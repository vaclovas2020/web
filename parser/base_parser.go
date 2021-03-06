/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

/* Weblang language parser */
package parser

import (
	"regexp"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/generator"
	"webimizer.dev/web/bytecode/loader"
	"webimizer.dev/web/core/server"
)

/* Weblang language syntax parser */
type Parser struct {
	Memory    *base.MemoryMap             // Global MemoryMap on WebLang VM
	Server    *server.Server              // Global server object
	Namespace string                      // class namespace (can be empty)
	generator generator.ByteCodeGenerator // ByteCodeGenerator for this class
	loader    loader.Loader               // Bytecodde loader
}

/* Parse from source code or bytecode and append result to class map */
func (parser *Parser) Parse(sourceCode string, sourceFileName string, byteCodeFileName string) error {
	return parser.parseSourceCode(sourceCode, sourceFileName, byteCodeFileName)
}

/* base function type to define diffrent parsers */
type parserFunc func(sourceCode *string, isApplicable *bool) error

/* Parse source code and append result to class map */
func (parser *Parser) parseSourceCode(sourceCode string, sourceFileName string, byteCodeFileName string) error {
	var err error = nil
	parser.loader = loader.Loader{SourceCodeFileName: sourceFileName, ByteCodeFileName: byteCodeFileName, Server: parser.Server}
	success, err := parser.loader.LoadClassAndObject(parser.Memory)
	if err != nil {
		return err
	}
	if success {
		return nil
	}
	parser.generator = generator.ByteCodeGenerator{ByteCodeFileName: byteCodeFileName, SourceCodeFileName: sourceFileName}
	err = parser.removeComments(&sourceCode)
	if err != nil {
		return err
	}
	err = parser.parserSourceCodeInternal([]parserFunc{
		parserFunc(parser.parseNamespace),
		parserFunc(parser.serverParser),
	}, sourceCode)
	if err != nil {
		return err
	}
	return nil
}

/* parsing source code using diffrent type parsers */
func (parser *Parser) parserSourceCodeInternal(parserFuncArray []parserFunc, sourceCode string) error {
	isApplicable := false
	var err error = nil
	for _, p := range parserFuncArray {
		if !isApplicable {
			err = p(&sourceCode, &isApplicable)
		}
	}
	if err != nil {
		return err
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
