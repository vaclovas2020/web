/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package generator

import (
	"crypto/sha256"
	"io/ioutil"
)

/* Generate source code file sha256 sum */
func (generator *ByteCodeGenerator) generateSha256Sum() error {
	data, err := ioutil.ReadFile(generator.SourceCodeFileName)
	if err != nil {
		return err
	}
	generator.Class.ByteCode.Header.SourceCodeHash = sha256.Sum256(data)
	return nil
}
