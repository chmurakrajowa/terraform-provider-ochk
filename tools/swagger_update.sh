#!/bin/bash

set -eo pipefail

output_file=ochk/sdk/swagger.json

echo "url:"${TF_VAR_host}
curl -s https://${TF_VAR_host}/v2/api-docs | jq '.' >${output_file}
echo "Updated swagger file: ${output_file}"
