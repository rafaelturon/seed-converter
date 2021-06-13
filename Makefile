.PHONY: all
all: help
FORCE: ;

SHELL  := env LIBRARY_ENV=$(LIBRARY_ENV) $(SHELL)
LIBRARY_ENV ?= dev
BINARY_NAME=print

BIN_DIR = $(PWD)/bin

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: build

## Build:
clean: ## Remove build related file
	rm -rf ./bin
	rm -rf ./out
	rm -rf ./vendor
	rm -f ./coverage.out

dependencies:
	go mod download

build: dependencies build-cmd ## Build your project and put the output binary in out/bin/

build-cmd:
	mkdir -p out/bin
	GO111MODULE=on go build -mod vendor -o out/bin/$(BINARY_NAME) .
	##go build -tags $(LIBRARY_ENV) -o ./bin/print cmd/dcrseed/main.go

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	go mod vendor

linux-binaries:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(LIBRARY_ENV) netgo" -installsuffix netgo -o $(BIN_DIR)/search cmd/main.go

run:
	go run cmd/dcrseed/main.go

ci: dependencies test	

## Test:
test: ## run unit tests
	go test -tags testing ./...

coverage: ## run coverage tests
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

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