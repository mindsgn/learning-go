.DEFAULT_GOAL := build

fmt: 
	go fmt ./...

vet: fmt
	go vet ./...

test: vet
	go test -v

build: test
	go build -o number

clean: 
	go clean
	rm -f number