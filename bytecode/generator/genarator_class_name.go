/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package generator

import (
	"encoding/binary"
	"fmt"
)

/* Set ClassName array to Bytecode struct (max 80 symbols allowed) */
func (generator *ByteCodeGenerator) writeClassName() error {
	if len(generator.ClassName) > 80 {
		return fmt.Errorf("class name '%v' is too long (max 80 allowed)", generator.ClassName)
	}
	err := binary.Write(generator.byteBuffer, binary.BigEndian, []byte(generator.ClassName))
	if err != nil {
		return err
	}
	return nil
}

/* Set ClassName array to Bytecode struct (max 80 symbols allowed) */
func (generator *ByteCodeGenerator) generateClassNameLength() error {
	if len(generator.ClassName) > 80 {
		return fmt.Errorf("class name '%v' is too long (max 80 allowed)", generator.ClassName)
	}
	generator.Class.ByteCode.Header.ClassNameLength = uint64(len(generator.ClassName))
	return nil
}
