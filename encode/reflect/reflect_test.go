//  Copyright 2016 Walter Schulze
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
	reflectparser "github.com/katydid/katydid/parser/reflect"
	"reflect"
	"testing"
)

func testCopy(t *testing.T, input interface{}) {
	p := reflectparser.NewReflectParser().Init(reflect.ValueOf(input))
	output := reflect.New(reflect.TypeOf(input).Elem()).Interface()
	if err := Encode(p, output); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(input, output) {
		t.Fatalf("expected equal, but got %#v", output)
	}
}

func testPartial(t *testing.T, input interface{}, expected interface{}) {
	p := reflectparser.NewReflectParser().Init(reflect.ValueOf(input))
	output := reflect.New(reflect.TypeOf(expected).Elem()).Interface()
	if err := Encode(p, output); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(expected, output) {
		t.Fatalf("expected equal, but got %#v", output)
	}
}

type A struct {
	Name  string
	Value int64
}

func TestCopyA(t *testing.T) {
	testCopy(t, &A{Name: "a", Value: 1})
	testCopy(t, &A{Name: "a"})
	testCopy(t, &A{Value: 1})
	testCopy(t, &A{})
}

func TestParialA(t *testing.T) {
	testPartial(t, &A{Name: "a", Value: 1}, &B{Name: "a"})
}

type B struct {
	Name string
	A    *A
}

func TestCopyB(t *testing.T) {
	testCopy(t, &B{Name: "b", A: &A{Name: "a", Value: 1}})
	testCopy(t, &B{Name: "b"})
}

func TestParialB(t *testing.T) {
	testPartial(t, &B{Name: "b", A: &A{Name: "a", Value: 1}}, &A{Name: "b"})
}

type C struct {
	Name string
	List []int64
}

func TestCopyC(t *testing.T) {
	testCopy(t, &C{Name: "c", List: []int64{1, 2, 3, 4}})
	testCopy(t, &C{Name: "c"})
}

func TestParialC(t *testing.T) {
	testPartial(t, &C{Name: "c", List: []int64{1, 2, 3, 4}}, &A{Name: "c"})
}

type D struct {
	Name string
	List []*A
}

func TestCopyD(t *testing.T) {
	testCopy(t, &D{Name: "d", List: []*A{&A{Name: "a", Value: 1}, &A{}, nil}})
}
