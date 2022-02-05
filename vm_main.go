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
}

/* Initialize VM with given context and arguments. Please provide correct sourceDir - directory of Web language source files */
func (vm *VM) InitVM(ctx context.Context, args []string, sourceDir string) {
	vm.classes = make(map[string]base.Class)
	wg := &sync.WaitGroup{}
	count := 0
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		panic(err.Error())
	}
	output := make(chan string)
	for _, file := range files {
		if !file.IsDir() {
			wg.Add(1)
			go vm.parseFileWorker(wg, fmt.Sprintf("%v/%v", sourceDir, file.Name()), ctx, output)
			count++
		}
	}
	go vm.monitorWorker(wg, output)
	done := make(chan bool, 1)
	go vm.printWorker(count, output, done)
	<-done
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
	vm.parser = &parser.Parser{Classes: &vm.classes}
	err = vm.parser.Parse(sourceCode)
	if err != nil {
		panic(err.Error())
	}
}
