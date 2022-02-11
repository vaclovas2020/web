package parser

import (
	"fmt"
	"log"
	"regexp"

	"webimizer.dev/web/base"
)

/* Weblang language syntax parser */
type Parser struct {
	Stack *base.MemoryStack
}

const regExpClassName string = "[[:alpha:]]\\w*"
const serverRegExpStart string = "^(server)\\s+" + regExpClassName + "\\s*[{]\\s*"
const serverRegExpParamName string = "(router|port|host)"
const serverRegExpParamValueStart string = "[(]\\s*[\"]*"
const serverRegExpParamValue string = "(\\w|[.])+"
const serverRegExpParamValueEnd string = "[\"]*\\s*[)]\\s+"
const serverRegExpOneParam string = "([@]" + serverRegExpParamName + serverRegExpParamValueStart + serverRegExpParamValue + serverRegExpParamValueEnd + ")"
const serverRegExpParams string = serverRegExpOneParam + "{3}"
const serverRegExpEnd string = "[}]\\s*$"
const serverRegExpFull string = serverRegExpStart + serverRegExpParams + serverRegExpEnd

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
func (parser *Parser) pushToMap(className string, class *base.Class) error {
	if _, found := (*parser.Stack).Classes[className]; !found {
		(*parser.Stack).Classes[className] = *class
		log.Printf("\033[32m[weblang]\033[0m Loaded class '%v' to VM environement successfully", className)
		return nil
	}
	return fmt.Errorf("class with name '%v' already exists", className)
}

/* Compile regexp */
func (parser *Parser) compileRegExp(regExp string) *regexp.Regexp {
	serverExpStart, err := regexp.Compile(regExp)
	if err != nil {
		panic(err)
	}
	return serverExpStart
}
