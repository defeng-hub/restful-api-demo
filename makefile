PROJECT_NAME=restful-api-demo
MAIN_FILE=main.go
PKG := "/Users/didi/Desktop/byweb/$(PROJECT_NAME)"
MOD_DIR := $(shell go env GOMODCACHE)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all tidy test  build linux clean protoc help

all: build

tidy: ## go mod tidy
	go mod tidy

test: ## Run unittests
	@go test -short ${PKG_LIST}

build: tidy ## 编译可执行文件
	@go build -ldflags "-s -w" -o dist/demo-api $(MAIN_FILE)

linux: tidy ## 编译成amd64版本
	@GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/demo-api $(MAIN_FILE)

run: # Run Develop server
	@go run $(MAIN_FILE) start -f etc/pro.toml

clean: ## 删除构建的内容
	@rm -f dist/*

protoc: ## 生成proto文件
	protoc -I=. --go_out=. --go_opt="" ./*.proto

help: ## 查看帮助
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

