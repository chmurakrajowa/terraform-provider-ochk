#!/bin/bash

set -e

LINT_BIN="$(go env GOPATH)/bin/golangci-lint"

if [[ ! -f ${LINT_BIN} ]]; then
  echo "Install golangci-lint"
  go install github.com/golangci/golangci-lint/cmd/golangci-lint
fi

${LINT_BIN} --version
echo "Run golangci-lint"
${LINT_BIN} run ./ochk/... -v