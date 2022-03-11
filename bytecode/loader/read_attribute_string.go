/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"errors"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class/attribute"
)

func (loader *Loader) readAttributeString(attrName string, attrPtr *attribute.Attribute, objPtr *base.Object) error {
	if attrPtr == nil {
		return errors.New("readAttributeString: attrPtr is nil")
	}
	objPtr.Attributes[attrName] = string(attrPtr.Value)
	return nil
}
