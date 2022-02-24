/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package generator

import "webimizer.dev/web/bytecode/class"

/* Generate class type (use correct uint8 constant) */
func (generator *ByteCodeGenerator) generateClassType() error {
	switch generator.Class.Type {
	case "server":
		generator.Class.ByteCode.Header.ClassType = class.ClassType_Server
	case "router":
		generator.Class.ByteCode.Header.ClassType = class.ClassType_Router
	case "controller":
		generator.Class.ByteCode.Header.ClassType = class.ClassType_Controller
	case "model":
		generator.Class.ByteCode.Header.ClassType = class.ClassType_Model
	case "view":
		generator.Class.ByteCode.Header.ClassType = class.ClassType_View
	case "object":
		generator.Class.ByteCode.Header.ClassType = class.ClassType_Object
	case "cms":
		generator.Class.ByteCode.Header.ClassType = class.ClassType_Cms
	case "firewall":
		generator.Class.ByteCode.Header.ClassType = class.ClassType_Firewall
	case "repository":
		generator.Class.ByteCode.Header.ClassType = class.ClassType_Repository
	case "service":
		generator.Class.ByteCode.Header.ClassType = class.ClassType_Service
	case "event":
		generator.Class.ByteCode.Header.ClassType = class.ClassType_Event
	}
	return nil
}
