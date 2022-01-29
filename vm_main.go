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
	var sourceCodes []string
	output := make(chan string)
	errCh := make(chan string)
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		panic(err.Error())
	}
	for _, file := range files {
		if !file.IsDir() {
			data, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", sourceDir, file.Name()))
			if err != nil {
				panic(err.Error())
			}
			sourceCodes = append(sourceCodes, string(data))
		}
	}
	for _, sourceCode := range sourceCodes {
		wg.Add(1)
		go vm.parseFileWorker(wg, sourceCode, ctx, output, errCh)
	}
	go vm.monitorWorker(wg, output, errCh)
	done := make(chan bool, 1)
	go vm.printWorker(len(sourceCodes), output, errCh, done)
	<-done
}

/* print output and/or error to screen */
func (vm *VM) printWorker(count int, output <-chan string, errCh <-chan string, done chan<- bool) {
	for i := 0; i < count; i++ {
		select {
		case <-output:
			fmt.Println(<-output)
		case <-errCh:
			fmt.Println("Fatal error: " + <-errCh)
		}
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
func (vm *VM) parseFileWorker(wg *sync.WaitGroup, sourceCode string, ctx context.Context, output chan<- string, errCh chan<- string) {
	defer wg.Done()
	vm.classes = make(map[string]Class)
	output <- sourceCode // Debug info for testing
	// TODO: parse source and create Class struct array
}
