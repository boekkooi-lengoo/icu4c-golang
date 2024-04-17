//go:build tools

package tools

// Adding new CLI tools to the project

// 1. Change directory to .tools
// 2. Run go get with the command to be installed
// 3. Add a new line in the import below with `_ "package/to/import"`
// 4. Run go mod vendor
// 5. Run go build -o bin/<name_of_package> vendor/package/main.go
// 6. You can now use the package CLI

// It is encouraged to add an abstraction for the steps 4 and 5 in the Makefile.
// Try to follow the same approach implemented for $(swag)

import (
	_ "github.com/xlab/c-for-go"
)
