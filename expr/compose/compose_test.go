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
	"github.com/katydid/katydid/expr"
	"github.com/katydid/katydid/expr/funcs"
	"github.com/katydid/katydid/parser"
	"strings"
	"testing"
)

func TestComposeNot(t *testing.T) {
	expr := &expr.Expr{
		Function: &expr.Function{
			Name: "not",
			Params: []*expr.Expr{
				{
					Terminal: &expr.Terminal{
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
	if str != "true" {
		t.Fatalf("trimming did not work: %s", str)
	}
}

//contains(toLower(decString(test.Address.Street.value)), toLower("TheStreet"))

func TestComposeContains(t *testing.T) {
	expr := expr.NewNestedFunction("contains",
		expr.NewNestedFunction("toLower", expr.NewStringVar()),
		expr.NewNestedFunction("toLower", expr.NewStringConst("TheStreet")),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	r, err := f.Eval(parser.NewStringValue("TheStreet"))
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	r, err = f.Eval(parser.NewStringValue("ThatStreet"))
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
	expr := expr.NewNestedFunction("eq",
		expr.NewNestedFunction("toLower", expr.NewStringVar()),
		expr.NewNestedFunction("toLower", expr.NewStringConst("TheStreet")),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	r, err := f.Eval(parser.NewStringValue("TheStreet"))
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestComposeListBool(t *testing.T) {
	expr := expr.NewNestedFunction("eq",
		expr.NewNestedFunction("length",
			expr.NewBoolList(
				expr.NewTrue(),
				expr.NewFalse(),
			),
		),
		expr.NewIntConst(2),
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
	if str != "true" {
		t.Fatalf("trimming did not work: %s", str)
	}
}

func TestComposeListInt64(t *testing.T) {
	expr := expr.NewNestedFunction("eq",
		expr.NewNestedFunction("elem",
			expr.NewNestedFunction("print",
				expr.NewIntList(
					expr.NewIntConst(1),
					expr.NewIntConst(2),
				),
			),
			expr.NewIntConst(1),
		),
		expr.NewIntConst(2),
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
	expr := expr.NewNestedFunction("regex",
		expr.NewStringConst("ab"),
		expr.NewStringVar(),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	r, err := f.Eval(parser.NewStringValue("a"))
	if err != nil {
		panic(err)
	}
	if r != false {
		t.Fatalf("expected false")
	}
}

func TestConst(t *testing.T) {
	expr := expr.NewNestedFunction("regex",
		expr.NewStringVar(),
		expr.NewStringConst("ab"),
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
	expr := expr.NewNestedFunction("regex",
		expr.NewStringConst(".*"),
		expr.NewStringConst("ab"),
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
	expr := expr.NewNestedFunction("eq",
		expr.NewNestedFunction("elem",
			expr.NewStringList(
				expr.NewNestedFunction("print",
					expr.NewStringVar(),
				),
			),
			expr.NewIntConst(0),
		),
		expr.NewStringConst("abc"),
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
	r, err := f.Eval(parser.NewStringValue("abc"))
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestList(t *testing.T) {
	expr := expr.NewNestedFunction("eq",
		expr.NewNestedFunction("elem",
			expr.NewStringList(
				expr.NewStringConst("abc"),
			),
			expr.NewIntConst(0),
		),
		expr.NewStringConst("abc"),
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
	if str != "true" {
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
	expr := expr.NewRegex(
		expr.NewStringConst("ab"),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	r, err := f.Eval(parser.NewStringValue("ab"))
	if err != nil {
		panic(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestComposeBuiltInEqual(t *testing.T) {
	expr := expr.NewEqual(
		expr.NewIntConst(123),
	)
	b, err := NewBool(expr)
	if err != nil {
		panic(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		panic(err)
	}
	r, err := f.Eval(parser.NewIntValue(124))
	if err != nil {
		panic(err)
	}
	if r != false {
		t.Fatalf("expected false")
	}
}
