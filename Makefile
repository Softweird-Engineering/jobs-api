.PHONY: build
build:
	go build -v ./src

.DEFAULT_GOAL := build