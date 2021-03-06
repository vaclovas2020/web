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
	if objPtr != nil {
		objPtr.Attributes = make(map[string]interface{})
		objPtr.AttributesType = make(map[string]uint8)
	}
	for i := int64(0); i < classPtr.ByteCode.Header.AttributesCount; i++ {
		if err := loader.loadClassAttribute(classPtr, objPtr); err != nil {
			return fmt.Errorf("loadClassAttribute %v: %v", i, err.Error())
		}
	}
	return nil
}
