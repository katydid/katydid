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

package interp_test

import (
	"github.com/katydid/katydid/expr/ast"
	exprparser "github.com/katydid/katydid/expr/parser"
	. "github.com/katydid/katydid/funcs"
	. "github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/relapse/interp"
	"testing"
)

func getExpr(exprStr string) *expr.Expr {
	expr, err := exprparser.NewParser().ParseExpr(exprStr)
	if err != nil {
		panic(err)
	}
	return expr
}

func TestAnal1(t *testing.T) {
	f := Sprint(StringEq(StringVar(), StringConst("#")))
	f2 := Sprint(StringEq(StringVar(), StringConst("?")))
	alwaysFalse := getExpr("false")
	g := G{
		"main": In("A", In("B", Field("c", f))),
	}.Grammar()
	leafs := interp.GetLeafs(g, "A", "B", "c")
	if len(leafs) != 1 {
		t.Fatalf("expected %d leafs, but got %d", 1, len(leafs))
	}
	if !interp.CompareVarExpr(leafs[0], getExpr(f2)) {
		t.Fatalf("expected %s but got %s", f2, leafs[0].String())
	}
	newg := interp.Replace(g, leafs[0], alwaysFalse, "A", "B", "c")
	if interp.Satisfiable(newg) {
		t.Fatalf("should be unsatisfiable: %s", newg)
	}
	if !interp.Satisfiable(g) {
		t.Fatalf("should be satisfiable: %s", g)
	}
}

func TestAnal2(t *testing.T) {
	f := Sprint(StringEq(StringVar(), StringConst("#")))
	f2 := Sprint(StringEq(StringVar(), StringConst("?")))
	alwaysFalse := getExpr("false")
	g := G{
		"main": ExactAnyOf(
			In("A", ExactAllOf(
				In("B", Field("c", f)),
				Any()),
			),
			In("D", Field("c", f)),
		),
	}.Grammar()
	leafs := interp.GetLeafs(g, "A", "B", "c")
	if len(leafs) != 1 {
		t.Fatalf("expected %d leafs, but got %d", 1, len(leafs))
	}
	if !interp.CompareVarExpr(leafs[0], getExpr(f2)) {
		t.Fatalf("expected %s but got %s", f2, leafs[0].String())
	}
	newg := interp.Replace(g, leafs[0], alwaysFalse, "A", "B", "c")
	if !interp.Satisfiable(newg) {
		t.Fatalf("should be satisfiable: %s", newg)
	}
	if !interp.Satisfiable(g) {
		t.Fatalf("should be satisfiable: %s", g)
	}
}

func TestAnal3(t *testing.T) {
	f := Sprint(StringEq(StringVar(), StringConst("#")))
	f2 := Sprint(StringEq(StringVar(), StringConst("?")))
	alwaysFalse := getExpr("false")
	g := G{
		"main": ExactAnyOf(
			In("A", ExactAllOf(
				Many(In("B", Field("c", f))),
				In("B", Field("d", f)),
			)),
			In("D", Field("c", f)),
		),
	}.Grammar()
	leafs := interp.GetLeafs(g, "A", "B", "c")
	if len(leafs) != 1 {
		t.Fatalf("expected %d leafs, but got %d", 1, len(leafs))
	}
	if !interp.CompareVarExpr(leafs[0], getExpr(f2)) {
		t.Fatalf("expected %s but got %s", f2, leafs[0].String())
	}
	newg := interp.Replace(g, leafs[0], alwaysFalse, "A", "B", "c")
	if !interp.Satisfiable(newg) {
		t.Fatalf("should be satisfiable: %s", newg)
	}
	if !interp.Satisfiable(g) {
		t.Fatalf("should be satisfiable: %s", g)
	}
}

func TestAnal4(t *testing.T) {
	f := Sprint(StringEq(StringVar(), StringConst("#")))
	g := G{
		"main": ExactAnyOf(
			In("A", ExactAllOf(
				Many(In("B", Field("c", f))),
				In("B", ExactAnyOf(
					Field("d", f),
					Field("c", f),
				)),
			)),
			In("D", Field("c", f)),
		),
	}.Grammar()
	leafs := interp.GetLeafs(g, "A", "B", "c")
	if len(leafs) != 2 {
		t.Fatalf("expected %d leafs, but got %d", 2, len(leafs))
	}

}
