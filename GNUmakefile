default: testacc

PROVIDER_VERSION="1.2.2"

ARCH:=$(shell uname -m)
ifeq ($(ARCH),x86_64)
  ARCH:=amd64
endif

OOS:=darwin


OOS:=linux
ARCH:=amd64
export GOARCH:=$(ARCH)
export GOOS:=$(OOS)

export GOARCH:=$(ARCH)
export GOOS:=$(OOS)

.PHONY: testacc

build:
	go build ./...

EXAMPLES_PROVIDER_DIR="examples/terraform.d/plugins/registry.terraform.io/chmurakrajowa/ochk/${PROVIDER_VERSION}/$(OOS)_$(ARCH)"
ENV_PROVIDER_DIR="env/terraform.d/plugins/registry.terraform.io/chmurakrajowa/ochk/${PROVIDER_VERSION}/$(OOS)_$(ARCH)"

build_local:
	mkdir -p ${EXAMPLES_PROVIDER_DIR}
	mkdir -p ${ENV_PROVIDER_DIR}
	mkdir bin || true
	go build -o bin ./...
	cp bin/terraform-provider-ochk ${EXAMPLES_PROVIDER_DIR}/terraform-provider-ochk_v${PROVIDER_VERSION}
	cp bin/terraform-provider-ochk ${ENV_PROVIDER_DIR}/terraform-provider-ochk_v${PROVIDER_VERSION}

testacc:
	TF_ACC=1 go test ./...

lint:
	tools/lint.sh

swagger-update:
	tools/swagger_update.sh

swagger:
	tools/swagger_gen.sh


