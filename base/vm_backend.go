/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* Weblang VM environment base struct types such as Function, MmemoryMap, Class and Object. */
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

/* Memory of all declared classes in VM */
type MemoryMap struct {
	Classes map[string]Class  // Class list
	Objects map[string]Object // All objects list
}

/* Weblang class represantation in VM */
type Class struct {
	Type     string              // Class type (string representation)
	Methods  map[string]Function // Function list
	ByteCode *bytecode.ByteCode  // Bytecode pointer
}

/* Object struct in VM environment */
type Object struct {
	Scope      uint                   // Object scope, 0 - global, 1 and larger - local
	Class      *Class                 // pointer to class struct
	Attributes map[string]interface{} // Attributes list
	Memory     *MemoryMap             // list of all declared global classes & objects in VM environment
}

/* User defined function handler */
type FunctionHandler interface {
	Invoke(args map[string]interface{}, funcPtr *Function, callerObj *Object, receiverObj *Object) error // Invoke function execution
}
