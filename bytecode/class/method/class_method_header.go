/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package method

/* Header struct for classMethod declaration in Weblang bytecode file */
type ClassMethod struct {
	MethodName               [80]byte // Class method name (max 80 symbols)
	MethodType               uint8    // Method type public, private, protected, external
	InstructionSetStartIndex uint64   // Index of first InstructionSet block
	InstructionSetEndIndex   uint64   // Index of last InstructionSet block
}
