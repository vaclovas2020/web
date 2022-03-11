/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class"
)

/* Bytecode loader handler function */
type LoaderFunc func(class *base.Class, obj *base.Object, memory *base.MemoryMap) error

/* Close bytecode file */
func (loader *Loader) closeFile() error {
	if loader.file != nil {
		return loader.file.Close()
	}
	return nil
}

/* Run bytcode parser if bytecode file is valid */
func (loader *Loader) parseIfValid(isValid bool, classPtr *base.Class, memory *base.MemoryMap) (bool, error) {
	if isValid {
		var err error = nil
		classPtr.Type, err = loader.detectClassType(classPtr)
		if err != nil {
			return false, err
		}
		var objPtr *base.Object
		if classPtr.ByteCode.Header.ClassType == class.ClassType_Object || classPtr.ByteCode.Header.ClassType == class.ClassType_Model {
			objPtr = nil
		}
		objPtr = &base.Object{Class: classPtr, Memory: memory}
		return isValid, loader.runLoader([]LoaderFunc{
			LoaderFunc(loader.loadClassName),
			LoaderFunc(loader.loadClassAttributes),
		}, classPtr, objPtr, memory)
	}
	return isValid, nil
}
