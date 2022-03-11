/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"fmt"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class/attribute"
)

/* Load class name from bytecode */
func (loader *Loader) loadClassName(classPtr *base.Class, objPtr *base.Object, memory *base.MemoryMap) error {
	if err := loader.detectBug(classPtr, objPtr, memory); err != nil {
		return fmt.Errorf("loadClassName: %v", err.Error())
	}
	data, err := loader.readData(attribute.AttributeHeaderSize)
	if err != nil {
		return fmt.Errorf("loadClassName: %v", err.Error())
	}
	nameStr := string(data)
	err = loader.writeToMemory(nameStr, classPtr, objPtr, memory)
	if err != nil {
		return fmt.Errorf("loadClassName: %v", err.Error())
	}
	return nil
}
