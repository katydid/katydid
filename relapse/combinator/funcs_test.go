//  Copyright 2018 Walter Schulze
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

package combinator

import (
	"math/rand"
	"testing"
	"time"

	"github.com/katydid/katydid/relapse/ast"
)

func TestConst(t *testing.T) {
	v := Value(BoolConst(true))
	if !v.GetLeafNode().GetExpr().GetTerminal().GetBoolValue() {
		t.Fatalf("expected true")
	}

	vs := Value(BoolsConst([]bool{true}))
	if !vs.GetLeafNode().GetExpr().GetList().GetElems()[0].GetTerminal().GetBoolValue() {
		t.Fatalf("expected true")
	}
}

func TestInt(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var intFunc func(*ast.Expr, *ast.Expr) *ast.Expr
	switch rand.Intn(3) {
	case 0:
		intFunc = Eq
	case 1:
		intFunc = GE
	case 2:
		intFunc = LE
	}
	t.Log(Value(ast.NewFunction("print", intFunc(IntVar(), IntConst(rand.Int63())))).String())
}

// TestImmutability tests that the `Value` function does not change the input expression and remove a comma.
func TestImmutability(t *testing.T) {
	eq := Eq(UintVar(), UintConst(10))
	s1 := eq.String()
	_ = Value(eq)
	s2 := eq.String()
	if s1 != s2 {
		t.Fatalf("s1: %v, s2:%v\n", s1, s2)
	}
}
