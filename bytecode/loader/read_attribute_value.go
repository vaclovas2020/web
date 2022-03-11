/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"fmt"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class/attribute"
)

/* Read attribute value from bytecode */
func (loader *Loader) readAttributeValue(attrName string, attrPtr *attribute.Attribute, classPtr *base.Class, objPtr *base.Object) error {
	var err error = nil
	attrPtr.Value, err = loader.readData(attrPtr.Header.AttributeValueSize)
	if err != nil {
		return fmt.Errorf("readAttributeValue: %v", err.Error())
	}
	if err := loader.convertAttributeValue(attrName, attrPtr, classPtr, objPtr); err != nil {
		return fmt.Errorf("readAttributeValue: %v", err.Error())
	}
	return nil
}
