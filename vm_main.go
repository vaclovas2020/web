package web

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"
)

/* Initialize VM with given context and arguments. Please provide correct sourceDir - directory of Web language source files */
func (vm *VM) InitVM(ctx context.Context, args []string, sourceDir string) {
	wg := &sync.WaitGroup{}
	output := make(chan string)
	errCh := make(chan string)
	files, err := ioutil.ReadDir(sourceDir)
	parseDone := make(chan bool, 1)
	if err != nil {
		panic(err.Error())
	}
	for _, file := range files {
		if !file.IsDir() {
			data, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", sourceDir, file.Name()))
			if err != nil {
				panic(err.Error())
			}
			wg.Add(1)
			go vm.parseFileWorker(wg, string(data), ctx, output, errCh)
		}
	}
	go vm.monitorWorker(wg, output, errCh, parseDone)
	done := make(chan bool, 1)
	go vm.printWorker(output, errCh, done, parseDone)
	<-done
}

/* print output and/or error to screen */
func (vm *VM) printWorker(output <-chan string, errCh <-chan string, done chan<- bool, parseDone <-chan bool) {
	for {
		select {
		case <-output:
			fmt.Println(<-output)
		case <-errCh:
			fmt.Println("Fatal error: " + <-errCh)
		case <-parseDone:
			done <- <-parseDone
		}
	}
}

/* wait until all workers finish and close channels */
func (vm *VM) monitorWorker(wg *sync.WaitGroup, output chan<- string, errCh chan<- string, done chan<- bool) {
	wg.Wait()
	close(errCh)
	close(output)
	done <- true
}

/* load source file from disk. Still not yet fully implemented */
func (vm *VM) parseFileWorker(wg *sync.WaitGroup, sourceCode string, ctx context.Context, output chan<- string, errCh chan<- string) {
	defer wg.Done()
	vm.classes = make(map[string]Class)
	output <- sourceCode // Debug info for testing
	// TODO: parse source and create Class struct array
}
