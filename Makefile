.DEFAULT_GOAL := build
.SILENT:

.PHONY: build
build:
	go build -o ./build/swagmod ./cmd/swagmod/main.go

.PHONY: run-demoserver
run-demoserver:
	go run ./cmd/demoserver/main.go

.PHONY: swag
swag:
	swag init -g ./cmd/demoserver/main.go -o ./docs
