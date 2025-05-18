.PHONY: all clean format vet build

all: clean fmt vet build

clean:
	go clean
	rm -rf ./bin

fmt:
	go fmt ./...

vet:
	go vet ./...

build:
	mkdir -p bin
	go build -o bin/ cmd/kafka-tryout/main.go