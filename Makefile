BINARY=keyfactor-go-client
VERSION=1.0.0

default: build

build:
	go build -o ${BINARY}