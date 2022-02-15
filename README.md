# Weblang language (Go package)

![Weblang](weblang.png "Weblang")

[![Go Reference](https://pkg.go.dev/badge/webimizer.dev/web.svg)](https://pkg.go.dev/webimizer.dev/web)

Go-powered new programming language for web development

This is still very early version, so please use only for testing purposes, because it's not production ready yet.

## Example
```go
package main

import (
	"log"

	"webimizer.dev/web"
)

func main() {
	vm := web.VM{}
	vm.InitVM("web/src", "web/generated")
	log.Fatal(vm.StartServer())
}
```

Full example: https://github.com/vaclovas2020/webtest
