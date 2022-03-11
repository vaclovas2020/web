/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class/attribute"
)

/* Load class attributes from bytecode */
func (loader *Loader) loadClassAttribute(classPtr *base.Class, objPtr *base.Object) error {
	if loader.filePos+attribute.AttributeHeaderSize > (*loader.fileStat).Size() {
		return fmt.Errorf("eof reached when try to read attribute header from '%v' file", loader.ByteCodeFileName)
	}
	data := make([]byte, attribute.AttributeHeaderSize)
	count, err := loader.file.ReadAt(data, loader.filePos)
	if err != nil {
		return err
	}
	loader.filePos += int64(count)
	buf := &bytes.Buffer{}
	_, err = buf.Write(data)
	if err != nil {
		return err
	}
	attrHeader := &attribute.AttributeHeader{}
	if err := binary.Read(buf, binary.BigEndian, attrHeader); err != nil {
		return err
	}
	attrPtr := &attribute.Attribute{Header: attrHeader}
	if objPtr == nil {
		classPtr.ByteCode.Attribute = append(classPtr.ByteCode.Attribute, attrPtr)
		return nil
	}
	if err := loader.readAttribute(attrPtr, classPtr, objPtr); err != nil {
		return err
	}
	classPtr.ByteCode.Attribute = append(classPtr.ByteCode.Attribute, attrPtr)
	return nil
}
