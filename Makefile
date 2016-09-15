#! /usr/bin/make

sources := $(shell find ./templates -name '*.tmpl')
functions := $(foreach name, $(sources), $(subst .tmpl,,$(subst ./templates/,,$(name))))
generated := $(foreach name, $(functions), $(name).gen.go)

.PHONY:

all: *.go $(generated)
	@go test -v

%.gen.go: templates/%.tmpl
	@echo generating... $@
	@gotpl $< < types.yml | gofmt | goimports > $@

test:
	@go test -v

tools:
	@go get cmd/gofmt
	@go get golang.org/x/tools/cmd/goimports
	@go get github.com/tsg/gotpl

clean:
	@rm *.gen.go
