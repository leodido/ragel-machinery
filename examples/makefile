SHELL := /bin/bash

.PHONY: all
all: newlines.go multiline.go

newlines.go: newlines.go.rl
	ragel -Z -G2 -e -o $@ $<
	@gofmt -w -s $@
	@sed -i '/^\/\/line/d' $@

multiline.go: multiline.go.rl
	ragel -Z -G2 -e -o $@ $<
	@gofmt -w -s $@
	@sed -i '/^\/\/line/d' $@

.PHONY: tests
tests: *_test.go
	go test -race -timeout 10s -v ./...
