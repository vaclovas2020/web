/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package generator

/* Generate ByteCode header */
func (generator *ByteCodeGenerator) generateHeader() error {
	err := generator.Class.ByteCode.Header.WriteHeader()
	if err != nil {
		return err
	}
	return nil
}
