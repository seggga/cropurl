.PHONY: build

build:
	go build -v ./cmd/cropurl/main.go

.DEFAULT_GOAL := build