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
	"github.com/katydid/katydid/relapse/ast"
	. "github.com/katydid/katydid/relapse/combinator"
	. "github.com/katydid/katydid/relapse/funcs"
	"github.com/katydid/katydid/relapse/interp"
	"testing"
)

var alwaysFalse = Value(BoolConst(false))
var f = Value(StringEq(StringVar(), StringConst("#")))
var f2 = Value(StringEq(StringVar(), StringConst("?")))

func getExpr(pattern *ast.Pattern) *ast.Expr {
	return pattern.GetLeafNode().GetExpr()
}

func TestAnal1(t *testing.T) {
	g := G{
		"main": In("A", In("B", In("c", f))),
	}.Grammar()
	newg := g.Clone()
	path := []*ast.NameExpr{ast.NewStringName("A"), ast.NewStringName("B"), ast.NewStringName("c")}
	leafs := interp.ChildrenOf(newg, path)
	leafs = interp.LeafNodesOf(newg, leafs...)
	if len(leafs) != 1 {
		t.Fatalf("expected %d leafs, but got %d", 1, len(leafs))
	}
	if !interp.CompareVarExpr(leafs[0].LeafNode.Expr, getExpr(f2)) {
		t.Fatalf("expected %s but got %s", f2, leafs[0].String())
	}
	leafs[0].LeafNode = alwaysFalse.LeafNode
	if interp.Satisfiable(newg) {
		t.Fatalf("should be unsatisfiable: %s", newg)
	}
	if !interp.Satisfiable(g) {
		t.Fatalf("should be satisfiable: %s", g)
	}
}

func TestAnal2(t *testing.T) {
	g := G{
		"main": AnyOf(
			In("A", AllOf(
				In("B", In("c", f)),
				Any()),
			),
			In("D", In("c", f)),
		),
	}.Grammar()
	newg := g.Clone()
	path := []*ast.NameExpr{ast.NewStringName("A"), ast.NewStringName("B"), ast.NewStringName("c")}
	leafs := interp.ChildrenOf(newg, path)
	leafs = interp.LeafNodesOf(newg, leafs...)
	if len(leafs) != 1 {
		t.Fatalf("expected %d leafs, but got %d", 1, len(leafs))
	}
	if !interp.CompareVarExpr(leafs[0].LeafNode.Expr, getExpr(f2)) {
		t.Fatalf("expected %s but got %s", f2, leafs[0].String())
	}
	leafs[0].LeafNode = alwaysFalse.LeafNode
	if !interp.Satisfiable(newg) {
		t.Fatalf("should be satisfiable: %s", newg)
	}
	if !interp.Satisfiable(g) {
		t.Fatalf("should be satisfiable: %s", g)
	}
}

func TestAnal3(t *testing.T) {
	g := G{
		"main": AnyOf(
			In("A", AllOf(
				Many(In("B", In("c", f))),
				In("B", In("d", f)),
			)),
			In("D", In("c", f)),
		),
	}.Grammar()
	newg := g.Clone()
	path := []*ast.NameExpr{ast.NewStringName("A"), ast.NewStringName("B"), ast.NewStringName("c")}
	leafs := interp.ChildrenOf(newg, path)
	leafs = interp.LeafNodesOf(newg, leafs...)
	if len(leafs) != 1 {
		t.Fatalf("expected %d leafs, but got %d", 1, len(leafs))
	}
	if !interp.CompareVarExpr(leafs[0].LeafNode.Expr, getExpr(f2)) {
		t.Fatalf("expected %s but got %s", f2, leafs[0].String())
	}
	leafs[0].LeafNode = alwaysFalse.LeafNode
	if !interp.Satisfiable(newg) {
		t.Fatalf("should be satisfiable: %s", newg)
	}
	if !interp.Satisfiable(g) {
		t.Fatalf("should be satisfiable: %s", g)
	}
}

func TestAnal4(t *testing.T) {
	g := G{
		"main": AllOf(
			In("A", AnyOf(
				In("B", In("c", f)),
				In("B", AllOf(
					In("d", f),
					In("c", f),
				)),
			)),
			In("D", In("c", f)),
		),
	}.Grammar()
	newg := g.Clone()
	path := []*ast.NameExpr{ast.NewStringName("A"), ast.NewStringName("B"), ast.NewStringName("c")}
	leafs := interp.ChildrenOf(newg, path)
	leafs = interp.LeafNodesOf(newg, leafs...)
	if len(leafs) != 2 {
		t.Fatalf("expected %d leafs, but got %d", 2, len(leafs))
	}
	for i := range leafs {
		if !interp.CompareVarExpr(leafs[i].LeafNode.Expr, getExpr(f2)) {
			t.Fatalf("expected %s but got %s", f2, leafs[0].String())
		}
		leafs[i].LeafNode = alwaysFalse.LeafNode
	}
	if interp.Satisfiable(newg) {
		t.Fatalf("should be unsatisfiable: %s", newg)
	}
	if !interp.Satisfiable(g) {
		t.Fatalf("should be satisfiable: %s", g)
	}
}

