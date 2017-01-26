VERSION ?= $$(git describe --tags)

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o bin/todo
