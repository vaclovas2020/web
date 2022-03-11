/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"bytes"
	"crypto/sha256"
	"os"

	"webimizer.dev/web/base"
)

/* Is valid bytecode file for use in VM environment */
func (loader *Loader) isValidByteCode(class *base.Class) (bool, error) {
	loader.filePos = 0
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
	err = loader.loadClassHeader(class)
	if err != nil {
		return false, err
	}
	return bytes.Equal(sha256Source[:], class.ByteCode.Header.SourceCodeHash[:]), nil
}
