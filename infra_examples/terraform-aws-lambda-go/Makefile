.PHONY: all clean install

all: build/deployment.zip

build/main: main.go
	GOOS=linux go build -o $@ $^

build/deployment.zip: build/main
	zip -j $@ $^

install: build/deployment.zip
	@echo See README.md for installation instructions.

clean:
	rm -f build/*
