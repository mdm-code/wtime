GO=go
GOFLAGS=--race
DEV_BIN=bin

all: test build

test:
	$(GO) test -v ./...

build:
	$(GO) build $(GOFLAGS) -o $(DEV_BIN)/wtime main.go

install: test
	go install ./...

.PHONY: all test build install
