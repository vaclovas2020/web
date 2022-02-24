/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package generator

func (generator *ByteCodeGenerator) generateAttributesCount() error {
	generator.Class.ByteCode.Header.AttributesCount = 0
	if generator.Object != nil {
		generator.Class.ByteCode.Header.AttributesCount = uint64(len(generator.Object.Attributes))
	}
	return nil
}
