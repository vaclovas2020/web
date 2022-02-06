package parser

import (
	"fmt"
	"log"
	"regexp"

	"webimizer.dev/web/base"
)

/* Weblang language syntax parser */
type Parser struct {
	Classes *map[string]base.Class
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

func (parser *Parser) pushToMap(className string, class *base.Class) error {
	if _, found := (*parser.Classes)[className]; !found {
		(*parser.Classes)[className] = *class
		log.Printf("\033[32m[weblang]\033[0m Loaded %v class '%v' successfully: %v", (*class).Type, className, *class)
		return nil
	}
	return fmt.Errorf("%v class with name '%v' already exists", (*class).Type, className)
}

func (parser *Parser) compileRegExp(regExp string) *regexp.Regexp {
	serverExpStart, err := regexp.Compile(regExp)
	if err != nil {
		panic(err)
	}
	return serverExpStart
}
