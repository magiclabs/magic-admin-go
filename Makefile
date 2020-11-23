PROJECT_PATH ?= $(shell pwd -P)

deps:
	$(info * Installing dependencies)
	@go mod download

.PHONY: build
build: deps
	$(info * Building magic-cli)
	@go build ./cmd/magic-cli

.PHONY: install
install: deps
	$(info * Installing magic-cli)
	@go install ./cmd/magic-cli

.PHONY: test
test: deps
	$(info * Running tests)
	@go test ./...
