VERSION ?= $$(git describe --tags)

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o bin/todo

release: build
	touch bin/empty
	tar czvf todo.tar.gz bin
