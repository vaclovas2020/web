# Weblang language (Go package)

![Weblang](weblang.png "Weblang")

[![Go Reference](https://pkg.go.dev/badge/webimizer.dev/web.svg)](https://pkg.go.dev/webimizer.dev/web)

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/weblang)

Go-powered new programming language for web development

This is still very early version, so please use only for testing purposes, because it's not production ready yet.

## Go Example
```go
package main

import "webimizer.dev/web/cmd"

func main() {
	cmd.RegisterAndExecute()
}
```
## Weblang debian package

You can install debian weblang package and no need to use this go package directly anymore (this is more easy way):

1. Add GPG public key:
```sh
curl -fsSL https://weblang.dev/deb-repo/pgp-key.public | sudo gpg --dearmor -o /usr/share/keyrings/weblang.gpg
```

2. Add deb package to package list:
```sh
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/weblang.gpg] https://weblang.dev/deb-repo stable main" | sudo tee /etc/apt/sources.list.d/weblang.list  > /dev/null
```

3. Update apt list and install weblang package:
```sh
sudo apt update && sudo apt install weblang
```

4. Run application with command:
```sh
weblang run
```

Full example: https://github.com/vaclovas2020/weblang_app_example
