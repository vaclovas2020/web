/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package web

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/go-git/go-git/v5"
	http2 "github.com/go-git/go-git/v5/plumbing/transport/http"
	"gopkg.in/yaml.v2"
	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class/method"
	"webimizer.dev/web/core/server"
	"webimizer.dev/web/parser"
	"webimizer.dev/webimizer"
)

/* Weblang version string */
const Version string = "weblang0.8.6"

/* Main VM struct */
type VM struct {
	memory base.MemoryMap          // Global MemoryMap
	server *server.Server          // Global Server
	events map[string]EventHandler // Global events handlers map
	wg     *sync.WaitGroup         // WaitGroup for goroutines control
	config *Config                 // Yaml parsed Config struct pointer
}

// Weblang now require configuration file weblang.yml in your project directory to work correctly.
// Initialize VM environment. Please provide correct configFile (config file name).
func (vm *VM) InitVM(configFile string) {
	err := vm.parseConfig(configFile)
	if err != nil {
		panic(err.Error())
	}
	sourceDir := vm.config.Project.Directories.SourceDir
	byteCodeDir := vm.config.Project.Directories.BytecodeDir
	log.Printf("\033[32m[weblang]\033[0m Starting application %v v%v in safe VM environment", vm.config.Project.Name, vm.config.Project.Version)
	err = vm.makeByteCodeDir(byteCodeDir)
	if err != nil {
		panic(err.Error())
	}
	vm.memory = base.MemoryMap{}
	vm.memory.Classes = make(map[string]*base.Class)
	vm.memory.Objects = make(map[string]*base.Object)
	vm.server = &server.Server{}
	vm.wg = &sync.WaitGroup{}
	count := 0
	output := make(chan string)
	vm.loadSourceDir(&count, sourceDir, byteCodeDir, output)
	go vm.monitorWorker(vm.wg, output)
	done := make(chan bool, 1)
	go vm.printWorker(count, output, done)
	<-done
}

/* Parse weblang.yml config file  */
func (vm *VM) parseConfig(configFile string) error {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	vm.config = &Config{}
	err = yaml.Unmarshal(data, vm.config)
	if err != nil {
		return err
	}
	return nil
}

/* Starts server process in VM environment */
func (vm *VM) StartServer() error {
	if vm.server != nil && vm.server.Host != "" && vm.server.Port > 0 {
		if err := vm.loadBeforeStaticFilesInitEvent(); err != nil {
			return err
		}
		return vm.server.Start()
	}
	return nil
}

/* Clone git repository to gitLocalDir and listen request from webHook url for git changes */
func (vm *VM) GitPreperWebHook(gitUrl string, gitUser string, gitToken string, gitLocalDir string, gitWebHook string) error {
	if _, err := os.Stat(gitLocalDir); os.IsNotExist(err) {
		_, err = git.PlainClone(gitLocalDir, false, &git.CloneOptions{
			Auth: &http2.BasicAuth{
				Username: gitUser,
				Password: gitToken,
			},
			URL:      gitUrl,
			Progress: os.Stdout,
		})
		if err != nil {
			return err
		}
	}
	http.Handle(gitWebHook, webimizer.HttpHandlerStruct{
		Handler: webimizer.HttpHandler(func(rw http.ResponseWriter, r *http.Request) {
			repository, err := git.PlainOpen(gitLocalDir)
			if err != nil {
				log.Fatal(err)
			}
			worktree, err := repository.Worktree()
			if err != nil {
				log.Fatal(err)
			}
			err = worktree.Pull(&git.PullOptions{
				RemoteName: "origin",
				Auth: &http2.BasicAuth{
					Username: gitUser,
					Password: gitToken,
				}, Progress: os.Stdout,
			})
			if err != nil && err != git.NoErrAlreadyUpToDate {
				log.Println(err.Error())
			}
			fmt.Fprintf(rw, "OK")
		}),
		AllowedMethods: []string{"POST"},
	}.Build())
	return nil
}

/* Set handler to specific class method (works with external methods only) */
func (vm *VM) DefineFunc(className string, methodName string, handler base.FunctionHandler) {
	if v, found := vm.memory.Classes[className].Methods[methodName]; found {
		if v.ClassMethod != nil &&
			(v.ClassMethod.MethodType == method.MethodType_ExternalPublic ||
				v.ClassMethod.MethodType == method.MethodType_ExternalPrivate ||
				v.ClassMethod.MethodType == method.MethodType_ExternalProtected) {
			v.Handler = handler
		}
	}
}

/* parse source file */
func (vm *VM) loadSourceDir(count *int, sourceDir string, byteCodeDir string, output chan<- string) {
	files, err := os.ReadDir(sourceDir)
	if err != nil {
		panic(err.Error())
	}
	for _, file := range files {
		if !file.IsDir() && (strings.Contains(file.Name(), ".web") || strings.Contains(file.Name(), ".weblang")) {
			vm.wg.Add(1)
			*count++
			log.Printf("\033[32m[weblang]\033[0m Loading %v worker(goroutine) for file '%v/%v' parsing...", *count, sourceDir, file.Name())
			info, err := file.Info()
			if err != nil {
				panic(err.Error())
			}
			go vm.parseFileWorker(vm.wg, fmt.Sprintf("%v/%v", sourceDir, file.Name()), fmt.Sprintf("%v/%v", byteCodeDir, strings.Replace(file.Name(), vm.getFileExt(&info), ".webc", 1)), output)
		} else if file.IsDir() {
			vm.makeByteCodeDir(fmt.Sprintf("%v/%v", byteCodeDir, file.Name()))
			vm.loadSourceDir(count, fmt.Sprintf("%v/%v", sourceDir, file.Name()), fmt.Sprintf("%v/%v", byteCodeDir, file.Name()), output)
		}
	}
}

func (vm *VM) getFileExt(file *fs.FileInfo) string {
	if strings.Contains((*file).Name(), ".web") {
		return ".web"
	}
	return ".weblang"
}

/* print output and/or error to screen */
func (vm *VM) printWorker(count int, output <-chan string, done chan<- bool) {
	for i := 0; i < count; i++ {
		log.Println(<-output)
	}
	done <- true
}

/* wait until all workers finish and close channels */
func (vm *VM) monitorWorker(wg *sync.WaitGroup, output chan<- string) {
	wg.Wait()
	close(output)
}

/* load source file from disk. Still not yet fully implemented */
func (vm *VM) parseFileWorker(wg *sync.WaitGroup, fileName string, byteCodeFileName string, output chan<- string) {
	defer wg.Done()
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	sourceCode := string(data)
	parser := &parser.Parser{Memory: &vm.memory, Server: vm.server}
	err = parser.Parse(sourceCode, fileName, byteCodeFileName)
	if err != nil {
		panic(err.Error())
	}
	output <- fmt.Sprintf("\033[32m[weblang]\033[0m Parsed file '%v'...", fileName)
}

func (vm *VM) makeByteCodeDir(byteCodeDir string) error {
	if _, err := os.Stat(byteCodeDir); os.IsNotExist(err) {
		err = os.Mkdir(byteCodeDir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create bytecode dir `%v`", byteCodeDir)
		} else {
			log.Printf("created bytecode dir `%v`", byteCodeDir)
		}
	}
	return nil
}
