package loader

import (
	"fmt"

	"webimizer.dev/web/base"
)

func (loader *Loader) updateServerParams(obj *base.Object) {
	if obj != nil && obj.Class != nil && obj.Class.Type == "server" {
		(*loader.Server).ServerObject = obj
		(*loader.Server).Host = fmt.Sprintf("%v", obj.Attributes["host"])
		(*loader.Server).Port = int(obj.Attributes["port"].(int64))
		if v, exists := obj.Attributes["staticFiles"]; exists {
			(*loader.Server).StaticFilesPath = fmt.Sprintf("%v", v)
		}
	}
}
