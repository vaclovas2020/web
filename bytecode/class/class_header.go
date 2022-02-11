package class

import (
	"bytes"
	"encoding/binary"
)

/* Constant for current Weblang  bytecode version */
const ByteCodeVersion uint16 = 0x0001

/* Header struct for Weblang bytecode class file. Binary data block size: 336 bytes */
type ClassHeader struct {
	FileFormatText      [8]byte  // First 8 bytes of file. It's contains "WEBLANG\x0f"
	ClassName           [80]byte // Class name (max 80 symbols)
	ClassType           uint8    // Class type: object, server, controller, router, model, repository, service and view
	ByteCodeVersion     uint16   // Version of bytecode file
	InstructionSetCount uint64   // Count of InstructionSet struct
	ClassMethodsCount   uint64   // Count of ClassMethodHeader struct
	AttributesCount     uint64   // Count of declared class atributes
	SourceCodeHash      [32]byte // Sha-256 hash of sourcecode file. For sourcefile changes detection
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