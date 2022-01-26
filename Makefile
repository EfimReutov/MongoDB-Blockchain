NAME=MongoDB-Blockchain

## Running current project in the terminal.
run:
	./scripts/run.sh

## DockerRun
docker-up:
	CGO_ENABLED=0 go build -o main ./cmd
	docker-compose -f docker-compose.yaml up -d
	rm -rf ./main

## Build
build:
	CGO_ENABLED=0 go build -o main ./cmd

## Format all code in the project.
format:
	@echo formating is running...
	go vet ./...
	go fmt ./...