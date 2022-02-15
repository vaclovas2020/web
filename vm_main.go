/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class"
	"webimizer.dev/web/bytecode/class/method"
	"webimizer.dev/web/parser"
	"webimizer.dev/web/server"
)

/* Weblang version string */
const Version string = "v0.3.15"

/* Main VM struct */
type VM struct {
	stack  base.MemoryStack // Global MemoryStack
	parser *parser.Parser   // Global Parser
	server *server.Server   // Global server
	wg     *sync.WaitGroup  // WaitGroup for goroutines control
}

/* Initialize VM with given context and arguments. Please provide correct sourceDir (directory of Web language source files) and byteCodeDir (direcotry for bytecode files) */
func (vm *VM) InitVM(sourceDir string, byteCodeDir string) {
	fmt.Println("----------------------")
	fmt.Printf("Welcome to Weblang %v (bytecode version %v)\n\n", Version, class.ByteCodeVersion)
	fmt.Println("Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved.")
	fmt.Println("License: BSD-3-Clause License")
	fmt.Println("----------------------")
	log.Println("\033[32m[weblang]\033[0m Preparing VM environment...")
	vm.stack = base.MemoryStack{}
	vm.stack.Classes = make(map[string]base.Class)
	vm.stack.Objects = make(map[string]base.Object)
	vm.server = &server.Server{}
	vm.parser = &parser.Parser{Stack: &vm.stack, Server: vm.server}
	vm.wg = &sync.WaitGroup{}
	count := 0
	output := make(chan string)
	vm.loadSourceDir(&count, sourceDir, byteCodeDir, output)
	go vm.monitorWorker(vm.wg, output)
	done := make(chan bool, 1)
	go vm.printWorker(count, output, done)
	<-done
}

/* Starts server process in VM environment */
func (vm *VM) StartServer() error {
	if vm.server != nil {
		return vm.server.Start()
	}
	return nil
}

/* Set handler to specific class method (works with external methods only) */
func (vm *VM) DefineFunc(className string, methodName string, handler base.FunctionHandler) {
	if v, found := vm.stack.Classes[className].Methods[methodName]; found {
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
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		panic(err.Error())
	}
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), ".web") {
			vm.wg.Add(1)
			*count++
			log.Printf("\033[32m[weblang]\033[0m Loading %v worker(goroutine) for file '%v/%v' parsing...", *count, sourceDir, file.Name())
			go vm.parseFileWorker(vm.wg, fmt.Sprintf("%v/%v", sourceDir, file.Name()), fmt.Sprintf("%v/%v", byteCodeDir, strings.Replace(file.Name(), ".web", "webc", 1)), output)
		} else if file.IsDir() {
			vm.loadSourceDir(count, fmt.Sprintf("%v/%v", sourceDir, file.Name()), fmt.Sprintf("%v/%v", byteCodeDir, file.Name()), output)
		}
	}
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
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("\033[33m[weblang]\033[0m %v", err.Error())
	}
	sourceCode := string(data)
	err = vm.parser.Parse(sourceCode, fileName, byteCodeFileName)
	if err != nil {
		log.Fatalf("\033[33m[weblang]\033[0m %v", err.Error())
	}
	output <- fmt.Sprintf("\033[32m[weblang]\033[0m Parsed file '%v'...", fileName)
}
