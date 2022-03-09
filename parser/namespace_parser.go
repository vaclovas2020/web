/* Copyright (c) */

package parser

import "strings"

func (parser *Parser) parseNamespace(sourceCode *string, isApplicable *bool) error {
	nameSpaceFull, err := parser.compileRegExp(regExpNamespace)
	if err != nil {
		return err
	}
	namespaceFullStr := nameSpaceFull.FindString(*sourceCode)
	nameSpaceName, err := parser.compileRegExp(regExpNamespaceName)
	if err != nil {
		return err
	}
	parser.Namespace = nameSpaceName.FindString(namespaceFullStr)
	*sourceCode = strings.ReplaceAll(*sourceCode, namespaceFullStr, "")
	return nil
}
