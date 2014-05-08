.PHONY: nuke regenerate gofmt build test

all: nuke regenerate gofmt build test

test:
	go test -v ./...

build:
	go build ./...

bench:
	go test -v -test.run=XXX -test.bench=. ./...

regenerate:
	(cd types && protoc --gogo_out=. -I=.:../../../../ types.proto)
	(cd asm && gocc asm.bnf)
	(cd asm/ast && protoc --gogo_out=. -I=.:../../../../../ asm.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../ person.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../ srctree.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../ taxonomy.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../ treeregister.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../ typewriterprison.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../ puddingmilkshake.proto)
	(cd funcs && protoc --gogo_out=. -I=.:../../../../ decode.proto)
	go install github.com/awalterschulze/katydid/funcs/funcs-gen
	funcs-gen ./funcs/
	go install github.com/awalterschulze/katydid/asm/compose/compose-gen
	compose-gen ./asm/compose/
	go install github.com/awalterschulze/katydid/asm/conv/conv-gen
	conv-gen ./asm/conv/
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