# メタ情報
NAME := go_development
VERSION := $(gobump show -r)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := "-X main.revision=$(REVISION)"

export GO111MODULE=on

## Install dependencies
.PHONY: deps
deps:
	go get -v -d

# 開発に必要な依存関係をインストールする
## Setup
.PHONY: devel-deps
devel-deps: deps
	GO111MODULE=off go get\
		github.com/motemen/gore/cmd/gore\
		github.com/mdempsky/gocode\
		github.com/k0kubun/pp\
		golang.org/x/tools/cmd/goimports\
		golang.org/x/lint/golint\
		github.com/motemen/gobump/cmd/gobump\
		github.com/Songmu/make2help/cmd/make2help

# テストを実行する
## Run tests
.PHONY: test
test: deps
	go test ./...

## Lint
.PHONY: lint
lint:
	go vet ./...
	golint -set_exit_status ./...

## build binaries ex. make bin/myproj
bin/%: cmd/%/main.go deps
	go build -ldflags "$(LDFLAGS)" -o $@ $<

## build binary
.PHONY: build
build: bin/go_development

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)