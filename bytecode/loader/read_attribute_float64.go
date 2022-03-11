/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class/attribute"
)

func (loader *Loader) readAttributeFloat64(attrName string, attrPtr *attribute.Attribute, objPtr *base.Object) error {
	if attrPtr == nil {
		return errors.New("readAttributeFloat64: attrPtr is nil")
	}
	buf := &bytes.Buffer{}
	_, err := buf.Write(attrPtr.Value)
	if err != nil {
		return fmt.Errorf("readAttributeFloat64: %v", err.Error())
	}
	var value float64
	err = binary.Read(buf, binary.BigEndian, &value)
	if err != nil {
		return fmt.Errorf("readAttributeFloat64: %v", err.Error())
	}
	if objPtr.Attributes == nil {
		objPtr.Attributes = make(map[string]interface{})
	}
	objPtr.Attributes[attrName] = value
	return nil
}
