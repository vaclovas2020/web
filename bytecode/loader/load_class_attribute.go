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
	data, err := loader.readData(attribute.AttributeHeaderSize)
	if err != nil {
		return fmt.Errorf("loadClassAttribute: %v", err.Error())
	}
	buf := &bytes.Buffer{}
	_, err = buf.Write(data)
	if err != nil {
		return fmt.Errorf("loadClassAttribute: %v", err.Error())
	}
	attrHeader := &attribute.AttributeHeader{}
	if err := binary.Read(buf, binary.BigEndian, attrHeader); err != nil {
		return fmt.Errorf("loadClassAttribute: %v", err.Error())
	}
	attrPtr := &attribute.Attribute{Header: attrHeader}
	if objPtr == nil {
		classPtr.ByteCode.Attribute = append(classPtr.ByteCode.Attribute, attrPtr)
		return nil
	}
	if err := loader.readAttribute(attrPtr, classPtr, objPtr); err != nil {
		return fmt.Errorf("loadClassAttribute: %v", err.Error())
	}
	classPtr.ByteCode.Attribute = append(classPtr.ByteCode.Attribute, attrPtr)
	return nil
}
