PHONY: deps
deps:
	go mod tidy

PHONY: build
build: deps
	go build -o bin/fruitpass