#! /usr/bin/make

DIR = $(shell go list)

.PHONY:

all:
	@echo "usage: make gen|test"

gen:
	@go run tools/gen.go > ./jarr.go

test:
	@go test -v
