package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
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
func (vm *VM) InitVM(sourceDir string) {
	fmt.Println("----------------------")
	fmt.Println("Welcome to Weblang\n\n ")
	fmt.Println("Copyright (c) 2022, Vaclovas Lapinskis")
	fmt.Println("All rights reserved.")
	fmt.Println("License: BSD-3-Clause License")
	fmt.Println("----------------------")
	log.Println("\033[32m[weblang]\033[0m Preparing VM environment...")
	vm.classes = make(map[string]base.Class)
	vm.parser = &parser.Parser{Classes: &vm.classes}
	vm.wg = &sync.WaitGroup{}
	count := 0
	output := make(chan string)
	vm.loadSourceDir(&count, sourceDir, output)
	go vm.monitorWorker(vm.wg, output)
	done := make(chan bool, 1)
	go vm.printWorker(count, output, done)
	<-done
}

/* Set handler to specific class method */
func (vm *VM) DefineFunc(className string, methodName string, handler base.FunctionHandler) {
	if v, found := vm.classes[className].Methods[methodName]; found {
		v.Handler = handler
	}
}

/* parse source file */
func (vm *VM) loadSourceDir(count *int, sourceDir string, output chan<- string) {
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		panic(err.Error())
	}
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), ".web") {
			vm.wg.Add(1)
			*count++
			log.Printf("\033[32m[weblang]\033[0m Loading %v worker(goroutine) for file '%v/%v' parsing purpose... ", *count, sourceDir, file.Name())
			go vm.parseFileWorker(vm.wg, fmt.Sprintf("%v/%v", sourceDir, file.Name()), output)
		} else if file.IsDir() {
			vm.loadSourceDir(count, fmt.Sprintf("%v/%v", sourceDir, file.Name()), output)
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
	log.Printf("Classes: %v", vm.classes)
	close(output)
}

/* load source file from disk. Still not yet fully implemented */
func (vm *VM) parseFileWorker(wg *sync.WaitGroup, fileName string, output chan<- string) {
	defer wg.Done()
	output <- fmt.Sprintf("\033[32m[weblang]\033[0m Parsing file '%v'...", fileName)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	sourceCode := string(data)
	err = vm.parser.Parse(sourceCode)
	if err != nil {
		panic(err.Error())
	}
}
