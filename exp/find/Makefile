all: clean regenerate build test

build:
	go build ./...
	go test -v ./...

regenerate:
	gocc find.bnf
	(cd ast && protoc --gogo_out=. -I=.:../../../../../../ find.proto)
	(cd test && make regenerate)

clean:
	rm -rf errors
	rm -rf lexer
	rm -rf parser
	rm -rf token
	rm -rf util
	rm *.txt || true
	(cd test && make clean)