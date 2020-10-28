#!/bin/sh

GIT_VER=`git describe --tags`
go build -ldflags "-X main.version=${GIT_VER}"