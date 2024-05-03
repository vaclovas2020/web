package loader

import (
	"errors"

	"webimizer.dev/web/base"
)

/* Detect bytecode reader bugs */
func (loader *Loader) detectBug(classPtr *base.Class, _ *base.Object, _ *base.MemoryMap) error {
	if classPtr == nil {
		return errors.New("bug detected: classPtr is nil")
	}
	if classPtr.ByteCode == nil {
		return errors.New("bug detected: classPtr.ByteCode is nil")
	}
	if classPtr.ByteCode.Header == nil {
		return errors.New("bug detected: classPtr.ByteCode.Header is nil")
	}
	return nil
}
