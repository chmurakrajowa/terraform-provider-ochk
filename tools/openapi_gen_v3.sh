#!/bin/bash

set -e

echo "Clear ochk/sdk/gen directory"
rm -rf ochk/api/v3/*

SWAGGER_BIN="$(go env GOPATH)/bin/swagger"
OPENAPI_BIN="$(go env GOPATH)/bin/oapi-codegen"

if [[ ! -f "${SWAGGER}" ]]; then
  echo "Install swagger"
  go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
fi

echo "Run openapi gen"

#${OPENAPI_BIN} -package=ochk_client -generate=types,client,spec -o=ochk/client/ochk_client.go ./ochk/api/swagger.json

#${SWAGGER_BIN} generate client -f ./ochk/api/swagger.json -t ./ochk/api/v3/ -A ochk


# Prepare swagger.json file to generate client and model i go


cat ./ochk/api/swagger.json

sed -i -e 's/hasPrivateKey/privateKeyPresent/g' ./ochk/api/swagger.json


awk '
BEGIN { count = 0 }
{
    while (match($0, /"hasValue"/)) {
        count++
        if (count == 1)
            $0 = substr($0, 1, RSTART-1) "\"valuePresent\"" substr($0, RSTART+RLENGTH)
        else if (count == 2)
            $0 = substr($0, 1, RSTART-1) "\"valueAvailable\"" substr($0, RSTART+RLENGTH)
        else
            break
    }
    print
}
' ./ochk/api/swagger.json > ./ochk/api/temp.json && mv ./ochk/api/temp.json ./ochk/api/swagger.json


sed -i -e 's/hasValue/valueExists/g' ./ochk/api/swagger.json

sed -i -e 's/hasDefaultValue/defaultValuePresent/g' ./ochk/api/swagger.json



sed -i -e 's/"format": "date-time",//g' ./ochk/api/swagger.json

sed -i '' '/"startDate": {/,/}/{
  /"format": "date-time"/d
  s/"type": "string",/"type": "string"/
}' ./ochk/api/swagger.json



openapi-generator generate -i ./ochk/api/swagger.json -g go -o ./ochk/api/v3 \
    --additional-properties=enumClassPrefix=true,useOneOfDiscriminatorLookup=true,generateMarshalJSON=false

rm -rf ./ochk/api/v3/.openapi-generator
rm -rf ./ochk/api/v3/api
rm -rf ./ochk/api/v3/test
rm -rf ./ochk/api/v3/docs
rm ./ochk/api/v3/.travis.yml
rm ./ochk/api/v3/.gitignore
rm ./ochk/api/v3/.openapi-generator-ignore
rm ./ochk/api/v3/git_push.sh
rm ./ochk/api/v3/README.md
rm -rf ./ochk/api/v3/go.mod
rm ./ochk/api/v3/go.sum


# na potrzeby rezerwacji allocation public address ip: musimy wyslac serviceList: []
sed -i -e 's/serviceList,omitempty/serviceList/g' ochk/api/v3/model_public_ip_allocation.go

# na potrzeby modyfikacji viertul machine  musimy wyslac dać mozliwość wysyłki   "tags": [],  w obiekcie virtual machine

sed -i -e 's/tags,omitempty/tags/g' ochk/api/v3/model_virtual_machine_instance.go


sed -i -e 's/Body Nullable/Body /g' ochk/api/v3/model_http_response.go

sed -i -e 's/BodyReader Nullable/BodyReader /g' ochk/api/v3/model_http_request.go


sed -i -e 's/Body Nullable/Body /g' ochk/api/v3/model_http_request.go

sed -i -e 's/BodyReader Nullable/BodyReader /g' ochk/api/v3/model_http_request.go

sed -i -e 's/Body Nullable/Body /g' ochk/api/v3/model_http_response.go

sed -i -e 's/Body Nullable/Body /g' ochk/api/v3/model_http_request.go

sed -i -e 's/BodyReader Nullable/BodyReader /g' ochk/api/v3/model_http_request.go




jq '
  walk(
    if type == "object" and has("bodyReader")
    then .bodyReader |= del(.nullable)
    else .
    end
  )
' ./ochk/api/swagger.json > ./ochk/api/tmp.json && mv ./ochk/api/tmp.json ./ochk/api/swagger.json

jq '
  walk(
    if type == "object" and has("body")
    then .body |= del(.nullable)
    else .
    end
  )
' ./ochk/api/swagger.json > ./ochk/api/tmp.json && mv ./ochk/api/tmp.json ./ochk/api/swagger.json




# change date-time type field to string in modification and creation date

# for AccountInstance
jq '
  .components.schemas.AccountInstance.properties.creationDate.format = "string"
' ./ochk/api/swagger.json > ./ochk/api/tmp.json && mv ./ochk/api/tmp.json ./ochk/api/swagger.json


#jq '
#  .components.schemas.AccountInstance.properties.modificationDate.format = "string"
#' ./ochk/api/swagger.json > ./ochk/api/tmp.json && mv ./ochk/api/tmp.json ./ochk/api/swagger.json

# for ProjectInstance
#jq '
#  .components.schemas.ProjectInstance.properties.creationDate.format = "string"
#' ./ochk/api/swagger.json > ./ochk/api/tmp.json && mv ./ochk/api/tmp.json ./ochk/api/swagger.json
#
#jq '
#  .components.schemas.ProjectInstance.properties.modificationDate.format = "string"
#' ./ochk/api/swagger.json > ./ochk/api/tmp.json && mv ./ochk/api/tmp.json ./ochk/api/swagger.json
make fmt

go mod tidy

