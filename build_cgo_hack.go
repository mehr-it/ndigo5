// +build required

package ndigo5

// This file exists purely to prevent the golang toolchain from stripping
// away the c source directories and files when `go mod vendor` is used
// to populate a `vendor/` directory of a project depending on `goccy/go-graphviz`.
//
// How it works:
//  - every directory which only includes c source files receives a dummy.go file.
//  - every directory we want to preserve is included here as a _ import.
//  - this file is given a build to exclude it from the regular build.
import (
	// Prevent go tooling from stripping out the c source files.
	_ "github.com/mehr-it/ndigo5/include"
)
