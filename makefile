.PHONY: format
format:
	go fmt ./...

.PHONY: build
build:
	go build ./cmd

.PHONY: lint
lint:
	golangci-lint run -c ./.golangci.yml > lint.txt

.PHONY: test
test:
	go test ./...

.PHONY: all
all:  format lint build test
