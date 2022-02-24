/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* Class struct of Weblang bytecode */
package class

import (
	"bytes"
	"encoding/binary"
)

/* Constant for current Weblang  bytecode version */
const ByteCodeVersion uint16 = 0x0002

/* Header struct for Weblang bytecode class file. Binary data block size: 336 bytes */
type ClassHeader struct {
	FileFormatText      [8]byte  // First 8 bytes of file. It's contains "WEBLANG\x0f"
	ByteCodeVersion     uint16   // Version of bytecode file
	ClassType           uint8    // Class type: object, server, controller, router, model, repository, service and view
	ClassName           [80]byte // Class name (max 80 symbols)
	InstructionSetCount uint64   // Count of InstructionSet struct
	ClassMethodsCount   uint64   // Count of ClassMethodHeader struct
	AttributesCount     uint64   // Count of declared class atributes
	SourceCodeHash      [32]byte // Sha-256 hash of sourcecode file. For sourcefile changes detection
}

/* Write header data to struct. Use to prepare for write header struct to file */
func (header *ClassHeader) WriteHeader() error {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, []byte("WEBLANG\x0f"))
	if err != nil {
		return err
	}
	data := buf.Bytes()
	for i, v := range data {
		header.FileFormatText[i] = v
	}
	header.ByteCodeVersion = ByteCodeVersion
	return nil
}
