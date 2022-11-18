# Version Vars
VERSION_TAG := $(shell git describe --tags --always)
VERSION_VERSION := $(shell git log --date=iso --pretty=format:"%cd" -1) $(VERSION_TAG)
VERSION_COMPILE := $(shell date +"%F %T %z") by $(shell go version)
VERSION_BRANCH  := $(shell git rev-parse --abbrev-ref HEAD)
VERSION_GIT_DIRTY := $(shell git diff --no-ext-diff 2>/dev/null | wc -l | awk '{print $1}')
VERSION_DEV_PATH:= $(shell pwd)

# Go Checkup
GOPATH ?= $(shell go env GOPATH)
GO111MODULE:=auto
export GO111MODULE
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
PATH := ${GOPATH}/bin:$(PATH)
GCFLAGS=-gcflags="all=-trimpath=${GOPATH}"
LDFLAGS=-ldflags="-s -w -X 'main.Version=${VERSION_VERSION}' -X 'main.Compile=${VERSION_COMPILE}' -X 'main.Branch=${VERSION_BRANCH}'"

GO = go

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m➡\033[0m")


#Commands
.PHONY: all
all: |help

.PHONY: init
init: ; $(info $(M) Install project tools dependencies...) @ ## Install tools dependencies
	pip3 install -U commitizen pre-commit

	# Install project pre-commit hooks from file
	pre-commit install --install-hooks
	pre-commit install --hook-type commit-msg

	# Go libraries for pre-commit test
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) install github.com/go-critic/go-critic/cmd/gocritic@latest
	$(GO) install github.com/sqs/goreturns@latest
	$(GO) install golang.org/x/tools/cmd/goimports@latest
	$(GO) install golang.org/x/vuln/cmd/govulncheck@latest

.PHONY: deps
deps: ; $(info $(M) Install project go dependecies...) @ ## Install project dependencies
	$Q $(GO) mod tidy

.PHONY: test
test: ; $(info $(M) Running unit tests...) @ ## Run unit test
	$Q $(GO) test -v -covermode=count -coverprofile coverage.out ./...

.PHONY: build
build: ; $(info $(M) Install project tools dependencies...) @ ## Build project binary
	$Q mkdir -p bin
	$Q $(GO) generate ./...
	$Q ret=0 && for d in $$($(GO) list -f '{{if (eq .Name "main")}}{{.ImportPath}}{{end}}' ./...); do \
		b=$$(basename $${d}) ; \
		$(GO) build ${LDFLAGS} ${GCFLAGS} -o bin/$${b} $$d || ret=$$? ; \
		echo "$(M) Building: bin/$${b}" ; \
		echo "$(M) Done!" ; \
  	done ; exit $$ret


.PHONY: run
run: ; $(info $(M) Running a development build...) @ ## Run a development build
	$Q ret=0 && for d in $$($(GO) list -f '{{if (eq .Name "main")}}{{.ImportPath}}{{end}}' ./...); do \
		b=$$(basename $${d}) ; \
		$(GO) run -race $$d || ret=$$? ; \
	done ; exit $$ret

help:
	$Q echo "KAPE TECH makefile\n----------------"
	$Q grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'