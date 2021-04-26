GO := go
COMPOSE := docker compose

.PHONY: build-go
build-go:
	$(GO) build -o ./bin/app

.PHONY: build-docker
build-docker:
	$(COMPOSE) build

.PHONY: build
build: build-go build-docker

.PHONY: start
start: build
	$(COMPOSE) up -d