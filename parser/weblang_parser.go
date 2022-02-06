package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"webimizer.dev/web/base"
)

/* Weblang language syntax parser */
type Parser struct {
	Classes *map[string]base.Class
}

const regExpClassName string = "[[:alpha:]]\\w*"
const serverRegExpStart string = "^(server)\\s+" + regExpClassName + "\\s*[{]\\s*"
const serverRegExpParamName = "(router|port|host)"
const serverRegExpOneParam string = "([@]" + serverRegExpParamName + "[(]\\s+[\"]*(\\w|[.])+[\"]*\\s+[)]\\s+)"
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

func (parser *Parser) pushToMap(className string, class *base.Class) error {
	if _, found := (*parser.Classes)[className]; !found {
		(*parser.Classes)[className] = *class
		return nil
	}
	return fmt.Errorf("class with name '%v' already exists", className)
}

func (parser *Parser) parseServerParams(class *base.Class, sourceCode string) {
	serverExpOneParam := parser.compileRegExp(serverRegExpOneParam)
	if serverExpOneParam.MatchString(sourceCode) {
		oneParam := serverExpOneParam.FindString(sourceCode)
		newSourceCode := strings.Replace(sourceCode, oneParam, "", 1)
		// TODO: parse param name and value
		if newSourceCode != "" {
			parser.parseServerParams(class, newSourceCode)
		}
	}
}

func (parser *Parser) parseServer(sourceCode string) error {
	var className string
	var class base.Class
	serverExpFull := parser.compileRegExp(serverRegExpFull)
	if serverExpFull.MatchString(sourceCode) {
		serverExpStart := parser.compileRegExp(serverRegExpStart)
		classNameExp := parser.compileRegExp(regExpClassName)
		className = classNameExp.FindString(strings.Replace(serverExpStart.FindString(sourceCode), "server", "", 1))
		parser.parseServerParams(&class, sourceCode)
	} else {
		return errors.New("incorrect class definition syntax")
	}
	return parser.pushToMap(className, &class)
}

func (parser *Parser) compileRegExp(regExp string) *regexp.Regexp {
	serverExpStart, err := regexp.Compile(regExp)
	if err != nil {
		panic(err)
	}
	return serverExpStart
}

func (parser *Parser) isServerClass(sourceCode string) bool {
	serverExpStart := parser.compileRegExp(serverRegExpStart)
	return serverExpStart.MatchString(sourceCode)
}
