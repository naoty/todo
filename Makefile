install: deps test
	go install

deps:
	go get github.com/codegangsta/cli
	go get github.com/ymotongpoo/goltsv

test:
	go test ./...
