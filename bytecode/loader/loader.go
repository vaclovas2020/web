/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* bytcode loader package */
package loader

import (
	"fmt"
	"os"

	"webimizer.dev/web/base"
	"webimizer.dev/web/core/server"
)

/* Bytecode loader */
type Loader struct {
	SourceCodeFileName string         // Source code file name
	ByteCodeFileName   string         // Byte code file name
	Server             *server.Server // Global server object
	file               *os.File       // Bytecode file reader
	filePos            int64          // Current bytecode file buffer position
	fileStat           *os.FileInfo   // Current bytecode file info
}

/* Load and fully parse bytecode data to *base.Class and *base.Object */
/* obj can be nil */
func (loader *Loader) LoadClassAndObject(memory *base.MemoryMap) (bool, error) {
	classPtr := &base.Class{}
	isValid, err := loader.isValidByteCode(classPtr)
	defer loader.closeFile()
	if err != nil {
		return false, fmt.Errorf("LoadClassAndObject: %v", err.Error())
	}
	return loader.parseIfValid(isValid, classPtr, memory)
}
