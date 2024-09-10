GOFMT_FILES?=$$(find . -name '*.go' | grep -v 'vendor')
PROVIDER_VERSION?=$$(sed -n -e '1{s/\#//g;p;}' CHANGELOG.md | awk '{print $$1}')

ARCH:=$(shell uname -m)
ifeq ($(ARCH),x86_64)
  ARCH:=amd64
endif
OS:=$(shell uname -s | tr '[:upper:]' '[:lower:]')

EXAMPLES_PROVIDER_DIR="examples/terraform.d/plugins/registry.terraform.io/chmurakrajowa/ochk/$(PROVIDER_VERSION)/$(OS)_$(ARCH)"
ENV_PROVIDER_DIR="env/terraform.d/plugins/registry.terraform.io/chmurakrajowa/ochk/$(PROVIDER_VERSION)/$(OS)_$(ARCH)"

default: build

clean:
	go clean ./...
build: fmtcheck
	go build ./...

build_local: fmtcheck
	mkdir -p $(EXAMPLES_PROVIDER_DIR)
	mkdir -p $(ENV_PROVIDER_DIR)
	mkdir -p bin
	go build -o bin ./...
	cp bin/terraform-provider-ochk $(EXAMPLES_PROVIDER_DIR)/terraform-provider-ochk_v$(PROVIDER_VERSION)
	cp bin/terraform-provider-ochk $(ENV_PROVIDER_DIR)/terraform-provider-ochk_v$(PROVIDER_VERSION)

test: fmtcheck
	go test ./...

testacc: fmtcheck
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

fmt:
	gofmt -w -s $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/tools/gofmtcheck.sh'"

vet:
	go vet

lint:
	@sh -c "'$(CURDIR)/tools/lint.sh'"

swagger-update:
	@sh -c "'$(CURDIR)/tools/swagger_update.sh'"

swagger-generate:
	@sh -c "'$(CURDIR)/tools/swagger_gen.sh'"

.PHONY: build test testacc fmt fmtcheck vet lint swagger-update swagger-generate
