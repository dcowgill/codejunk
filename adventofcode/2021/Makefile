.PHONY: all clean default main test

all: test

default: test

test: main
	@./main

main:
	go build -o main main.go

clean:
	@rm main
