package bytecode

import (
	"webimizer.dev/web/bytecode/class"
	"webimizer.dev/web/bytecode/class/method"
	"webimizer.dev/web/bytecode/instructionset"
)

/* Weblang bytecode struct for use in Weblang VM */
type ByteCode struct {
	Header       *class.ClassHeader               // Class header struct pointer
	ClassMethods []*method.ClassMethod            // Array of all class methods
	Instructions []*instructionset.InstructionSet // Instructions array
}
