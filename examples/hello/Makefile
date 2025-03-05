WRANGLER 		?= npx -y wrangler
GO 				?= go
GO_RUN_TOOLS	?= $(GO) run -modfile ../../tools/go.mod
GO_RUN 			?= $(GO) run -modfile ../../go.mod
TINYGO 			?= tinygo

.PHONY: dev
dev:
	$(WRANGLER) dev --ip 127.0.0.1

.PHONY: build
build:
	$(GO_RUN) ../../cmd/gen/main.go
	GOOS=js GOARCH=wasm go build -o ./build/app.wasm
	# $(TINYGO) build -o ./build/app.wasm -target wasm -no-debug ./...

.PHONY: deploy
deploy:
	$(WRANGLER) deploy

.PHONY: clean
clean:
	@rm -rf build .wrangler node_modules