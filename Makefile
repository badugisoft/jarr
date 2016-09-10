#! /usr/bin/make

.PHONY:

test:
	@go test -v

gen:
	@go run tools/gen.go > ./jarr.go
