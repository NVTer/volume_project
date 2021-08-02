.PHONY: format
format:
	go fmt ./...

.PHONY: build
build:
	go build ./cmd/main.go

.PHONY: lint
lint:
	golangci-lint run -c ./.golangci.yml > lint.txt

.PHONY: test
test:
	go test -v ./cmd/main_test.go

.PHONY: all
all:  format lint build test
