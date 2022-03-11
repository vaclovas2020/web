/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"bytes"
	"encoding/binary"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class"
)

/* Load bytecode to class struct */
func (loader *Loader) loadClassHeader(classPtr *base.Class) error {
	data := make([]byte, class.HeaderSize)
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
	err = binary.Read(buf, binary.BigEndian, classPtr.ByteCode.Header)
	if err != nil {
		return err
	}
	return nil
}
