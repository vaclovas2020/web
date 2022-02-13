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

/* Parse source code and append result to class map */
func (parser *Parser) parseSourceCode(sourceCode string, sourceFileName string, byteCodeFileName string) error {
	if parser.isServerClass(sourceCode) {
		err := parser.parseServer(sourceCode)
		if err != nil {
			return err
		}
	}
	return nil
}

/* Compile regexp */
func (parser *Parser) compileRegExp(regExp string) *regexp.Regexp {
	serverExpStart, err := regexp.Compile(regExp)
	if err != nil {
		panic(err)
	}
	return serverExpStart
}
