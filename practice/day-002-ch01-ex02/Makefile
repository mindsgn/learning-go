
.DEFAULT_GOAL := build

.PHONY: fmt vet clean build

fmt: clean
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o hello_world .

clean:
	rm -rvf ./hello_world
