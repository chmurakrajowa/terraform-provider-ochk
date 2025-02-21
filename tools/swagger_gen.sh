#!/bin/bash

set -e

echo "Clear ochk/sdk/gen directory"
rm -rf ochk/api/v3/*

SWAGGER_BIN="$(go env GOPATH)/bin/swagger"

if [[ ! -f "${SWAGGER}" ]]; then
  echo "Install swagger"
  go get github.com/go-swagger/go-swagger/cmd/swagger
  go install github.com/go-swagger/go-swagger/cmd/swagger
fi

echo "Run swagger gen"
${SWAGGER_BIN} generate client -f ./ochk/api/swagger.json -t ./ochk/api/v3/ -A ochk
