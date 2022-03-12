/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package generator

import (
	"encoding/binary"
)

/* Set ClassName array to Bytecode struct (max 80 symbols allowed) */
func (generator *ByteCodeGenerator) writeClassName() error {
	err := binary.Write(generator.byteBuffer, binary.BigEndian, []byte(generator.ClassName))
	if err != nil {
		return err
	}
	return nil
}

/* Set ClassName array to Bytecode struct (max 80 symbols allowed) */
func (generator *ByteCodeGenerator) generateClassNameLength() error {
	generator.Class.ByteCode.Header.ClassNameLength = int64(len(generator.ClassName))
	return nil
}
