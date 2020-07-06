#!/bin/bash

set -ex

echo "Install golangci-lint"
go install github.com/golangci/golangci-lint/cmd/golangci-lint

echo "Run golangci-lint"
golangci-lint run ./ochk/... -v