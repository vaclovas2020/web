/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package handler

import (
	"webimizer.dev/web/base"
)

/* Struct for calling class method and invoke bytecode execution in VM environment */
type ClassMethodHandler struct{}

/* Execute method bytecode in VM environemt */
func (method *ClassMethodHandler) Invoke(args map[string]interface{}, funcPtr *base.Function, obj *base.Object) error {
	// TODO: implement bytecode parser and start bytecode execution
	return nil
}
