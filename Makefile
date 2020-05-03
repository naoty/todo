VERSION ?= $$(git describe --tags)

.PHONY: build
build: test
	go build -ldflags "-X main.Version=$(VERSION)" -o bin/todo

.PHONY: test
test:
	go test ./...
