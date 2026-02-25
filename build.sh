#!/bin/bash

set -e

version=${1:-"v1.0.0"}

# macOS x86_64
GOOS=darwin
GOARCH=amd64
go build -o internal-proxy-darwin-amd64-${version} .

# macOS arm64(Apple Silicon)
GOOS=darwin
GOARCH=arm64
go build -o internal-proxy-darwin-arm64-${version} .

# Linux
# GOOS=linux
# GOARCH=amd64
# go build -o internal-proxy-linux-amd64-${version} .

# Windows
GOOS=windows
GOARCH=amd64
go build -o internal-proxy-windows-amd64-${version} .