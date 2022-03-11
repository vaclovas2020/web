package loader

import (
	"fmt"

	"webimizer.dev/web/base"
)

func (loader *Loader) writeToMemory(nameStr string, classPtr *base.Class, objPtr *base.Object, memory *base.MemoryMap) error {
	if _, exist := memory.Classes[nameStr]; exist {
		return fmt.Errorf("class name '%v' already exists", nameStr)
	}
	memory.Classes[nameStr] = classPtr
	if objPtr != nil {
		if obj, exist := memory.Objects[nameStr]; exist && obj.Scope == 0 {
			return fmt.Errorf("object name '%v' already exists in global scope", nameStr)
		}
		memory.Objects[nameStr] = objPtr
	}
	return nil
}