func TestAnal5Repeated(t *testing.T) {
	g := G{
		"main": In("A", InAny(In("c", InAny(f)))),
	}.Grammar()
	newg := g.Clone()
	path := []*ast.NameExpr{ast.NewStringName("A"), ast.NewAnyName(), ast.NewStringName("c"), ast.NewAnyName()}
	leafs := interp.ChildrenOf(newg, path)
	leafs = interp.LeafNodesOf(newg, leafs...)
	if len(leafs) != 1 {
		t.Fatalf("expected %d leafs, but got %d", 1, len(leafs))
	}
	if !interp.CompareVarExpr(leafs[0].LeafNode.Expr, getExpr(f2)) {
		t.Fatalf("expected %s but got %s", f2, leafs[0].String())
	}
	leafs[0].LeafNode = alwaysFalse.LeafNode
	if interp.Satisfiable(newg) {
		t.Fatalf("should be unsatisfiable: %s", newg)
	}
	if !interp.Satisfiable(g) {
		t.Fatalf("should be satisfiable: %s", g)
	}
}

func TestAnal6TreeNode(t *testing.T) {
	g := G{
		"main": In("A", In("B", In("c", f))),
	}.Grammar()
	newg := g.Clone()
	path := []*ast.NameExpr{ast.NewStringName("A")}
	trees := interp.ChildrenOf(newg, path)
	trees = interp.TreeNodesOf(newg, trees...)
	if len(trees) != 1 {
		t.Fatalf("expected %d trees, but got %d", 1, len(trees))
	}
	trees[0].TreeNode.Name = ast.NewAnyNameExcept(ast.NewAnyName())
	if interp.Satisfiable(newg) {
		t.Fatalf("should be unsatisfiable: %s", newg)
	}
	if !interp.Satisfiable(g) {
		t.Fatalf("should be satisfiable: %s", g)
	}
}

func TestAnal7TreeNodes(t *testing.T) {
	g := G{
		"main": AllOf(
			In("A", AnyOf(
				In("B", In("c", f)),
				In("B", AllOf(
					In("d", f),
					In("c", f),
				)),
			)),
			In("D", In("c", f)),
		),
	}.Grammar()
	newg := g.Clone()
	path := []*ast.NameExpr{ast.NewStringName("A")}
	trees := interp.ChildrenOf(newg, path)
	trees = interp.TreeNodesOf(newg, trees...)
	if len(trees) != 2 {
		t.Fatalf("expected %d trees, but got %d", 2, len(trees))
	}
	for i := range trees {
		trees[i].TreeNode.Name = ast.NewAnyNameExcept(ast.NewAnyName())
	}
	if interp.Satisfiable(newg) {
		t.Fatalf("should be unsatisfiable: %s", newg)
	}
	if !interp.Satisfiable(g) {
		t.Fatalf("should be satisfiable: %s", g)
	}
}

func TestAnal8TreeNodes2(t *testing.T) {
	g := G{
		"main": AllOf(
			In("A", AnyOf(
				In("B", In("c", f)),
				In("C", AllOf(
					In("d", f),
					In("c", f),
				)),
			)),
			In("D", In("c", f)),
		),
	}.Grammar()
	newg := g.Clone()
	path := []*ast.NameExpr{ast.NewStringName("A")}
	trees := interp.ChildrenOf(newg, path)
	trees = interp.TreeNodesOf(newg, trees...)
	if len(trees) != 2 {
		t.Fatalf("expected %d trees, but got %d", 2, len(trees))
	}
	for i := range trees {
		if trees[i].TreeNode.Name.Equal(ast.NewStringName("B")) {
			trees[i].TreeNode.Name = ast.NewAnyNameExcept(ast.NewAnyName())
		}
	}
	if !interp.Satisfiable(newg) {
		t.Fatalf("should be satisfiable: %s", newg)
	}
	if !interp.Satisfiable(g) {
		t.Fatalf("should be satisfiable: %s", g)
	}
	newerg := interp.SimplifyGrammar(newg)
	t.Logf("%v", newerg)
	trees = interp.ChildrenOf(newerg, path)
	trees = interp.TreeNodesOf(newerg, trees...)
	if len(trees) != 1 {
		t.Fatalf("expected %d trees, but got %d", 1, len(trees))
	}
}
