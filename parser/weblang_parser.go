/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package parser

import (
	"fmt"
	"log"
	"regexp"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class"
)

/* Weblang language syntax parser */
type Parser struct {
	Stack *base.MemoryStack // Global MemoryStack on WebLang VM
}

const regExpClassName string = "[[:alpha:]]\\w*"

/* Parse source code and append result to class map */
func (parser *Parser) Parse(sourceCode string) error {
	if parser.isServerClass(sourceCode) {
		err := parser.parseServer(sourceCode)
		if err != nil {
			return err
		}
	}
	return nil
}

/* Push class to MemoryStack */
func (parser *Parser) pushToMap(objName string, className string, classPtr *base.Class, obj *base.Object) error {
	if _, found := (*parser.Stack).Classes[className]; !found {
		(*parser.Stack).Classes[className] = *classPtr
		log.Printf("\033[32m[weblang]\033[0m Loaded class '%v' to VM environement successfully", className)
	} else {
		return fmt.Errorf("class with name '%v' already exists", className)
	}
	if (*classPtr).ByteCode != nil && ((*classPtr).ByteCode.Header.ClassType == class.ClassType_Object ||
		(*classPtr).ByteCode.Header.ClassType == class.ClassType_Model) {
		return nil // if class type is model or object than no need to add obj to MemoryStack therefore we return and exit function
	}
	if o, found := (*parser.Stack).Objects[objName]; !found || o.Scope > 0 {
		(*parser.Stack).Objects[objName] = *obj
		log.Printf("\033[32m[weblang]\033[0m Loaded %v object '%v' to VM environement successfully: %v", className, objName, (*obj).Attributes)
		return nil
	}
	return fmt.Errorf("object with name '%v' already exists", objName)
}

/* Compile regexp */
func (parser *Parser) compileRegExp(regExp string) *regexp.Regexp {
	serverExpStart, err := regexp.Compile(regExp)
	if err != nil {
		panic(err)
	}
	return serverExpStart
}
