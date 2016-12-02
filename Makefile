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

all: nuke dep regenerate build test vet checklicense

dep:
	go install github.com/gogo/protobuf/protoc-gen-gogo
	go install -v github.com/goccmack/gocc

checklicense:
	go get github.com/katydid/checklicense
	checklicense . \
	person.proto \
	srctree.proto \
	puddingmilkshake.proto \
	taxonomy.proto \
	treeregister.proto \
	typewriterprison.proto \
	proto/tokens/test.proto \
	parser/debug/debug.proto \
	bnf \
	doc.go \
	.svg \
	.txt \
	COPIED_FROM_GO \
	.travis.yml \
	install-protobuf.sh \
	install-godeps.sh

test:
	go test ./...

build:
	go build ./...

install:
	go install ./...

bench:
	go test -v -test.run=XXX -test.bench=. ./...

vet:
	go vet ./encode/...
	go vet ./gen/...
	go vet ./relapse/...

regenerate:
	(cd parser && make regenerate)
	(cd relapse && make regenerate)
	(cd relapse/funcs && go test -test.run=GenFuncList 2>../../list_of_functions.txt)
	(cd encode && make regenerate)
	find . -name "*.pb.go" | xargs gofmt -l -s -w
	find . -name "*.gen.go" | xargs gofmt -l -s -w
	find . -name "*.gen_test.go" | xargs gofmt -l -s -w

clean:
	go clean ./...
	(cd relapse && make clean)

nuke: clean
	(cd parser && make nuke)
	(cd relapse && make nuke)
	rm list_of_functions.txt || true
	go clean -i ./...

gofmt:
	gofmt -l -s -w .

errcheck:
	go get github.com/kisielk/errcheck
	errcheck ./...
