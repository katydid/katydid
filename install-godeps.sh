#!/usr/bin/env bash
set -xe
mkdir -p $GOPATH/src/github.com/gogo
git clone https://github.com/gogo/protobuf $GOPATH/src/github.com/gogo/protobuf
mkdir -p $GOPATH/src/github.com/goccmack
git clone https://github.com/goccmack/gocc $GOPATH/src/github.com/goccmack/gocc
mkdir -p $GOPATH/src/github.com/awalterschulze
git clone https://github.com/awalterschulze/goderive $GOPATH/src/github.com/awalterschulze/goderive
mkdir -p $GOPATH/src/github.com/katydid
git clone https://github.com/katydid/testsuite $GOPATH/src/github.com/katydid/testsuite
