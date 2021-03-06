/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* Instuctionset struct of Weblang bytecode files */
package instructionset

/* Header struct for dynamic length instruction block for Weblang VM */
type InstructionSetHeader struct {
	Index  int64  // Index number
	Type   uint16 // InstructionSet type
	Length int64  // Instruction block size (bytes)
}

/* Dynamic length instruction set for Weblang Vm */
type InstructionSet struct {
	Header   *InstructionSetHeader // Heaader info: index number, type and length
	ByteCode []byte                // data array (bytes)
}
