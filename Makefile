.PHONY: build

default: build

BINARY=go_amis_admin
GIT_VERSION := $(shell git rev-parse HEAD)
GO_VERSION := $(shell go version)
BUILD_TIME := $(shell date +%FT%T%z)

LDFLAGS=-ldflags '-s -X "github.com/fs714/go-amis-admin/utils/version.GitVersion=${GIT_VERSION}" -X "github.com/fs714/go-amis-admin/utils/version.GoVersion=${GO_VERSION}" -X "github.com/fs714/go-amis-admin/utils/version.BuildTime=${BUILD_TIME}"'

bindata:
	go-bindata-assetfs  static/...
build:
	go-bindata-assetfs  static/...
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/${BINARY} ${LDFLAGS}
clean:
	rm bindata.go
	rm -rf bin/
