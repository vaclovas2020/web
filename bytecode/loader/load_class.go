/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"webimizer.dev/web/base"
)

/* Bytecode loader handler function */
type LoaderFunc func(class *base.Class) error

/* Load and fully parse bytecode data to *base.Class */
func (loader *Loader) LoadClass(class *base.Class) error {
	isValid, err := loader.isValidByteCode()
	defer loader.closeFile()
	if err != nil {
		return err
	}
	return loader.runIfValid(isValid, class)
}

/* Close bytecode file */
func (loader *Loader) closeFile() error {
	if loader.file != nil {
		return loader.file.Close()
	}
	return nil
}

/* Run bytcode parser if bytecode file is valid */
func (loader *Loader) runIfValid(isValid bool, class *base.Class) error {
	if isValid {
		return loader.runLoader([]LoaderFunc{LoaderFunc(loader.loadClassHeader)}, class)
	}
	return nil
}
