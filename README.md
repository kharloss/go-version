# go-version

## Introduction

In all of my programs, I like to have a pretty little line at the top stating
some key facts about the programs compile.

Example:

```
cloudkey version v1.0.0-rc2 git revision fe3428954dfd97d850ca687badcace28102e351d go version go1.11
```

I was tiring of writing this over, and over, and over again, so I just made a
package.

## Usage

At the top of `main.go`

```
import (
	_ "github.com/jnovack/go-version"
)
```

Then, when you build, add these variables to the command line:

```
VERSION=$(shell git describe --tags)
REVISION=$(shell git rev-parse HEAD)
RELEASE=$(shell git describe --tags | cut -d"-" -f 1,2)
BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%S+00:00")
APPLICATION=$(shell basename `pwd`)
GO_LDFLAGS="-X github.com/jnovack/go-version/build.Application=${application} -X github.com/jnovack/go-version/build.Version=${version} -X github.com/jnovack/go-version/build.Revision=${revision} -X github.com/jnovack/go-version/build.BuildDate=${build_date}"

go build -ldflags GO_LDFLAGS main.go
```

or use the Makefile and do not be a scrub.
