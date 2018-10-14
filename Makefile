version := $(shell git describe --tags)
revision := $(shell git rev-parse HEAD)
release := $(shell git describe --tags | cut -d"-" -f 1,2)
build_date := $(shell date -u +"%Y-%m-%dT%H:%M:%S+00:00")
application := $(shell basename `pwd`)

GO_LDFLAGS := "-X github.com/jnovack/go-version/build.Application=${application} -X github.com/jnovack/go-version/build.Version=${version} -X github.com/jnovack/go-version/build.Revision=${revision}"

all: build

.PHONY: build
build:
	go build -ldflags $(GO_LDFLAGS) main.go
