/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"bytes"
	"encoding/binary"
	"errors"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class/attribute"
)

func (loader *Loader) readAttributeFloat64(attrName string, attrPtr *attribute.Attribute, objPtr *base.Object) error {
	if attrPtr == nil {
		return errors.New("bug detected: attrPtr is nil")
	}
	buf := &bytes.Buffer{}
	_, err := buf.Write(attrPtr.Value)
	if err != nil {
		return err
	}
	var value float64
	err = binary.Read(buf, binary.BigEndian, &value)
	if err != nil {
		return err
	}
	objPtr.Attributes[attrName] = value
	return nil
}
