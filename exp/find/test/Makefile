all: regenerate build

build:
	go build .
	go test -c

regenerate:
	protoc --gogo_out=. test.proto

clean:
	go clean .
	rm *.pb.go || true
	rm *.json || true
	rm *.txt || true
	rm *.dot || true
	rm *.pdf || true