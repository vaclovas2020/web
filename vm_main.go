package web

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"webimizer.dev/web/base"
	"webimizer.dev/web/parser"
)

/* Main VM struct */
type VM struct {
	classes map[string]base.Class
	parser  *parser.Parser
	wg      *sync.WaitGroup
}

/* Initialize VM with given context and arguments. Please provide correct sourceDir - directory of Web language source files */
func (vm *VM) InitVM(ctx context.Context, args []string, sourceDir string) {
	vm.classes = make(map[string]base.Class)
	vm.parser = &parser.Parser{Classes: &vm.classes}
	vm.wg = &sync.WaitGroup{}
	count := 0
	output := make(chan string)
	vm.loadSourceDir(&count, ctx, sourceDir, output)
	go vm.monitorWorker(vm.wg, output)
	done := make(chan bool, 1)
	go vm.printWorker(count, output, done)
	<-done
}

/* parse source file */
func (vm *VM) loadSourceDir(count *int, ctx context.Context, sourceDir string, output chan<- string) {
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		panic(err.Error())
	}
	for _, file := range files {
		if !file.IsDir() {
			vm.wg.Add(1)
			go vm.parseFileWorker(vm.wg, fmt.Sprintf("%v/%v", sourceDir, file.Name()), ctx, output)
			*count++
		} else {
			vm.loadSourceDir(count, ctx, fmt.Sprintf("%v/%v", sourceDir, file.Name()), output)
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
func (vm *VM) parseFileWorker(wg *sync.WaitGroup, fileName string, ctx context.Context, output chan<- string) {
	defer wg.Done()
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	sourceCode := string(data)
	err = vm.parser.Parse(sourceCode)
	if err != nil {
		panic(err.Error())
	}
	output <- "\033[32m[weblang]\033[0m Parsed file '" + fileName + "' successfully"
}
