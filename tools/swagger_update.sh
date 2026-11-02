#!/bin/bash

set -eo pipefail

output_file=ochk/api/swagger.json
output_file2=ochk/api/swagger2.json
output_file_yaml=ochk/api/swagger.yaml

TF_VAR_host=pckproxy.ochk.pilot
echo "url:"${TF_VAR_host}
curl -s -v https://${TF_VAR_host}/swagger/v1/swagger.json | jq '.' >${output_file}
curl -s -v https://${TF_VAR_host}/swagger/v1/swagger.json  >  ${output_file2}
curl -s -v https://${TF_VAR_host}/swagger/v1/swagger.yaml  >  ${output_file_yaml}
#curl -s -v http://localhost:8060/swagger/v1/swagger.json | jq '.' >${output_file}
#curl -s -v http://localhost:8060/swagger/v1/swagger.json  >  ${output_file2}
echo "Updated swagger file: ${output_file}"
echo "Updated swagger yaml file: ${output_file_yaml}"