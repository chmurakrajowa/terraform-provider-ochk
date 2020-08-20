// +build tools

package main

import (
	_ "github.com/bflad/tfproviderdocs"
	_ "github.com/bflad/tfproviderlint/cmd/tfproviderlint"
	_ "github.com/go-swagger/go-swagger"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
