#! /usr/bin/make

sources := $(shell find ./templates -name '*.go.tmpl')
functions := $(foreach name, $(sources), $(subst .go.tmpl,,$(subst ./templates/,,$(name))))
generated := $(foreach name, $(functions), $(name).gen.go)

.PHONY:

all: *.go $(generated)
	@go test -v

%.gen.go: templates/%.go.tmpl types.yml
	@echo generating... $@
	@go-template $< < types.yml | gofmt | goimports > $@

test:
	@go test -v

deps:
	@go get -u cmd/gofmt
	@go get -u golang.org/x/tools/cmd/goimports
	@go get -u github.com/badugisoft/go-template
	@go get -u github.com/badugisoft/go-regexp

doc:
	@cat *_test.go | go-regexp "Examples=(?s:func Example.+?\n})"\
		| go-template templates/README.md.tmpl > README.md

perf:
	@go run perf/main.go > perf/README.md

clean:
	@rm *.gen.go
