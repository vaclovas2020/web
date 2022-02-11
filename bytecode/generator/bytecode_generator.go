package generator

import (
	"webimizer.dev/web/bytecode/class"
	"webimizer.dev/web/bytecode/class/method"
	"webimizer.dev/web/bytecode/instructionset"
)

/* Weblang bytecode class file generator for use in Weblang VM */
type ByteCodeGenerator struct {
	Header       *class.ClassHeader               // Class header struct pointer
	ClassMethods []*method.ClassMethodHeader      // Array of all class methods
	Instructions []*instructionset.InstructionSet // Instructions array
}
