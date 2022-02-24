/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package generator

import (
	"encoding/binary"
	"io/ioutil"
)

/* Write bytecode data to buffer */
func (generator *ByteCodeGenerator) write() error {
	err := generator.generateInner([]GeneratorHandler{
		GeneratorHandler(generator.writeHeader),
		GeneratorHandler(generator.writeBufferToFile),
	})
	if err != nil {
		return err
	}
	return nil
}

/* Write ClassHeader to buffer */
func (generator *ByteCodeGenerator) writeHeader() error {
	err := binary.Write(generator.byteBuffer, binary.BigEndian, *generator.Class.ByteCode.Header)
	if err != nil {
		return err
	}
	return nil
}

/* Write buffer to file */
func (generator *ByteCodeGenerator) writeBufferToFile() error {
	err := ioutil.WriteFile(generator.ByteCodeFileName, generator.byteBuffer.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}
