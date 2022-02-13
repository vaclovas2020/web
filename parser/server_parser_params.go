/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package parser

import (
	"fmt"
	"strconv"
	"strings"

	"webimizer.dev/web/base"
)

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
