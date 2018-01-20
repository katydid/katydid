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
	"strings"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/parser/debug"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/funcs"
)

func TestComposeNot(t *testing.T) {
	expr := &ast.Expr{
		Function: &ast.Function{
			Name: "not",
			Params: []*ast.Expr{
				{
					Terminal: &ast.Terminal{
						BoolValue: proto.Bool(false),
					},
				},
			},
		},
	}
	b, err := NewBool(expr)
	if err != nil {
		t.Fatal(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		t.Fatal(err)
	}
	r, err := f.Eval(nil)
	if err != nil {
		t.Fatal(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	str := funcs.Sprint(f.(*composedBool).Func)
	if str != "->true" {
		t.Fatalf("trimming did not work: %s", str)
	}
}

func TestComposeContains(t *testing.T) {
	expr := ast.NewNestedFunction("contains",
		ast.NewNestedFunction("toLower", ast.NewStringVar()),
		ast.NewNestedFunction("toLower", ast.NewStringConst("TheStreet")),
	)
	b, err := NewBool(expr)
	if err != nil {
		t.Fatal(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		t.Fatal(err)
	}
	r, err := f.Eval(debug.NewStringValue("TheStreet"))
	if err != nil {
		t.Fatal(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	r, err = f.Eval(debug.NewStringValue("ThatStreet"))
	if err != nil {
		t.Fatal(err)
	}
	if r != false {
		t.Fatalf("expected false")
	}
	if strings.Contains(funcs.Sprint(f.(*composedBool).Func), "toLower(`TheStreet`)") {
		t.Fatalf("trimming did not work")
	}
}

func TestComposeStringEq(t *testing.T) {
	expr := ast.NewNestedFunction("eq",
		ast.NewNestedFunction("toLower", ast.NewStringVar()),
		ast.NewNestedFunction("toLower", ast.NewStringConst("TheStreet")),
	)
	b, err := NewBool(expr)
	if err != nil {
		t.Fatal(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		t.Fatal(err)
	}
	r, err := f.Eval(debug.NewStringValue("TheStreet"))
	if err != nil {
		t.Fatal(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestComposeListBool(t *testing.T) {
	expr := ast.NewNestedFunction("eq",
		ast.NewNestedFunction("length",
			ast.NewBoolList(
				ast.NewTrue(),
				ast.NewFalse(),
			),
		),
		ast.NewIntConst(2),
	)
	b, err := NewBool(expr)
	if err != nil {
		t.Fatal(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		t.Fatal(err)
	}
	r, err := f.Eval(nil)
	if err != nil {
		t.Fatal(err)
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
	expr := ast.NewNestedFunction("eq",
		ast.NewNestedFunction("elem",
			ast.NewNestedFunction("print",
				ast.NewIntList(
					ast.NewIntConst(1),
					ast.NewIntConst(2),
				),
			),
			ast.NewIntConst(1),
		),
		ast.NewIntConst(2),
	)
	b, err := NewBool(expr)
	if err != nil {
		t.Fatal(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		t.Fatal(err)
	}
	r, err := f.Eval(nil)
	if err != nil {
		t.Fatal(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
	t.Logf("%s", funcs.Sprint(f.(*composedBool).Func))
}

func TestComposeRegex(t *testing.T) {
	expr := ast.NewNestedFunction("regex",
		ast.NewStringConst("ab"),
		ast.NewStringVar(),
	)
	b, err := NewBool(expr)
	if err != nil {
		t.Fatal(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		t.Fatal(err)
	}
	r, err := f.Eval(debug.NewStringValue("a"))
	if err != nil {
		t.Fatal(err)
	}
	if r != false {
		t.Fatalf("expected false")
	}
}

func TestConst(t *testing.T) {
	expr := ast.NewNestedFunction("regex",
		ast.NewStringVar(),
		ast.NewStringConst("ab"),
	)
	b, err := NewBool(expr)
	if err != nil {
		t.Fatal(err)
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
	expr := ast.NewNestedFunction("regex",
		ast.NewStringConst(".*"),
		ast.NewStringConst("ab"),
	)
	b, err := NewBool(expr)
	if err != nil {
		t.Fatal(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		t.Fatal(err)
	}
	r, err := f.Eval(nil)
	if err != nil {
		t.Fatal(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestNoTrim(t *testing.T) {
	expr := ast.NewNestedFunction("eq",
		ast.NewNestedFunction("elem",
			ast.NewStringList(
				ast.NewNestedFunction("print",
					ast.NewStringVar(),
				),
			),
			ast.NewIntConst(0),
		),
		ast.NewStringConst("abc"),
	)
	b, err := NewBool(expr)
	if err != nil {
		t.Fatal(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		t.Fatal(err)
	}
	str := funcs.Sprint(f.(*composedBool).Func)
	if str == "false" {
		t.Fatalf("too much trimming")
	}
	t.Logf("trimmed = %s", str)
	r, err := f.Eval(debug.NewStringValue("abc"))
	if err != nil {
		t.Fatal(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestList(t *testing.T) {
	expr := ast.NewNestedFunction("eq",
		ast.NewNestedFunction("elem",
			ast.NewStringList(
				ast.NewStringConst("abc"),
			),
			ast.NewIntConst(0),
		),
		ast.NewStringConst("abc"),
	)
	b, err := NewBool(expr)
	if err != nil {
		t.Fatal(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		t.Fatal(err)
	}
	str := funcs.Sprint(f.(*composedBool).Func)
	if str != "->true" {
		t.Fatalf("not enough trimming on %s", str)
	}
	r, err := f.Eval(nil)
	if err != nil {
		t.Fatal(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestComposeBuiltInRegex(t *testing.T) {
	expr := ast.NewRegex(
		ast.NewStringConst("ab"),
	)
	b, err := NewBool(expr)
	if err != nil {
		t.Fatal(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		t.Fatal(err)
	}
	r, err := f.Eval(debug.NewStringValue("ab"))
	if err != nil {
		t.Fatal(err)
	}
	if r != true {
		t.Fatalf("expected true")
	}
}

func TestComposeBuiltInEqual(t *testing.T) {
	expr := ast.NewEqual(
		ast.NewIntConst(123),
	)
	b, err := NewBool(expr)
	if err != nil {
		t.Fatal(err)
	}
	f, err := NewBoolFunc(b)
	if err != nil {
		t.Fatal(err)
	}
	r, err := f.Eval(debug.NewIntValue(124))
	if err != nil {
		t.Fatal(err)
	}
	if r != false {
		t.Fatalf("expected false")
	}
}
