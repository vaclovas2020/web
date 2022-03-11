/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* bytcode loader package */
package loader

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"os"

	"webimizer.dev/web/bytecode/class"
)

/* Bytecode loader */
type Loader struct {
	SourceCodeFileName string   // Source code file name
	ByteCodeFileName   string   // Byte code file name
	file               *os.File // Bytecode file reader
}

/* Is valid bytecode file for use in VM environment */
func (loader *Loader) IsValidByteCode() (bool, error) {
	bytecodeFile, err := os.Open(loader.ByteCodeFileName)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	defer bytecodeFile.Close() // Close bytecode file at the end no matter what
	sourceCode, err := os.ReadFile(loader.SourceCodeFileName)
	if err != nil {
		return false, err
	}
	sha256Source := sha256.Sum256(sourceCode)
	header := class.ClassHeader{}
	err = binary.Read(bytecodeFile, binary.BigEndian, &header)
	if err != nil {
		return false, err
	}
	return (bytes.Equal(sha256Source[:], header.SourceCodeHash[:])), nil
}
