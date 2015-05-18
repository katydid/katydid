//  Copyright 2015 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package protoparser

import (
	"github.com/gogo/protobuf/parser"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
)

// import (
// 	"github.com/dsymonds/gotoc/internal/gendesc"
// 	"github.com/dsymonds/gotoc/internal/parser"
// 	"github.com/gogo/protobuf/proto"
// )

func ParseFile(filename string, paths ...string) (*descriptor.FileDescriptorSet, error) {
	return parser.ParseFile(filename, paths...)
}

// func ParseFile(filename string, paths ...string) (*descriptor.FileDescriptorSet, error) {
// 	fs, err := parser.ParseFiles([]string{filename}, paths)
// 	if err != nil {
// 		return nil, err
// 	}
// 	golangfds, err := gendesc.Generate(fs)
// 	if err != nil {
// 		return nil, err
// 	}
// 	data, err := proto.Marshal(golangfds)
// 	if err != nil {
// 		return nil, err
// 	}
// 	fds := &descriptor.FileDescriptorSet{}
// 	if err := proto.Unmarshal(data, fds); err != nil {
// 		return nil, err
// 	}
// 	return fds, nil
// }
