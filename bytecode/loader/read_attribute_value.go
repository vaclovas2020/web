/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"fmt"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class/attribute"
)

/* Read attribute value from bytecode */
func (loader *Loader) readAttributeValue(attrName string, attrPtr *attribute.Attribute, classPtr *base.Class, objPtr *base.Object) error {
	if loader.filePos+int64(attrPtr.Header.AttributeValueSize) > (*loader.fileStat).Size() {
		return fmt.Errorf("eof reached when try to read attribute name from '%v' file", loader.ByteCodeFileName)
	}
	attrPtr.Value = make([]byte, attrPtr.Header.AttributeValueSize)
	count, err := loader.file.ReadAt(attrPtr.Value, loader.filePos)
	if err != nil {
		return err
	}
	loader.filePos += int64(count)
	if err := loader.convertAttributeValue(attrName, attrPtr, classPtr, objPtr); err != nil {
		return err
	}
	return nil
}
