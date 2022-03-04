/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* Weblang bytecode generator package */
package generator

import (
	"bytes"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode"
	"webimizer.dev/web/bytecode/class"
)

/* Weblang bytecode class file generator for use in Weblang VM */
type ByteCodeGenerator struct {
	Class              *base.Class   // Pointer to Class struct
	Object             *base.Object  // Poniter to Object struct.
	ClassName          string        // Class name
	ByteCodeFileName   string        // ByteCode file relative path
	SourceCodeFileName string        // SourceCode file relative path
	byteBuffer         *bytes.Buffer // Write Buffer
}

/* Generate bytecode file */
func (generator *ByteCodeGenerator) Generate() error {
	generator.Class.ByteCode = &bytecode.ByteCode{Header: &class.ClassHeader{}}
	generator.byteBuffer = &bytes.Buffer{}
	err := generator.generateInner([]GeneratorHandler{
		GeneratorHandler(generator.generateHeader),
		GeneratorHandler(generator.generateClassNameLength),
		GeneratorHandler(generator.generateClassType),
		GeneratorHandler(generator.generateSha256Sum),
		GeneratorHandler(generator.generateAttributesCount),
		GeneratorHandler(generator.write),
	})
	if err != nil {
		return err
	}
	return nil
}

/* Call to all GenerationHandler functions */
func (generator *ByteCodeGenerator) generateInner(handlerList []GeneratorHandler) error {
	for _, handler := range handlerList {
		err := handler()
		if err != nil {
			return err
		}
	}
	return nil
}
