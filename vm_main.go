package web

import (
	"context"
	"io/ioutil"
	"sync"
)

/* Initialize VM with given context and arguments. Please provide correct sourceDir - directory of Web language source files */
func (vm *VM) InitVM(ctx context.Context, args []string, sourceDir string, output chan<- string, errCh chan<- string, done chan<- bool) {
	wg := &sync.WaitGroup{}
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		errCh <- err.Error()
	}
	for _, file := range files {
		if !file.IsDir() {
			wg.Add(1)
			go vm.loadFile(wg, file.Name(), ctx, output, errCh)
		}
	}
	go wait(wg, output, errCh, done)
}

func wait(wg *sync.WaitGroup, output chan<- string, err chan<- string, done chan<- bool) {
	wg.Wait()
	close(err)
	close(output)
	done <- true
}

/* load source file from disk. Still not yet fully implemented */
func (vm *VM) loadFile(wg *sync.WaitGroup, fileName string, ctx context.Context, output chan<- string, err chan<- string) {
	defer wg.Done()
	data, error := ioutil.ReadFile(fileName)
	if error != nil {
		err <- error.Error()
	}
	sourceString := string(data)
	output <- sourceString
	// TODO: parse source and create Class struct array
}
