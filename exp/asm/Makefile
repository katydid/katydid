.PHONY: nuke regenerate build test

all: nuke regenerate build test

test:
	go test -v ./...

build:
	go build ./...

regenerate:
	gocc asm.bnf
	(cd ast && protoc --gogo_out=. -I=.:../../../../../../ asm.proto)
	(cd test && protoc --gogo_out=. -I=.:../../../../../../ test.proto)

clean:
	rm *.txt || true

nuke: clean
	rm -rf errors
	rm -rf lexer
	rm -rf parser
	rm -rf token
	rm -rf util