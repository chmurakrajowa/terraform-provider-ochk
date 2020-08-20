#!/bin/bash

set -ex

echo "Install golangci-lint"
go install github.com/golangci/golangci-lint/cmd/golangci-lint

echo "Run golangci-lint"
$(go env GOPATH)/bin/golangci-lint run ./ochk/... -v