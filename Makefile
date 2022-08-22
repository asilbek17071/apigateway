.PHONY: build

CURRENT_DIR=$(shell pwd)
APP=crm
APP_CMD_DIR=./cmd

build:
	CGO_ENABLED=0 GOOS=linux go build -mod='' -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto:
	./scripts/gen-proto.sh	${CURRENT_DIR}

lint: ## Run golangci-lint with printing to stdout
	golangci-lint -c .golangci.yaml run --build-tags "musl" ./...