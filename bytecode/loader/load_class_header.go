/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode"
	"webimizer.dev/web/bytecode/class"
)

/* Load bytecode to class struct */
func (loader *Loader) loadClassHeader(classPtr *base.Class) error {
	if classPtr == nil {
		return errors.New("loadClassHeader: classPtr is nil")
	}
	classPtr.ByteCode = &bytecode.ByteCode{Header: &class.ClassHeader{}}
	data, err := loader.readData(class.HeaderSize)
	if err != nil {
		return fmt.Errorf("loadClassHeader: loader.readData %v", err.Error())
	}
	buf := &bytes.Buffer{}
	_, err = buf.Write(data)
	if err != nil {
		return fmt.Errorf("loadClassHeader: buf.Write(data) %v", err.Error())
	}
	err = binary.Read(buf, binary.BigEndian, classPtr.ByteCode.Header)
	if err != nil {
		return fmt.Errorf("loadClassHeader: binary.Read %v", err.Error())
	}
	return nil
}
