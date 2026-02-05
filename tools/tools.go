//go:build tools

package tools

import (
	_ "github.com/mattn/goveralls"
	_ "github.com/modocache/gover"
	_ "golang.org/x/tools/cmd/cover"

    // Also include any other library dependencies that are only used for testing
	_ "golang.org/x/net/html/charset"
	_ "golang.org/x/crypto/ssh/terminal"
)
