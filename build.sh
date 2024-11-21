#!/usr/bin/env sh
set -e
cd "$(dirname "$0")"


build_backend() {
    cd be/
    go mod tidy
    go generate ./...
    CGO_ENABLED=0 go build -o build/server cmd/main.go
    cd ../
}


build_backend
