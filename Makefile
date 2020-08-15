SHELL := /bin/bash

.PHONY: install run test

.PHONY: all test

all: help

help:
	@echo
	@echo "usage: make <command>"
	@echo
	@echo "commands:"
	@echo "    install   - get the modules"
	@echo "    test      - run all tests"
	@echo "    run       - execute"
	@echo

install:
	go mod tidy

test:
	go test -tags=integration,e2e ./... -v

run:
	go run ./cmd/main.go