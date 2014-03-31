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

package funcs

import (
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	"testing"
)

type which struct {
	exp string
}

func (this which) test(t *testing.T, name string, params ...descriptor.FieldDescriptorProto_Type) {
	uniq, err := funcsMap.which(name, params...)
	if err != nil {
		panic(err)
	}
	if uniq != this.exp {
		t.Fatalf("expected %v got %v", this.exp, uniq)
	}
}

func TestWhichStringEq(t *testing.T) {
	which{"stringEq"}.test(t, "eq", descriptor.FieldDescriptorProto_TYPE_STRING, descriptor.FieldDescriptorProto_TYPE_STRING)
}

func TestWhichInt64Eq(t *testing.T) {
	which{"int64Eq"}.test(t, "eq", descriptor.FieldDescriptorProto_TYPE_INT64, descriptor.FieldDescriptorProto_TYPE_INT64)
}

func TestWhichInt64Ge(t *testing.T) {
	which{"int64Ge"}.test(t, "ge", descriptor.FieldDescriptorProto_TYPE_INT64, descriptor.FieldDescriptorProto_TYPE_INT64)
}
