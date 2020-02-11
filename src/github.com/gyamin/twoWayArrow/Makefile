ROOT := $(shell cd ./../../../../; pwd)
BIN_DIR := $(ROOT)/bin
BINARY_NAME := textToTable
GOPATH := $(ROOT)

VERSION := $(shell git describe --tags --abbrev=0)
LDFLAGS := -X 'main.version=$(VERSION)'

build:
	go build -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/$(BINARY_NAME)

prepare:
	GOPATH=$(GOPATH) GOBIN=$(BIN_DIR) go get -u github.com/golang/dep/cmd/dep
	GOPATH=$(GOPATH) $(BIN_DIR)/dep ensure -v

test:
	GOPATH=$(GOPATH) go test -v cmd/...

.PHONY: test