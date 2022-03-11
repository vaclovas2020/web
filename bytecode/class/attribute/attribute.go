/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* Class attributes struct of Weblang bytcode */
package attribute

/* Size of one AttributeHeader */
const AttributeHeaderSize int64 = 17

/* Class attributes header struct */
type AttributeHeader struct {
	AttributeNameLength int64 // Attribute name length (bytes)
	AttributeType       uint8 // attribute type: string, int, float or objectReference
	AttributeValueSize  int64 // Data size of attribute value in bytes
}

/* Class attribute struct */
type Attribute struct {
	Header *AttributeHeader
	Value  []byte
}
