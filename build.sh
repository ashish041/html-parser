#!/bin/bash

export CGO_ENABLED=0
export GO111MODULE=on
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin

go mod tidy
go build cmd/httpserver/main.go
