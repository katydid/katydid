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

package compose

import (
	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/parser/debug"
	"github.com/katydid/katydid/relapse"
	"github.com/katydid/katydid/relapse/funcs"
	"strings"
	"testing"
)

func TestComposeNot(t *testing.T) {
	expr := &relapse.Expr{
		Function: &relapse.Function{
			Name: "not",
			Params: []*relapse.Expr{
				{
					Terminal: &relapse.Terminal{
						BoolValue: proto.Bool(false),
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	r, err := f.Eval(nil)
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	str := funcs.Sprint(f.(*composedBool).Func)
	if str != "->true" {
		t.Fatalf("trimming did not work: %s", str)
	}
}

//contains(toLower(decString(test.Address.Street.value)), toLower("TheStreet"))

func TestComposeContains(t *testing.T) {
	expr := relapse.NewNestedFunction("contains",
		relapse.NewNestedFunction("toLower", relapse.NewStringVar()),
		relapse.NewNestedFunction("toLower", relapse.NewStringConst("TheStreet")),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	r, err := f.Eval(debug.NewStringValue("TheStreet"))
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	r, err = f.Eval(debug.NewStringValue("ThatStreet"))
	if err != nil {
		panic(err)
	}
	if r != false {
		t.Fatalf("expected false")
	}
	if strings.Contains(funcs.Sprint(f.(*composedBool).Func), "toLower(`TheStreet`)") {
		t.Fatalf("trimming did not work")
	}
}

func TestComposeStringEq(t *testing.T) {
	expr := relapse.NewNestedFunction("eq",
		relapse.NewNestedFunction("toLower", relapse.NewStringVar()),
		relapse.NewNestedFunction("toLower", relapse.NewStringConst("TheStreet")),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	r, err := f.Eval(debug.NewStringValue("TheStreet"))
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestComposeListBool(t *testing.T) {
	expr := relapse.NewNestedFunction("eq",
		relapse.NewNestedFunction("length",
			relapse.NewBoolList(
				relapse.NewTrue(),
				relapse.NewFalse(),
			),
		),
		relapse.NewIntConst(2),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	r, err := f.Eval(nil)
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	str := funcs.Sprint(f.(*composedBool).Func)
	if str != "->true" {
		t.Fatalf("trimming did not work: %s", str)
	}
}

func TestComposeListInt64(t *testing.T) {
	expr := relapse.NewNestedFunction("eq",
		relapse.NewNestedFunction("elem",
			relapse.NewNestedFunction("print",
				relapse.NewIntList(
					relapse.NewIntConst(1),
					relapse.NewIntConst(2),
				),
			),
			relapse.NewIntConst(1),
		),
		relapse.NewIntConst(2),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	r, err := f.Eval(nil)
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	t.Logf("%s", funcs.Sprint(f.(*composedBool).Func))
}

func TestComposeRegex(t *testing.T) {
	expr := relapse.NewNestedFunction("regex",
		relapse.NewStringConst("ab"),
		relapse.NewStringVar(),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	r, err := f.Eval(debug.NewStringValue("a"))
	if err != nil {
		panic(err)
	}
	if r != false {
		t.Fatalf("expected false")
	}
}

func TestConst(t *testing.T) {
	expr := relapse.NewNestedFunction("regex",
		relapse.NewStringVar(),
		relapse.NewStringConst("ab"),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	_, err = NewBoolFunc(b)
	if err == nil {
		t.Fatalf("expected error")
	}
	if !strings.Contains(err.Error(), "regex has constant") || !strings.Contains(err.Error(), "has a variable parameter") {
		t.Fatalf("expected more specific error %s", err.Error())
	}
}

func TestTrimInit(t *testing.T) {
	expr := relapse.NewNestedFunction("regex",
		relapse.NewStringConst(".*"),
		relapse.NewStringConst("ab"),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	r, err := f.Eval(nil)
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestNoTrim(t *testing.T) {
	expr := relapse.NewNestedFunction("eq",
		relapse.NewNestedFunction("elem",
			relapse.NewStringList(
				relapse.NewNestedFunction("print",
					relapse.NewStringVar(),
				),
			),
			relapse.NewIntConst(0),
		),
		relapse.NewStringConst("abc"),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	str := funcs.Sprint(f.(*composedBool).Func)
	if str == "false" {
		t.Fatalf("too much trimming")
	}
	t.Logf("trimmed = %s", str)
	r, err := f.Eval(debug.NewStringValue("abc"))
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestList(t *testing.T) {
	expr := relapse.NewNestedFunction("eq",
		relapse.NewNestedFunction("elem",
			relapse.NewStringList(
				relapse.NewStringConst("abc"),
			),
			relapse.NewIntConst(0),
		),
		relapse.NewStringConst("abc"),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	str := funcs.Sprint(f.(*composedBool).Func)
	if str != "->true" {
		t.Fatalf("not enough trimming on %s", str)
	}
	r, err := f.Eval(nil)
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestComposeBuiltInRegex(t *testing.T) {
	expr := relapse.NewRegex(
		relapse.NewStringConst("ab"),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	r, err := f.Eval(debug.NewStringValue("ab"))
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestComposeBuiltInEqual(t *testing.T) {
	expr := relapse.NewEqual(
		relapse.NewIntConst(123),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	r, err := f.Eval(debug.NewIntValue(124))
	if err != nil {
		panic(err)
	}
	if r != false {
		t.Fatalf("expected false")
	}
}
