package bytecode

/* 255 bytes Instruction block for Weblang VM */
type InstructionSet struct {
	Index    uint64    // Index number
	Type     uint16    // InstructionSet type
	ByteCode [255]byte //  Array of 255 bytes
}
