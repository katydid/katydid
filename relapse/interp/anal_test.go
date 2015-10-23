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
	. "github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/relapse/ast"
	. "github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/relapse/interp"
	"testing"
)

func getExpr(pattern *relapse.Pattern) *expr.Expr {
	return pattern.GetLeafNode().GetExpr()
}

func TestAnal1(t *testing.T) {
	f := Value(StringEq(StringVar(), StringConst("#")))
	f2 := Value(StringEq(StringVar(), StringConst("?")))
	alwaysFalse := getExpr(Value(BoolConst(false)))
	g := G{
		"main": In("A", In("B", In("c", f))),
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
	f := Value(StringEq(StringVar(), StringConst("#")))
	f2 := Value(StringEq(StringVar(), StringConst("?")))
	alwaysFalse := getExpr(Value(BoolConst(false)))
	g := G{
		"main": AnyOf(
			In("A", AllOf(
				In("B", In("c", f)),
				Any()),
			),
			In("D", In("c", f)),
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
	f := Value(StringEq(StringVar(), StringConst("#")))
	f2 := Value(StringEq(StringVar(), StringConst("?")))
	alwaysFalse := getExpr(Value(BoolConst(false)))
	g := G{
		"main": AnyOf(
			In("A", AllOf(
				Many(In("B", In("c", f))),
				In("B", In("d", f)),
			)),
			In("D", In("c", f)),
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
	f := Value(StringEq(StringVar(), StringConst("#")))
	g := G{
		"main": AnyOf(
			In("A", AllOf(
				Many(In("B", In("c", f))),
				In("B", AnyOf(
					In("d", f),
					In("c", f),
				)),
			)),
			In("D", In("c", f)),
		),
	}.Grammar()
	leafs := interp.GetLeafs(g, "A", "B", "c")
	if len(leafs) != 2 {
		t.Fatalf("expected %d leafs, but got %d", 2, len(leafs))
	}

}
