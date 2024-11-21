#!/usr/bin/env sh
set -e
cd "$(dirname "$0")"


build_backend() {
    cd be/
    CGO_ENABLED=0 go build -o build/server cmd/server/main.go
    cd ../
}


build_backend
