#!/bin/sh
set -e

repodir=$(cd $(dirname $0);pwd)
cd $repodir
mkdir -p bin

cd $(go env GOPATH)/src/github.com/mackerelio/mackerel-agent
GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=0 make build

mv build/mackerel-agent "$repodir/bin/mackerel-agent"
