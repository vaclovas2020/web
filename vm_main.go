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
	done := make(chan bool, 1)
	go vm.printWorker(output, errCh, done)
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		errCh <- err.Error()
	}
	for _, file := range files {
		if !file.IsDir() {
			wg.Add(1)
			go vm.loadFileWorker(wg, file.Name(), ctx, output, errCh)
		}
	}
	go vm.monitorWorker(wg, output, errCh)
	<-done
}

/* print output and/or error to screen */
func (vm *VM) printWorker(output <-chan string, err <-chan string, done chan<- bool) {
	for {
		select {
		case <-output:
			fmt.Println(<-output)
		case <-err:
			fmt.Println("Fatal error:" + <-err)
		default:
			done <- true
		}
	}
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
	vm.classes = make(map[string]Class)
	output <- fileName
	// TODO: parse source and create Class struct array
}
