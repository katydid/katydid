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
	. "github.com/katydid/katydid/lang/combinator"
	"github.com/katydid/katydid/lang/interp"
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
	f := `eq($string, "#")`
	f2 := `eq($string, "?")`
	g := G{
		"main": MatchIn("A", MatchIn("B", MatchField("c", f))),
	}.Grammar()
	leafs := interp.GetLeafs(g, "A", "B", "c")
	if len(leafs) != 1 {
		t.Fatalf("expected %d leafs, but got %d", 1, len(leafs))
	}
	if !interp.CompareVarExpr(leafs[0], getExpr(f2)) {
		t.Fatalf("expected %s but got %s", f2, leafs[0].String())
	}
}

func TestAnal2(t *testing.T) {
	f := `eq($string, "#")`
	f2 := `eq($string, "?")`
	g := G{
		"main": Or(MatchIn("A", And(MatchIn("B", MatchField("c", f)), Any())), MatchIn("D", MatchField("c", f))),
	}.Grammar()
	leafs := interp.GetLeafs(g, "A", "B", "c")
	if len(leafs) != 1 {
		t.Fatalf("expected %d leafs, but got %d", 1, len(leafs))
	}
	if !interp.CompareVarExpr(leafs[0], getExpr(f2)) {
		t.Fatalf("expected %s but got %s", f2, leafs[0].String())
	}
}

func TestAnal3(t *testing.T) {
	f := `eq($string, "#")`
	f2 := `eq($string, elem([]string{"?"}, int(0)))`
	g := G{
		"main": Or(
			MatchIn("A", And(
				Many(MatchIn("B", MatchField("c", f))),
				MatchIn("B", MatchField("d", f)),
			)),
			MatchIn("D", MatchField("c", f)),
		),
	}.Grammar()
	leafs := interp.GetLeafs(g, "A", "B", "c")
	if len(leafs) != 1 {
		t.Fatalf("expected %d leafs, but got %d", 1, len(leafs))
	}
	if !interp.CompareVarExpr(leafs[0], getExpr(f2)) {
		t.Fatalf("expected %s but got %s", f2, leafs[0].String())
	}
}

func TestAnal4(t *testing.T) {
	f := `eq($string, "#")`
	g := G{
		"main": Or(
			MatchIn("A", And(
				Many(MatchIn("B", MatchField("c", f))),
				MatchIn("B", Or(
					MatchField("d", f),
					MatchField("c", f),
				)),
			)),
			MatchIn("D", MatchField("c", f)),
		),
	}.Grammar()
	leafs := interp.GetLeafs(g, "A", "B", "c")
	if len(leafs) != 2 {
		t.Fatalf("expected %d leafs, but got %d", 2, len(leafs))
	}
}
