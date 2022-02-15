/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package parser

import (
	"errors"
	"fmt"
	"log"

	"webimizer.dev/web/base"
)

/* Push class to MemoryStack */
func (parser *Parser) pushToMap(objName string, className string, classPtr *base.Class, obj *base.Object) error {
	if _, found := (*parser.Stack).Classes[className]; !found {
		(*parser.Stack).Classes[className] = *classPtr
		log.Printf("\033[32m[weblang]\033[0m Loaded class '%v' to VM environment successfully", className)
	} else {
		return fmt.Errorf("class with name '%v' already exists", className)
	}
	if (*parser.Server).ServerObject != nil && (*classPtr).Type == "server" {
		return errors.New("server type class already exists")
	}
	if (*parser.Server).RouterObject != nil && (*classPtr).Type == "router" {
		return errors.New("router type class already exists")
	}
	if (*classPtr).Type == "server" {
		(*parser.Server).ServerObject = obj
		(*parser.Server).Host = fmt.Sprintf("%v", obj.Attributes["host"])
		(*parser.Server).Port = obj.Attributes["port"].(int)
	}
	if (*classPtr).Type == "router" {
		(*parser.Server).RouterObject = obj
	}
	if (*classPtr).Type == "object" || (*classPtr).Type == "model" {
		return nil // if class type is model or object than no need to add obj to MemoryStack therefore we return and exit function
	}
	if o, found := (*parser.Stack).Objects[objName]; !found || o.Scope > 0 {
		(*parser.Stack).Objects[objName] = *obj
		log.Printf("\033[32m[weblang]\033[0m Loaded %v object '%v' to VM environment successfully: %v", className, objName, (*obj).Attributes)
		return nil
	}
	return fmt.Errorf("object with name '%v' already exists", objName)
}
