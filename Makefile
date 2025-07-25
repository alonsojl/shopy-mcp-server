# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: up
up: ## Start container services.
	@docker compose up -d
	
.PHONY: down
down: ## Stop container services.
	@docker compose down

.PHONY: linter
linter: ## Lint source code.
	@golangci-lint run -c .golangci.yml > linter.txt

.PHONY: clean
clean: ## Clean build files and cache.
	@go clean
	@rm -rf ./bin/mcp-server

.PHONY: build
build: clean ## Build application.
	@go mod tidy
	@go build -o ./bin/mcp-server ./cmd/main.go

.PHONY: run
run: build ## Run application.
	@./bin/mcp-server

.PHONY: live
live:  ## Live reload for Go applications.
	@air -c .air.toml

.PHONY: tools
tools: ## Install tools.
	@go install github.com/air-verse/air@v1.61.7
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.3.0


