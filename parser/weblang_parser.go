package parser

import "webimizer.dev/web/base"

type Parser struct {
	Classes *map[string]base.Class
}

func (parser *Parser) Parse(sourceCode string) error {
	// TODO: implement parser
	return nil
}
