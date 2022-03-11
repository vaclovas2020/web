/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"encoding/binary"
	"os"

	"webimizer.dev/web/base"
)

/* Load bytecode to class struct */
func (loader *Loader) loadClassHader(class *base.Class) error {
	bytecodeFile, err := os.Open(loader.ByteCodeFileName)
	if err != nil {
		return err
	}
	defer bytecodeFile.Close() // Close bytecode file at the end no matter what
	err = binary.Read(bytecodeFile, binary.BigEndian, &class.ByteCode.Header)
	if err != nil {
		return err
	}
	return nil
}
