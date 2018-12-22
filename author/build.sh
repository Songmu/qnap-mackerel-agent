#!/bin/sh
set -e

repodir=$(cd "$(dirname "$0")"/..;pwd)
cd "$repodir"

cd "$(go env GOPATH)/src/github.com/mackerelio/mackerel-agent"
GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=0 make build

mv build/mackerel-agent "$repodir/bin/mackerel-agent"

perl "$repodir/author/bump_version.pl" CHANGELOG.md "$repodir/qpkg.conf.txt"
