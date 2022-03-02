/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package parser

import (
	"fmt"
	"strconv"
	"strings"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class/attribute"
)

/* Parse server class parameters */
func (parser *Parser) parseServerParams(obj *base.Object, sourceCode string, className string) error {
	serverExpOneParam, err := parser.compileRegExp(serverRegExpOneParam)
	if err != nil {
		return err
	}
	if serverExpOneParam.MatchString(sourceCode) {
		oneParam := serverExpOneParam.FindString(sourceCode)
		paramName, newSourceCode, err := parser.parseServerParamName(obj, oneParam, sourceCode, className)
		if err != nil {
			return err
		}
		paramValue, err := parser.parseServerParamValue(obj, paramName, oneParam, sourceCode, className)
		if err != nil {
			return err
		}
		if paramName == "port" {
			intVar, err := strconv.Atoi(paramValue)
			if err != nil {
				return err
			}
			(*obj).Attributes[paramName] = int64(intVar)
			(*obj).AttributesType[paramName] = attribute.AttributeType_Int
		} else if paramName == "router" {
			(*obj).Attributes[paramName] = paramValue
			(*obj).AttributesType[paramName] = attribute.AttributeType_ObjReference
		} else {
			(*obj).Attributes[paramName] = paramValue
			(*obj).AttributesType[paramName] = attribute.AttributeType_String
		}
		if newSourceCode != "" {
			return parser.parseServerParams(obj, newSourceCode, className)
		}
	}
	return nil
}

/* parse server param name */
func (parser *Parser) parseServerParamName(obj *base.Object, oneParam string, sourceCode string, className string) (string, string, error) {
	var paramName, newSourceCode string
	newSourceCode = strings.Replace(sourceCode, oneParam, "", 1)
	paramNameExp, err := parser.compileRegExp(serverRegExpParamName)
	if err != nil {
		return paramName, newSourceCode, err
	}
	paramName = paramNameExp.FindString(oneParam)
	if _, found := (*obj).Attributes[paramName]; found {
		return paramName, newSourceCode, fmt.Errorf("class %v already has attribute %v defined", className, paramName)
	}
	return paramName, newSourceCode, nil
}

func (parser *Parser) parseServerParamValue(obj *base.Object, paramName string, oneParam string, sourceCode string, className string) (string, error) {
	var paramValue string
	paramValueFull := strings.Replace(oneParam, "@"+paramName, "", 1)
	paramValueStartExp, err := parser.compileRegExp(serverRegExpParamValueStart)
	if err != nil {
		return paramValue, err
	}
	paramValueEndExp, err := parser.compileRegExp(serverRegExpParamValueEnd)
	if err != nil {
		return paramValue, err
	}
	paramValue = strings.ReplaceAll(paramValueFull, paramValueStartExp.FindString(paramValueFull), "")
	paramValue = strings.ReplaceAll(paramValue, paramValueEndExp.FindString(paramValue), "")
	return paramValue, nil
}
