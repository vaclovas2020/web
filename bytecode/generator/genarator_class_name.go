/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package generator

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/* Set ClassName array to Bytecode struct (max 80 symbols allowed) */
func (generator *ByteCodeGenerator) generateClassName() error {
	if len(generator.ClassName) > 80 {
		return fmt.Errorf("class name '%v' is too long (max 80 allowed)", generator.ClassName)
	}
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, []byte(generator.ClassName))
	if err != nil {
		return err
	}
	data := buf.Bytes()
	for i, v := range data {
		generator.Class.ByteCode.Header.ClassName[i] = v
	}
	return nil
}
