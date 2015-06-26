# Copyright 2013 Walter Schulze
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: nuke regenerate gofmt build test

all: nuke dep regenerate gofmt build test checklicense

dep:
	go install github.com/gogo/protobuf/protoc-gen-gogo
	go install -v code.google.com/p/gocc/...

checklicense:
	go install ./cmd/checklicense
	checklicense .

test:
	go test -v ./expr/...
	go test -v ./funcs/...
	go test -v ./gen/...
	go test -v ./graphviz/...
	go test -v ./lang/...
	go test -v ./protoparser/...
	go test -v ./serialize/...
	go test -v ./tests/...
	go test -v ./types/...

build:
	go build ./...

install:
	go install ./...

bench:
	go test -v -test.run=XXX -test.bench=. ./...

regenerate:
	(cd types && protoc --gogo_out=. -I=.:../../../../:../../../../github.com/gogo/protobuf/protobuf types.proto)
	(cd expr && make regenerate)
	(cd asm && make regenerate)
	(cd lang && make regenerate)
	(cd tests && make regenerate)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf person.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf srctree.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf taxonomy.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf treeregister.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf typewriterprison.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf puddingmilkshake.proto)
	go install github.com/katydid/katydid/funcs/funcs-gen
	funcs-gen ./funcs/
	go install github.com/katydid/katydid/serialize/serialize-gen
	serialize-gen ./serialize
	(cd funcs && go test -test.run=GenFuncList 2>../list_of_functions.txt)
	make gofmt

clean:
	(cd asm && rm *.txt || true)
	rm -rf asm/test/example/*
	go clean ./...

example:
	(cd asm/test && go test -c && ./test.test)

nuke: clean
	rm -rf funcs/compare.gen.go
	rm -rf funcs/newfunc.gen.go
	rm -rf funcs/const.gen.go
	rm -rf funcs/elem.gen.go
	rm -rf funcs/length.gen.go
	rm -rf funcs/list.gen.go
	rm -rf funcs/print.gen.go
	rm -rf asm/compose/compose.gen.go
	rm -rf asm/conv/conv.gen.go
	rm list_of_functions.txt || true
	rm -rf asm/errors
	rm -rf asm/lexer
	rm -rf asm/parser/action.go
	rm -rf asm/parser/actiontable.go
	rm -rf asm/parser/gototable.go
	rm -rf asm/parser/parser.go
	rm -rf asm/parser/productionstable.go
	rm -rf asm/token
	rm -rf asm/util
	go clean -i ./...

gofmt:
	gofmt -l -s -w .

drone:
	sudo apt-get install protobuf-compiler
	sudo apt-get install graphviz
	mkdir -p $(GOPATH)/src/github.com/gogo
	mkdir -p $(GOPATH)/src/code.google.com/p
	(cd $(GOPATH)/src/code.google.com/p && git clone https://code.google.com/p/gocc )
	(cd $(GOPATH)/src/code.google.com/p/gocc && git checkout gocc2 && go install ./... )
	(cd $(GOPATH)/src/github.com/gogo && git clone https://github.com/gogo/protobuf )
	(cd $(GOPATH)/src/github.com/gogo/protobuf && make )
	go get -v code.google.com/p/go.text/unicode/norm
	(cd $(GOPATH)/src/github.com/katydid/katydid/ && make)
