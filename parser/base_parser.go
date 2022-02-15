/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package parser

import (
	"regexp"

	"webimizer.dev/web/base"
	"webimizer.dev/web/server"
)

/* Weblang language syntax parser */
type Parser struct {
	Stack  *base.MemoryStack // Global MemoryStack on WebLang VM
	Server *server.Server    // Global server object
}

/* Parse from source code or bytecode and append result to class map */
func (parser *Parser) Parse(sourceCode string, sourceFileName string, byteCodeFileName string) error {
	return parser.parseSourceCode(sourceCode, sourceFileName, byteCodeFileName)
}

/* base funxc type to define diffrent parsers */
type parserFunc func(sourceCode string, isApplicable *bool) error

/* Parse source code and append result to class map */
func (parser *Parser) parseSourceCode(sourceCode string, sourceFileName string, byteCodeFileName string) error {
	var err error = nil
	err = parser.removeComments(&sourceCode)
	if err != nil {
		return err
	}
	err = parser.parserSourceCodeInternal([]parserFunc{parserFunc(parser.serverParser)}, sourceCode)
	if err != nil {
		return err
	}
	return nil
}

/* parsing source code using diffrent type parsers */
func (parser *Parser) parserSourceCodeInternal(parserFuncArray []parserFunc, sourceCode string) error {
	isApplicable := false
	var err error
	for _, p := range parserFuncArray {
		if !isApplicable {
			err = p(sourceCode, &isApplicable)
		} else {
			return err
		}
	}
	return nil
}

/* remove all comments (one line comments and multiple lines comments) from source code */
func (parser *Parser) removeComments(sourceCode *string) error {
	oneLineCommentsExp, err := parser.compileRegExp(regExpComments)
	if err != nil {
		return err
	}
	*sourceCode = oneLineCommentsExp.ReplaceAllString(*sourceCode, "")
	return nil
}

/* Compile regexp */
func (parser *Parser) compileRegExp(regExp string) (*regexp.Regexp, error) {
	serverExpStart, err := regexp.Compile(regExp)
	if err != nil {
		return nil, err
	}
	return serverExpStart, nil
}
