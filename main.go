package main

import (
	acli "github.com/aideneyre/acli/cmd"
)

// main invokes our root acli command located at cmd/root.go
func main() {
	acli.Execute()
}
