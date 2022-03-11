/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"encoding/binary"

	"webimizer.dev/web/base"
)

/* Load bytecode to class struct */
func (loader *Loader) loadClassHeader(class *base.Class) error {
	err := binary.Read(loader.file, binary.BigEndian, &class.ByteCode.Header)
	if err != nil {
		return err
	}
	return nil
}
