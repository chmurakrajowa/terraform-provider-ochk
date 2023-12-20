#!/bin/bash

set -e

LINT_VERSION="v1.52.2"
LINT_BIN="$(go env GOPATH)/bin/golangci-lint"

echo "Download golangci-lint binary"
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin ${LINT_VERSION}

${LINT_BIN} --version
echo "Run golangci-lint"
${LINT_BIN} run ./ochk/... -v