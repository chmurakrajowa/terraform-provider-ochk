#!/bin/bash

set -eo pipefail

output_file=ochk/sdk/swagger.json

curl --insecure https://iaas-api-proxy.ochk.pilot/v2/api-docs | jq . >${output_file}

echo "Updated swagger file: ${output_file}"
