/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* bytcode loader package */
package loader

import (
	"os"

	"webimizer.dev/web/base"
)

/* Bytecode loader */
type Loader struct {
	SourceCodeFileName string       // Source code file name
	ByteCodeFileName   string       // Byte code file name
	file               *os.File     // Bytecode file reader
	filePos            int64        // Current bytecode file buffer position
	fileStat           *os.FileInfo // Current bytecode file info
}

/* Load and fully parse bytecode data to *base.Class and *base.Object */
/* obj can be nil */
func (loader *Loader) LoadClassAndObject(class *base.Class, obj *base.Object) (bool, error) {
	isValid, err := loader.isValidByteCode(class)
	defer loader.closeFile()
	if err != nil {
		return false, err
	}
	return loader.runIfValid(isValid, class, obj)
}
