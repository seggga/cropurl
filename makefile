.PHONY: build
build:
	go build -o ./cropurl ./cmd/cropurl/main.go

.PHONY: lint
lint:
	/home/coder/go/bin/golangci-lint run  /home/coder/git/cropurl/... -v --config /home/coder/git/cropurl/golangci.yml

.PHONY: test
test:
	go test ./...
	
.DEFAULT_GOAL := build


.PHONY: build_container
build_container:
	docker build -t docker.io/seggga/cropurl:1.0.0 .

.PHONY: docker_push
docker_puwh:
	 docker push docker.io/seggga/cropurl:1.0.0