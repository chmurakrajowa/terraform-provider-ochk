#!/bin/bash

set -ex

echo "Install swagger"
go install github.com/go-swagger/go-swagger/cmd/swagger

echo "Run swagger gen"
"$(go env GOPATH)/bin/swagger" generate client -f ./ochk/sdk/swagger.json -t ./ochk/sdk/gen/ -A ochk