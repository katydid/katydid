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

regenerate:
	(cd test && protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf person.proto)
	(cd test && protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf srctree.proto)
	(cd test && protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf taxonomy.proto)
	(cd test && protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf treeregister.proto)
	(cd test && protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf typewriterprison.proto)
	(cd test && protoc --gogo_out=. -I=.:../../../../../:../../../../../github.com/gogo/protobuf/protobuf puddingmilkshake.proto)
	./build.sh
	(cd ast && make regenerate)

clean:
	rm *.txt || true
	rm -rf ./test/example/*

example:
	(cd test && go test -c && ./test.test)

nuke: clean
	rm -rf ./errors
	rm -rf ./lexer
	rm -rf ./parser/action.go
	rm -rf ./parser/actiontable.go
	rm -rf ./parser/gototable.go
	rm -rf ./parser/parser.go
	rm -rf ./parser/productionstable.go
	rm -rf ./token
	rm -rf ./util