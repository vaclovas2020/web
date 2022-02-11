package bytecode

/* Header struct for classMethod declaration in Weblang bytecode file */
type ClassMethodHeader struct {
	ClassName                [80]byte // Class name
	InstructionSetStartIndex uint64   // Index of first InstructionSet block
	InstructionSetEndIndex   uint64   // Index of last InstructionSet block
}
