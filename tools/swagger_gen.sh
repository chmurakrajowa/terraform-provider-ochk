#!/bin/bash

set -ex

echo "Install swagger"
go install github.com/go-swagger/go-swagger/cmd/swagger

echo "Run swagger for doctor"
swagger generate client -f ./ochk/sdk/swagger.yaml -t ./ochk/sdk/gen/ -A ochk