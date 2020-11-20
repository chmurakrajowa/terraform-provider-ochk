default: testacc

.PHONY: testacc

build:
	go build ./...

EXAMPLES_PROVIDER_DIR=examples/terraform.d/plugins/registry.terraform.io/chmurakrajowa/ochk/0.1.0/darwin_amd64
ENV_PROVIDER_DIR=env/terraform.d/plugins/registry.terraform.io/chmurakrajowa/ochk/0.1.0/darwin_amd64

build_local:
	mkdir -p ${EXAMPLES_PROVIDER_DIR}
	mkdir -p ${ENV_PROVIDER_DIR}
	mkdir bin || true
	go build -o bin ./...
	cp bin/terraform-provider-ochk ${EXAMPLES_PROVIDER_DIR}/terraform-provider-ochk_v0.1.0
	cp bin/terraform-provider-ochk ${ENV_PROVIDER_DIR}/terraform-provider-ochk_v0.1.0

testacc:
	TF_ACC=1 go test ./...

lint:
	tools/lint.sh

swagger-update:
	tools/swagger_update.sh

swagger:
	tools/swagger_gen.sh


