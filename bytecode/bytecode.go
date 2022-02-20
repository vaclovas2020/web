/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* Weblang bytecode main package */
package bytecode

import (
	"webimizer.dev/web/bytecode/class"
	"webimizer.dev/web/bytecode/class/attribute"
	"webimizer.dev/web/bytecode/class/method"
	"webimizer.dev/web/bytecode/instructionset"
)

/* Weblang bytecode struct for use in Weblang VM */
type ByteCode struct {
	Header       *class.ClassHeader               // Class header struct pointer
	ClassMethods []*method.ClassMethod            // Array of all class methods
	Attribute    []*attribute.Attribute           // Attribute array
	Instructions []*instructionset.InstructionSet // Instructions array
}
