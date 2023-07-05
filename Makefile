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
	@/bin/sh -c 'echo "${GREEN}빌드를 시작합니다.${NC}"'
	@mkdir -p bin
	@go mod download && go mod verify
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/
.PHONY: build

# Update Swagger
docs:
	@swag init --parseDependency --parseInternal
.PHONY: docs