#!/bin/bash

set -eo pipefail

output_file=ochk/api/swagger.json
output_file2=ochk/api/swagger2.json
TF_VAR_host=pckproxy.ochk.pilot
echo "url:"${TF_VAR_host}
curl -s https://${TF_VAR_host}/swagger/v1/swagger.json | jq '.' >${output_file}
curl -s https://${TF_VAR_host}/swagger/v1/swagger.json  >  ${output_file2}
echo "Updated swagger file: ${output_file}"