#!/bin/sh
set -e

repodir=$(cd "$(dirname "$0")"/..;pwd)
cd "$repodir"

GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=0 go build -o bin/mkrinst ./cmd/mkrinst
