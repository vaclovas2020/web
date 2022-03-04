/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* Class attributes struct of Weblang bytcode */
package attribute

/* Class attributes header struct */
type AttributeHeader struct {
	AttributeNameLength uint64 // Attribute name length (bytes)
	AttributeType       uint8  // attribute type: string, int, float, bool or objectReference
	AttributeValueSize  uint64 // Data size of attribute value in bytes
}

/* Class attribute struct */
type Attribute struct {
	Header *AttributeHeader
	Value  []byte
}
