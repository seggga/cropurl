.PHONY: build
build:
	go build -v ./cmd/cropurl/main.go

.PHONY: lint
lint:
	/home/coder/go/bin/golangci-lint run  /home/coder/git/cropurl/... -v --config /home/coder/git/cropurl/golangci.yml

.PHONY: test
test:
	go test ./...
	
.DEFAULT_GOAL := build