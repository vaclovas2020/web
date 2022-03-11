/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package generator

import (
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
		if err := generator.generateAttributeNameLength(&attrHeader, attrName); err != nil {
			return err
		}
		if err := generator.writeAttributeHeader(&attrHeader); err != nil {
			return err
		}
		if err := generator.writeAttributeName(&attrHeader, attrName); err != nil {
			return err
		}
		if err := generator.writeAttributeValue(attrName, &el); err != nil {
			return err
		}
	}
	return nil
}

/* Get attribute size */
func (generator *ByteCodeGenerator) getAttributeSize(attributeName string, el *interface{}) (int64, error) {
	attrType := generator.Object.AttributesType[attributeName]
	var size int64 = 0
	switch attrType {
	case attribute.AttributeType_Float:
		size = 64
	case attribute.AttributeType_Int:
		size = 64
	case attribute.AttributeType_ObjReference:
		size = int64(len((*el).(string)))
	case attribute.AttributeType_String:
		size = int64(len((*el).(string)))
	}
	return size, nil
}

/* Write AttributeName to byteBuffer  */
func (generator *ByteCodeGenerator) writeAttributeName(header *attribute.AttributeHeader, attributeName string) error {
	if len(attributeName) > 80 {
		return fmt.Errorf("attribute name '%v' is too long (max 80 allowed)", attributeName)
	}
	err := binary.Write(generator.byteBuffer, binary.BigEndian, []byte(attributeName))
	if err != nil {
		return err
	}
	return nil
}

/* Write AttributeNameLength to AttributeHeader  */
func (generator *ByteCodeGenerator) generateAttributeNameLength(header *attribute.AttributeHeader, attributeName string) error {
	if len(attributeName) > 80 {
		return fmt.Errorf("attribute name '%v' is too long (max 80 allowed)", attributeName)
	}
	header.AttributeNameLength = int64(len(attributeName))
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

/* Write Attribute Value to buffer */
func (generator *ByteCodeGenerator) writeAttributeValue(attributeName string, el *interface{}) error {
	attrType := generator.Object.AttributesType[attributeName]
	switch attrType {
	case attribute.AttributeType_Float:
		err := binary.Write(generator.byteBuffer, binary.BigEndian, *el)
		if err != nil {
			return err
		}
	case attribute.AttributeType_Int:
		err := binary.Write(generator.byteBuffer, binary.BigEndian, *el)
		if err != nil {
			return err
		}
	case attribute.AttributeType_ObjReference:
		err := binary.Write(generator.byteBuffer, binary.BigEndian, []byte((*el).(string)))
		if err != nil {
			return err
		}
	case attribute.AttributeType_String:
		err := binary.Write(generator.byteBuffer, binary.BigEndian, []byte((*el).(string)))
		if err != nil {
			return err
		}
	}
	return nil
}
