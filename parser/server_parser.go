/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package parser

import (
	"errors"
	"strings"

	"webimizer.dev/web/base"
)

/* parse server class */
func (parser *Parser) parseServer(sourceCode string) error {
	var className string
	var objName string
	var class base.Class
	obj := &base.Object{Class: &class, Scope: 0, Stack: parser.Stack}
	serverExpFull := parser.compileRegExp(serverRegExpFull)
	if serverExpFull.MatchString(sourceCode) {
		serverExpStart := parser.compileRegExp(serverRegExpStart)
		classNameExp := parser.compileRegExp(regExpClassName)
		className = classNameExp.FindString(strings.Replace(serverExpStart.FindString(sourceCode), "server", "", 1))
		obj.Attributes = make(map[string]interface{})
		objName = className
		err := parser.parseServerParams(obj, sourceCode, className)
		if err != nil {
			return err
		}
	} else {
		return errors.New("incorrect server class definition syntax")
	}
	return parser.pushToMap(objName, className, &class, obj)
}

/* check if it is server class type */
func (parser *Parser) isServerClass(sourceCode string) bool {
	serverExpStart := parser.compileRegExp(serverRegExpStart)
	return serverExpStart.MatchString(sourceCode)
}
