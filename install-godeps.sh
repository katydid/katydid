#!/usr/bin/env bash
set -xe
mkdir -p $GOPATH/src/github.com/gogo
git clone https://github.com/gogo/protobuf $GOPATH/src/github.com/gogo/protobuf
mkdir -p $GOPATH/src/githbub.com/goccmack
git clone https://github.com/goccmack/gocc $GOPATH/src/github.com/goccmack/gocc

