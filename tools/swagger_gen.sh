#!/bin/bash

set -e

echo "Clear ochk/sdk/gen directory"
rm -rf ochk/sdk/gen/*

SWAGGER_BIN="$(go env GOPATH)/bin/swagger"

if [[ ! -f "${SWAGGER}" ]]; then
  echo "Install swagger"
  go install github.com/go-swagger/go-swagger/cmd/swagger
fi

echo "Run swagger gen"
${SWAGGER_BIN} generate client -f ./ochk/sdk/swagger.json -t ./ochk/sdk/gen/ -A ochk
