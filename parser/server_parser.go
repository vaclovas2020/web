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
	class := &base.Class{Type: "server"}
	obj := &base.Object{Class: class, Scope: 0, Stack: parser.Stack}
	serverExpFull, err := parser.compileRegExp(serverRegExpFull)
	if err != nil {
		return err
	}
	if serverExpFull.MatchString(sourceCode) {
		className, err = parser.getServerClassName(sourceCode)
		if err != nil {
			return err
		}
		obj.Attributes = make(map[string]interface{})
		objName = className
		err := parser.parseServerParams(obj, sourceCode, className)
		if err != nil {
			return err
		}
	} else {
		return errors.New("incorrect server class definition syntax")
	}
	return parser.pushToMap(objName, className, class, obj)
}

/* get server class namefrom source code  */
func (parser *Parser) getServerClassName(sourceCode string) (string, error) {
	var className string
	serverExpStart, err := parser.compileRegExp(serverRegExpStart)
	if err != nil {
		return className, err
	}
	classNameExp, err := parser.compileRegExp(regExpClassName)
	if err != nil {
		return className, err
	}
	className = classNameExp.FindString(strings.Replace(serverExpStart.FindString(sourceCode), "server", "", 1))
	return className, nil
}

/* check if it is server class type */
func (parser *Parser) isServerClass(sourceCode string) (bool, error) {
	serverExpStart, err := parser.compileRegExp(serverRegExpStart)
	if err != nil {
		return false, err
	}
	return serverExpStart.MatchString(sourceCode), nil
}
