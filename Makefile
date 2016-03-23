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

.PHONY: nuke dep regenerate gofmt build test

all: nuke dep regenerate build test checklicense

dep:
	go install github.com/gogo/protobuf/protoc-gen-gogo
	go install -v github.com/goccmack/gocc/...

checklicense:
	go install ./cmd/checklicense
	checklicense .

test:
	go test ./...

build:
	go build ./...

install:
	go install ./...

bench:
	go test -v -test.run=XXX -test.bench=. ./...

regenerate:
	(cd expr && make regenerate)
	(cd tests && make regenerate)
	(cd parser && make regenerate)
	(cd viper && make regenerate)
	(cd relapse && make regenerate)
	(cd expr/funcs && go test -test.run=GenFuncList 2>../../list_of_functions.txt)
	find . -name "*.pb.go" | xargs gofmt -l -s -w
	find . -name "*.gen.go" | xargs gofmt -l -s -w
	find . -name "*.gen_test.go" | xargs gofmt -l -s -w

clean:
	go clean ./...
	(cd expr && make clean)
	(cd viper && make clean)
	(cd relapse && make clean)

nuke: clean
	(cd parser && make nuke)
	(cd expr && make nuke)
	(cd viper && make nuke)
	(cd relapse && make nuke)
	rm list_of_functions.txt || true
	go clean -i ./...

gofmt:
	gofmt -l -s -w .

drone:
	sudo apt-get install protobuf-compiler
	sudo apt-get install graphviz
	mkdir -p $(GOPATH)/src/github.com/gogo
	mkdir -p $(GOPATH)/src/github.com/goccmack
	(cd $(GOPATH)/src/github.com/goccmack && git clone https://github.com/goccmack/gocc )
	(cd $(GOPATH)/src/github.com/goccmack/gocc && go install ./... )
	(cd $(GOPATH)/src/github.com/gogo && git clone https://github.com/gogo/protobuf )
	(cd $(GOPATH)/src/github.com/gogo/protobuf && make install)
	(cd $(GOPATH)/src/github.com/katydid/katydid/ && make)
