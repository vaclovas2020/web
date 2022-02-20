package parser

import (
	"fmt"

	"webimizer.dev/web/base"
)

func (parser *Parser) validateServerParams(objName string, obj *base.Object) error {
	if _, existPort := obj.Attributes["port"]; !existPort {
		return fmt.Errorf("server object '%v' hasn't required parameter port", objName)
	}
	if _, existHost := obj.Attributes["host"]; !existHost {
		return fmt.Errorf("server object '%v' hasn't required parameter host", objName)
	}
	if _, existHost := obj.Attributes["router"]; !existHost {
		return fmt.Errorf("server object '%v' hasn't required parameter router", objName)
	}
	return nil
}
