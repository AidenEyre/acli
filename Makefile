GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=acli
VERSION?=1.4.0

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test build vendor

all: help

## Build:
build: ## Build your project and put the output binary in out/bin/
	mkdir -p out/bin
	GO111MODULE=on $(GOCMD) build -o out/bin/$(BINARY_NAME) .

clean: ## Remove build related file
	rm -fr ./bin
	rm -fr ./out

## Install:
install: ## Install your project onto your machine
	$(GOCMD) install .

## Test:
test: ## Run the tests of the project
	$(GOTEST) -v -race ./...

## Release:
release: ## Build binaries for different OS types
	GOOS=windows GOARCH=amd64 go build -o bin/acli-amd64.exe .
	GOOS=darwin GOARCH=arm64 go build -o bin/acli-macos-arm64 .
	GOOS=linux GOARCH=386 go build -o bin/acli-linux-386 .

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)