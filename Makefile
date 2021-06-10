.PHONY: all build

build:
	go build -o build/nwaku nwaku.go

all: build
