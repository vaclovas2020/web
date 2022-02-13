/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"webimizer.dev/web/base"
)

const serverRegExpStart string = "^(server)\\s+" + regExpClassName + "\\s*[{]\\s*"
const serverRegExpParamName string = "(router|port|host)"
const serverRegExpParamValueStart string = "[(]\\s*[\"]*"
const serverRegExpParamValue string = "(\\w|[.])+"
const serverRegExpParamValueEnd string = "[\"]*\\s*[)]\\s+"
const serverRegExpOneParam string = "([@]" + serverRegExpParamName + serverRegExpParamValueStart + serverRegExpParamValue + serverRegExpParamValueEnd + ")"
const serverRegExpParams string = serverRegExpOneParam + "{3}"
const serverRegExpEnd string = "[}]\\s*$"
const serverRegExpFull string = serverRegExpStart + serverRegExpParams + serverRegExpEnd

/* Parse server class parameters */
func (parser *Parser) parseServerParams(obj *base.Object, sourceCode string, className string) error {
	serverExpOneParam := parser.compileRegExp(serverRegExpOneParam)
	if serverExpOneParam.MatchString(sourceCode) {
		oneParam := serverExpOneParam.FindString(sourceCode)
		newSourceCode := strings.Replace(sourceCode, oneParam, "", 1)
		paramNameExp := parser.compileRegExp(serverRegExpParamName)
		paramName := paramNameExp.FindString(oneParam)
		if _, found := (*obj).Attributes[paramName]; found {
			return fmt.Errorf("class %v already has attribute %v defined", className, paramName)
		}
		paramValueFull := strings.Replace(oneParam, "@"+paramName, "", 1)
		paramValueStartExp := parser.compileRegExp(serverRegExpParamValueStart)
		paramValueEndExp := parser.compileRegExp(serverRegExpParamValueEnd)
		paramValue := strings.ReplaceAll(paramValueFull, paramValueStartExp.FindString(paramValueFull), "")
		paramValue = strings.ReplaceAll(paramValue, paramValueEndExp.FindString(paramValue), "")
		if paramName == "port" {
			intVar, err := strconv.Atoi(paramValue)
			if err != nil {
				return err
			}
			(*obj).Attributes[paramName] = intVar
		} else {
			(*obj).Attributes[paramName] = paramValue
		}
		if newSourceCode != "" {
			return parser.parseServerParams(obj, newSourceCode, className)
		}
	}
	return nil
}

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
