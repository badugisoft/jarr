#! /usr/bin/make

.PHONY:

all: test

test: *.go
	@go test -v

each.go: templates/each.tmpl types.yml
	@echo generating... each.go
	@gotpl templates/each.tmpl < types.yml | gofmt > each.go

map.go: templates/map.tmpl types.yml
	@echo generating... map.go
	@gotpl templates/map.tmpl < types.yml | gofmt > map.go

tools:
	@go get cmd/gofmt
