#! /usr/bin/make

sources := $(shell find ./templates -name '*.tmpl')
functions := $(foreach name, $(sources), $(subst .tmpl,,$(subst ./templates/,,$(name))))
generated := $(foreach name, $(functions), $(name).gen.go)

.PHONY:

all: *.go $(generated)
	@go test -v

%.gen.go: templates/%.tmpl types.yml
	@echo generating... $@
	@go-template $< < types.yml | gofmt | goimports > $@

test:
	@go test -v

tools:
	@go get -u cmd/gofmt
	@go get -u golang.org/x/tools/cmd/goimports
	@go get -u github.com/badugisoft/go-template

perf:
	@go run perf/main.go > perf/README.md

clean:
	@rm *.gen.go
