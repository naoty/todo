VERSION ?= $$(git describe --tags)

.PHONY: build
build:
	go build -ldflags "-X main.Version=$(VERSION)" -o bin/todo
