SHELL := /bin/bash

example.go: example.go.rl
	ragel -Z -G2 -e -o $@ $<
	@gofmt -w -s $@
	@sed -i '/^\/\/line/d' $@

.PHONY: build
build: example.go

.PHONY: tests
tests: *_test.go
	go test -race -timeout 10s -v ./...
