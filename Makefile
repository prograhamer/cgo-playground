.PHONY: all
all: clean build

clean:
	rm -f test

build:
	CGO_ENABLED=true go build ./test.go
