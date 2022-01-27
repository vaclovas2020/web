package web

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
)

/* Initialize VM with given context and arguments. Please provide correct sourceDir - directory of Web language source files */
func (vm *VM) InitVM(ctx context.Context, args []string, sourceDir string, output chan<- string) error {
	errCh := make(chan string)
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			go vm.loadFile(file.Name(), ctx, output, errCh)
		}
		select {
		case <-errCh:
			return errors.New("Fatal error: " + <-errCh)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}

/* load source file from disk. Still not yet fully implemented */
func (vm *VM) loadFile(fileName string, ctx context.Context, output chan<- string, err chan<- string) {
	data, error := ioutil.ReadFile(fileName)
	if error != nil {
		err <- error.Error()
	}
	sourceString := string(data)
	output <- sourceString
	// TODO: parse source and create Class struct array
}
