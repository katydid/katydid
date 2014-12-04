//  Copyright 2013 Walter Schulze
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

package tokens

import (
	"fmt"
	"testing"

	protoparser "github.com/gogo/protobuf/parser"
)

func TestTokensPerson(t *testing.T) {
	fileDescriptorSet, err := protoparser.ParseFile("test.proto", ".")
	if err != nil {
		panic(err)
	}
	m, err := new("test", "Person", fileDescriptorSet, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", m.Dot())
}

func TestTokensSrcTree(t *testing.T) {
	fileDescriptorSet, err := protoparser.ParseFile("test.proto", ".")
	if err != nil {
		panic(err)
	}
	m, err := new("test", "SrcTree", fileDescriptorSet, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", m.Dot())
}
