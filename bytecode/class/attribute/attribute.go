/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package attribute

/* Class attributes header struct */
type AttributeHeader struct {
	AttributeName      [80]byte // Attribute name (max 80 symbols)
	AttributeType      uint8    // attribute type: string, int, float bool or objectName'
	AttributeValueSize uint64   // Data size of attribute value in bytes
}

/* Class attribute struct */
type Attribute struct {
	Header *AttributeHeader
	Value  []byte
}