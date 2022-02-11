package bytecode

import (
	"bytes"
	"encoding/binary"
)

/* Constant for current Weblang  bytecode version */
const ByteCodeVersion uint16 = 0x0001

/* Header struct for Weblang bytecode class file */
type ClassHeader struct {
	FileFormatText      [8]byte  // First 8 bytes of file. It's contains "WEBLANG\x0f"
	ClassName           [80]byte // Class name (max 80 symbols)
	ClassType           uint8    // Class type: class, server, controller, router, model, repository, service and view
	ByteCodeVersion     uint16   // Version of bytecode file
	InstructionSetCount uint64   // Count of InstructionSet struct
	ClassMethodsCount   uint64   // Count of ClassMethodHeader struct
}

/* Write header data to struct. Use to prepare for write header struct to file */
func (header *ClassHeader) WriteHeader() {
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.LittleEndian, []byte("WEBLANG\x0f"))
	data := buf.Bytes()
	for i, v := range data {
		header.FileFormatText[i] = v
	}
}
