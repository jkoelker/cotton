// +build tools

// The toools package is used to track tools used by the devkit and allow
// dependabot to manage their versions in the go.mod file.

package tools

import (
	_ "github.com/daixiang0/gci"
	_ "gotest.tools/gotestsum"
	_ "mvdan.cc/gofumpt"
	_ "golang.org/x/tools/cmd/goimports"
)
