/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"webimizer.dev/web/base"
)

/* Bytecode loader handler function */
type LoaderFunc func(class *base.Class, obj *base.Object) error

/* Close bytecode file */
func (loader *Loader) closeFile() error {
	if loader.file != nil {
		return loader.file.Close()
	}
	return nil
}

/* Run bytcode parser if bytecode file is valid */
func (loader *Loader) runIfValid(isValid bool, class *base.Class, obj *base.Object) (bool, error) {
	if isValid {
		return isValid, loader.runLoader([]LoaderFunc{
			LoaderFunc(loader.loadClassAttributes),
		}, class, obj)
	}
	return isValid, nil
}
