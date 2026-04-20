build:
	go build -o bin/go-editor .

run:
	go run ./...

.PHONY: build run
