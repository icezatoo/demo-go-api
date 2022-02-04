SELL := /bin/base

.PHONY: all build test deps deps-clearcahe

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin

all: test build

${BINARY_DIR}:
	mkdir -p $(BINARY_DIR)

build: ${BINARY_DIR}
	$(GOCMD) build -o $(BINARY_DIR) -v ./cmd/api

run:
	$(GOCMD) run ./cmd/api

test:
	$(GOCMD) test ./...
