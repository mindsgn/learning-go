.DEFAULT_GOAL := build

.PHONY: fmt vet test build clean

fmt: 
	go fmt ./...

vet: fmt
	go vet ./...

test: vet
	go test -v

build: test
	go build -o hello_world

clean:
	go clean
	rm -f hello_world

