# web
Go-powered new programming language for web development

This is still very early version, so please use only for testing purposes, because it's not production ready yet.

## Example
```go
package main

import (
	"context"

	"webimizer.dev/web"
)

func main() {
	vm := web.VM{}
	vm.InitVM(context.Background(), []string{}, "test")
}
```