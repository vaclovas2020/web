package handler

import (
	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode"
	"webimizer.dev/web/bytecode/class/method"
)

/* Struct for calling class method and invoke bytecode execution in VM environemt */
type ClassMethodHandler struct {
	ClasMethod *method.ClassMethod
	ByteCode   *bytecode.ByteCode
}

/* Execute method bytecode in VM environemt */
func (method *ClassMethodHandler) Invoke(args map[string]interface{}, funcPtr *base.Function) error {
	// TODO: implement bytecode parser and start bytecode execution
	return nil
}
