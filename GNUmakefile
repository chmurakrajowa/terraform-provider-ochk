default: testacc

.PHONY: testacc

build:
	go build ./...

testacc:
	TF_ACC=1 go test ./...

lint:
	tools/lint.sh

swagger-update:
	tools/swagger_update.sh

swagger:
	tools/swagger_gen.sh

