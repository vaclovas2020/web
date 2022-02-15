/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package generator

import (
	"webimizer.dev/web/base"
)

/* Weblang bytecode class file generator for use in Weblang VM */
type ByteCodeGenerator struct {
	Class  *base.Class
	Object *base.Object
}

/* Generate bytecode */
func (generator *ByteCodeGenerator) Generate(byteCodeFileName string, sourceCodeFileName string) error {
	// TODO: implement method
	return nil
}
