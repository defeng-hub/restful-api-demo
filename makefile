PROJECT_NAME=restful-api-demo
MAIN_FILE=main.go
PKG := "github.com/defeng-hub/$(PROJECT_NAME)"

.PHONY: all dep lint vet test test-coverage build clean

run:
	@go run $(MAIN_FILE) start -f etc/pro.toml

