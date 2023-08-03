# .env 존재시 export
ENV_FILE_EXISTS := $(wildcard .env)
ifneq ($(ENV_FILE_EXISTS),)
	include .env
	export
endif

# install dependencies
ref:
	@go mod tidy && go mod vendor
.PHONY: ref

# build
build:
	@/bin/sh -c 'echo "${GREEN}docker 빌드 및 push를 시작합니다. Version: ${VERSION} ${NC}"'
	@docker build -t ghcr.io/nexters/pinterest:${VERSION} .
	@/bin/sh -c 'echo "${GREEN}[Push Image] Pushing version: ${VERSION} ... ${NC}"'
	@docker push ghcr.io/nexters/pinterest:${VERSION}
.PHONY: build

# Update Swagger
docs:
	@swag init
.PHONY: docs