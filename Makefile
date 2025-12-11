.PHONY: test-% test-all build run clean

all: test-all

test-%:
	cd day-$* && go test -v

test-all:
	go test ./...

build:
	go build -o aoc2025

run: build
	./aoc2025

clean:
	rm -f aoc2025
