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
	if err != nil {
		panic(err.Error())
	}
	for _, file := range files {
		if !file.IsDir() {
			wg.Add(1)
			go vm.loadFileWorker(wg, file.Name(), ctx, output, errCh)
		}
	}
	go vm.monitorWorker(wg, output, errCh)
	done := make(chan bool, 1)
	go vm.printWorker(output, errCh, done)
	<-done
}

/* print output and/or error to screen */
func (vm *VM) printWorker(output <-chan string, errCh <-chan string, done chan<- bool) {
	for line := range output {
		fmt.Println(line)
	}
	for errLine := range errCh {
		fmt.Println("Fatal error: " + errLine)
	}
	done <- true
}

/* wait until all workers finish and close channels */
func (vm *VM) monitorWorker(wg *sync.WaitGroup, output chan<- string, errCh chan<- string) {
	wg.Wait()
	close(errCh)
	close(output)
}

/* load source file from disk. Still not yet fully implemented */
func (vm *VM) loadFileWorker(wg *sync.WaitGroup, fileName string, ctx context.Context, output chan<- string, errCh chan<- string) {
	defer wg.Done()
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		errCh <- err.Error()
	}
	vm.classes = make(map[string]Class)
	output <- fmt.Sprintf("DEBUG: Source file '%v' content:%v", fileName, string(data)) // Debug info for testing
	// TODO: parse source and create Class struct array
}
