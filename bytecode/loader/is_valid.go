/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"os"

	"webimizer.dev/web/bytecode/class"
)

/* Is valid bytecode file for use in VM environment */
func (loader *Loader) isValidByteCode() (bool, error) {
	err := loader.openByteCodeFile()
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	sourceCode, err := os.ReadFile(loader.SourceCodeFileName)
	if err != nil {
		return false, err
	}
	sha256Source := sha256.Sum256(sourceCode)
	header := class.ClassHeader{}
	err = binary.Read(loader.file, binary.BigEndian, &header)
	if err != nil {
		return false, err
	}
	return bytes.Equal(sha256Source[:], header.SourceCodeHash[:]), nil
}
