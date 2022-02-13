# Weblang language (Go package)

![Weblang](weblang.png "Weblang")

[![Go Reference](https://pkg.go.dev/badge/webimizer.dev/web.svg)](https://pkg.go.dev/webimizer.dev/web)

Go-powered new programming language for web development

This is still very early version, so please use only for testing purposes, because it's not production ready yet.

## Example
```go
package main

import (
	"context"

	"webimizer.dev/web/controller"
	"webimizer.dev/web"
)

func main() {
	vm := web.VM{}
	vm.InitVM("web/src", "web/generated")
	vm.DefineFunc("MainController", "index", controller.Handler(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Welcome from weblang!")
	}))
}
```

Full example: https://github.com/vaclovas2020/webtest
