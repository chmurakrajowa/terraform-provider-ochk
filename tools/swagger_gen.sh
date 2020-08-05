#!/bin/bash

set -ex

echo "Clear ochk/sdk/gen directory"
rm -rf ochk/sdk/gen/*

echo "Install swagger"
go install github.com/go-swagger/go-swagger/cmd/swagger

echo "Run swagger gen"
"$(go env GOPATH)/bin/swagger" generate client -f ./ochk/sdk/swagger.json -t ./ochk/sdk/gen/ -A ochk