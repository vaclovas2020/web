package bytecode

/* Header struct for classMethod declaration in Weblang bytecode file */
type ClassMethodHeader struct {
	MethodName               [80]byte // Class method name (max 80 symbols)
	InstructionSetStartIndex uint64   // Index of first InstructionSet block
	InstructionSetEndIndex   uint64   // Index of last InstructionSet block
}
