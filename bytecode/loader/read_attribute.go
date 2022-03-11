/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"fmt"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class/attribute"
)

/* Read attribute name from bytecode */
func (loader *Loader) readAttribute(attrPtr *attribute.Attribute, classPtr *base.Class, objPtr *base.Object) error {
	if loader.filePos+int64(attrPtr.Header.AttributeNameLength) > (*loader.fileStat).Size() {
		return fmt.Errorf("eof reached when try to read attribute name from '%v' file", loader.ByteCodeFileName)
	}
	data := make([]byte, attrPtr.Header.AttributeNameLength)
	count, err := loader.file.ReadAt(data, loader.filePos)
	if err != nil {
		return err
	}
	loader.filePos += int64(count)
	if err := loader.readAttributeValue(string(data), attrPtr, classPtr, objPtr); err != nil {
		return err
	}
	return nil
}
