BINARY_NAME=cliapp-wc

#GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

build:
	go build -o $(BINARY_NAME)-$(GOARCH) .

run:
	go run .

test:
	go test -v .


