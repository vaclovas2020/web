package base

import (
	"webimizer.dev/web/bytecode"
	"webimizer.dev/web/bytecode/class/method"
)

/* Web function representation in VM */
type Function struct {
	Args        map[string]interface{} // arguments list
	Handler     FunctionHandler
	ClassMethod *method.ClassMethod // bytecode ClassMethod pointer
	Class       *Class              // pointer to class struct
	Stack       *MemoryStack        // list of all declared global classes & objects in VM environment
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
	Object   *Object             //Pointer to class global object. For static objects only, etc: server, router, controller
	Stack    *MemoryStack        // list of all declared global classes & objects in VM environment
}

/* Object struct in VM environment */
type Object struct {
	Scope      uint                   // Object scope, 0 - global, 1 and larger - local
	Class      *Class                 // pointer to class struct
	Attributes map[string]interface{} // Attributes list
}

/* User defined function handler */
type FunctionHandler interface {
	Invoke(args map[string]interface{}, funcPtr *Function) error // Invoke function execution
}
