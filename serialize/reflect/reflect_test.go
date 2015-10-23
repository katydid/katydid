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

package reflect

import (
	"github.com/katydid/katydid/serialize/debug"
	"reflect"
	"testing"
)

type TheStruct struct {
	A int64
	B []string
	C *TheStruct
	D *int32
	E []*TheStruct
}

func TestReflect(t *testing.T) {
	p := NewReflectParser()
	three := int32(3)
	four := int32(4)
	s := &TheStruct{A: int64(1), B: []string{"b2", "b3"}, C: &TheStruct{A: int64(2), D: &three, E: []*TheStruct{&TheStruct{B: []string{"b4"}}, &TheStruct{B: []string{"b5"}}}}, D: &four}
	p.Init(reflect.ValueOf(s))
	m := debug.Walk(p)
	t.Logf("%v", m)
	if len(m[1].Children) != 2 {
		t.Fatalf("expected list of length 2")
	}
	if m[2].Children[2].Children[0].Label != "\"0\"" {
		t.Fatalf("expected child in slices of structs got %v", m[2].Children[2].Children[0].Label)
	}
}
