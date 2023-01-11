PHONY: deps
deps:
	go mod tidy

PHONY: build
build: deps
	go build -o bin/fruitpass

PHONY: install
install: deps build
	cp bin/fruitpass /usr/local/bin/fruitpass

PHONY: uninstall
uninstall:
	rm /usr/local/bin/fruitpass