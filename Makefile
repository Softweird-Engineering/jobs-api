.PHONY: build
build:
		go build -v ./src/apiserver

.DEFAULT_GOAL := build