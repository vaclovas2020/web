/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* bytcode loader package */
package loader

import (
	"os"

	"webimizer.dev/web/base"
)

/* Bytecode loader */
type Loader struct {
	SourceCodeFileName string   // Source code file name
	ByteCodeFileName   string   // Byte code file name
	file               *os.File // Bytecode file reader
}

/* Load and fully parse bytecode data to *base.Class */
func (loader *Loader) LoadClass(class *base.Class) error {
	isValid, err := loader.isValidByteCode()
	defer loader.closeFile()
	if err != nil {
		return err
	}
	return loader.runIfValid(isValid, class)
}
