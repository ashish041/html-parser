#!/bin/bash

export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin

go test internal/core/domain/domain_test.go
