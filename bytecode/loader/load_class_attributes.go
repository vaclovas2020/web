/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"fmt"

	"webimizer.dev/web/base"
)

/* Load class attributes from bytecode */
func (loader *Loader) loadClassAttributes(classPtr *base.Class, objPtr *base.Object, memory *base.MemoryMap) error {
	if err := loader.detectBug(classPtr, objPtr, memory); err != nil {
		return fmt.Errorf("loadClassAttributes: %v", err.Error())
	}
	for i := uint64(0); i < classPtr.ByteCode.Header.AttributesCount; i++ {
		if err := loader.loadClassAttribute(classPtr, objPtr); err != nil {
			return fmt.Errorf("loadClassAttributes: %v", err.Error())
		}
	}
	return nil
}
