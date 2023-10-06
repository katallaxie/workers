.DEFAULT_GOAL := build

GO ?= go
GO_RUN_TOOLS ?= $(GO) run -modfile ./tools/go.mod
GO_TEST = $(GO) test
GO_RELEASER ?= $(GO_RUN_TOOLS) github.com/goreleaser/goreleaser
GO_MOD ?= $(shell ${GO} list -m)
GO_GOOS ?= js
GO_GOARCH ?= wasm

# Module name
MODULE_NAME ?= github.com/katallaxie/template-go

.PHONY: build
build: ## Build the binary file.
	$(GO_RELEASER) build --snapshot --rm-dist

.PHONY: generate
generate: ## Generate code.
	$(GO) generate ./...

.PHONY: fmt
fmt: ## Run go fmt against code.
	$(GO_RUN_TOOLS) mvdan.cc/gofumpt -w .

.PHONY: vet
vet: ## Run go vet against code.
	GOOS=$(GO_GOOS) GOARCH=$(GO_GOARCH) $(GO) vet ./...

.PHONY: test
test: fmt vet ## Run tests.
	GOOS=$(GO_GOOS) GOARCH=$(GO_GOARCH) $(GO_TEST) ./...

.PHONY: lint
lint: ## Run lint.
	$(GO_RUN_TOOLS) github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout 5m -c .golangci.yml

.PHONY: clean
clean: ## Remove previous build.
	rm -rf .test .dist
	find . -type f -name '*.gen.go' -exec rm {} +
	git checkout go.mod

.PHONY: setup
setup: ## Setup the project.
	$(GO) mod edit -module $(MODULE_NAME)
	find . -type f -name '*.go' -exec sed -i -e 's,${GO_MOD},${MODULE_NAME},g' {} \;

.PHONY: help
help: ## Display this help screen.
	@grep -E '^[a-z.A-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'