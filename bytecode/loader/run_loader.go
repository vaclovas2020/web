/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import "webimizer.dev/web/base"

/* Run loader handlers and parse bytecode to class struct */
func (loader *Loader) runLoader(handlers []LoaderFunc, class *base.Class, obj *base.Object) error {
	for _, handler := range handlers {
		err := handler(class, obj)
		if err != nil {
			return err
		}
	}
	return nil
}
