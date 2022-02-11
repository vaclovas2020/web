package base

import (
	"webimizer.dev/web/bytecode"
	"webimizer.dev/web/bytecode/class/method"
)

/* Web function representation in VM */
type Function struct {
	Args        map[string]interface{} // arguments list
	Handler     FunctionHandler        // FunctionHandler interface for function invoke
	ClassMethod *method.ClassMethod    // ClassMethod struct from bytecode file
}

/* Stack of all declared classes in VM */
type MemoryStack struct {
	Classes map[string]Class  // Class list
	Objects map[string]Object // All objects list
}

/* Weblang class represantation in VM */
type Class struct {
	Methods  map[string]Function // Function list
	ByteCode *bytecode.ByteCode  // Bytecode pointer
}

/* Object struct in VM environment */
type Object struct {
	Scope      uint                   // Object scope, 0 - global, 1 and larger - local
	Class      *Class                 // pointer to class struct
	Attributes map[string]interface{} // Attributes list
	Stack      *MemoryStack           // list of all declared global classes & objects in VM environment
}

/* User defined function handler */
type FunctionHandler interface {
	Invoke(args map[string]interface{}, funcPtr *Function, obj *Object) error // Invoke function execution
}
