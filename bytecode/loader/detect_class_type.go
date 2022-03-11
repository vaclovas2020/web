/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package loader

import (
	"errors"

	"webimizer.dev/web/base"
	"webimizer.dev/web/bytecode/class"
)

/* Detect class type from bytecode */
func (loader *Loader) detectClassType(classPtr *base.Class) (string, error) {
	if classPtr.ByteCode == nil {
		return "", errors.New("bug detected: classPtr.ByteCode is nil")
	}
	if classPtr.ByteCode.Header == nil {
		return "", errors.New("bug detected: classPtr.ByteCode.Header is nil")
	}
	switch classPtr.ByteCode.Header.ClassType {
	case class.ClassType_Cms:
		return "cms", nil
	case class.ClassType_Controller:
		return "controller", nil
	case class.ClassType_Event:
		return "event", nil
	case class.ClassType_Firewall:
		return "firewall", nil
	case class.ClassType_Model:
		return "model", nil
	case class.ClassType_Object:
		return "object", nil
	case class.ClassType_Repository:
		return "repository", nil
	case class.ClassType_Router:
		return "router", nil
	case class.ClassType_Server:
		return "server", nil
	case class.ClassType_Service:
		return "service", nil
	case class.ClassType_View:
		return "view", nil
	}
	return "", errors.New("undefined class type")
}
