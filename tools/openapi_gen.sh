#!/bin/bash

set -e

echo "Clear ochk/sdk/gen directory"
rm -rf ochk/api/v3/*

SWAGGER_BIN="$(go env GOPATH)/bin/swagger"
OPENAPI_BIN="$(go env GOPATH)/bin/oapi-codegen"

if [[ ! -f "${SWAGGER}" ]]; then
  echo "Install swagger"
#  go get github.com/go-swagger/go-swagger/cmd/swagger
  #go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
  go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
fi

echo "Run openapi gen"

#${OPENAPI_BIN} -package=ochk_client -generate=types,client,spec -o=ochk/client/ochk_client.go ./ochk/api/swagger.json

#${SWAGGER_BIN} generate client -f ./ochk/api/swagger.json -t ./ochk/api/v3/ -A ochk


openapi-generator generate -i ./ochk/api/swagger.json -g go -o ./ochk/api/openapi/v3