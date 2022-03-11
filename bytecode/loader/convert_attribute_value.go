/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"fmt"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class/attribute"
)

/* Convert byte[] slice to variable */
func (loader *Loader) convertAttributeValue(attrName string, attrPtr *attribute.Attribute, classPtr *base.Class, objPtr *base.Object) error {
	if objPtr.AttributesType == nil {
		objPtr.AttributesType = make(map[string]uint8)
	}
	objPtr.AttributesType[attrName] = attrPtr.Header.AttributeType
	switch attrPtr.Header.AttributeType {
	case attribute.AttributeType_Int:
		err := loader.readAttributeInt64(attrName, attrPtr, objPtr)
		if err != nil {
			return fmt.Errorf("convertAttributeValue: %v", err.Error())
		}
	case attribute.AttributeType_Float:
		err := loader.readAttributeFloat64(attrName, attrPtr, objPtr)
		if err != nil {
			return fmt.Errorf("convertAttributeValue: %v", err.Error())
		}
	case attribute.AttributeType_String:
		err := loader.readAttributeString(attrName, attrPtr, objPtr)
		if err != nil {
			return fmt.Errorf("convertAttributeValue: %v", err.Error())
		}
	case attribute.AttributeType_ObjReference:
		err := loader.readAttributeString(attrName, attrPtr, objPtr)
		if err != nil {
			return fmt.Errorf("convertAttributeValue: %v", err.Error())
		}
	}
	return nil
}
