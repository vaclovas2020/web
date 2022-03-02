/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package generator

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"webimizer.dev/web/bytecode/class/attribute"
)

/* Generate attributes and write to bytecode buffer */
func (generator *ByteCodeGenerator) generateAttributes() error {
	for attrName, el := range generator.Object.Attributes {
		size, err := generator.getAttributeSize(attrName, &el)
		if err != nil {
			return err
		}
		attrHeader := attribute.AttributeHeader{AttributeType: generator.Object.AttributesType[attrName], AttributeValueSize: size}
		if err := generator.generateAttributeName(&attrHeader, attrName); err != nil {
			return err
		}
		if err := generator.writeAttributeHeader(&attrHeader); err != nil {
			return err
		}
		if err := generator.writeAttribute(attrName, &el); err != nil {
			return err
		}
	}
	return nil
}

/* Get attribute size */
func (generator *ByteCodeGenerator) getAttributeSize(attributeName string, el *interface{}) (uint64, error) {
	attrType := generator.Object.AttributesType[attributeName]
	var size uint64 = 0
	switch attrType {
	case attribute.AttributeType_Float:
	case attribute.AttributeType_Int:
		size = 64
	case attribute.AttributeType_ObjReference:
	case attribute.AttributeType_String:
		size = uint64(len((*el).(string)))
	}
	return size, nil
}

/* Set AttributeName array to Bytecode struct (max 80 symbols allowed) */
func (generator *ByteCodeGenerator) generateAttributeName(header *attribute.AttributeHeader, attributeName string) error {
	if len(attributeName) > 80 {
		return fmt.Errorf("attribute name '%v' is too long (max 80 allowed)", attributeName)
	}
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.BigEndian, []byte(attributeName))
	if err != nil {
		return err
	}
	data := buf.Bytes()
	for i, v := range data {
		header.AttributeName[i] = v
	}
	return nil
}

/* Write AttributeHeader to buffer */
func (generator *ByteCodeGenerator) writeAttributeHeader(header *attribute.AttributeHeader) error {
	err := binary.Write(generator.byteBuffer, binary.BigEndian, *header)
	if err != nil {
		return err
	}
	return nil
}

/* Write AttributeHeader to buffer */
func (generator *ByteCodeGenerator) writeAttribute(attributeName string, el *interface{}) error {
	attrType := generator.Object.AttributesType[attributeName]
	switch attrType {
	case attribute.AttributeType_Float:
	case attribute.AttributeType_Int:
		err := binary.Write(generator.byteBuffer, binary.BigEndian, *el)
		if err != nil {
			return err
		}
	case attribute.AttributeType_ObjReference:
	case attribute.AttributeType_String:
		err := binary.Write(generator.byteBuffer, binary.BigEndian, []byte((*el).(string)))
		if err != nil {
			return err
		}
	}
	return nil
}
