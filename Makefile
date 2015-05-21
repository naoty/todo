deps:
	go get github.com/codegangsta/cli
	go get github.com/ymotongpoo/goltsv

install: deps
	go install
