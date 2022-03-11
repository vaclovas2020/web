/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"fmt"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class/attribute"
)

/* Read attribute name from bytecode */
func (loader *Loader) readAttribute(attrPtr *attribute.Attribute, classPtr *base.Class, objPtr *base.Object) error {
	data, err := loader.readData(int64(attrPtr.Header.AttributeNameLength))
	if err != nil {
		return fmt.Errorf("readAttribute: %v", err.Error())
	}
	if err := loader.readAttributeValue(string(data), attrPtr, classPtr, objPtr); err != nil {
		return fmt.Errorf("readAttribute: %v", err.Error())
	}
	return nil
}
