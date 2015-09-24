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

package convert

import (
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/interp"
	//"github.com/katydid/katydid/serialize/debug"
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/tests"
	"testing"
)

func child(t *testing.T, refs relapse.RefLookup, name *relapse.NameExpr, expected *relapse.Pattern) {
	t.Logf("input = %s", refs["main"])
	t.Logf("name = %s", name)
	dun := deriv(refs, refs["main"], nil, name)
	du := dun.(*callAndReturn)
	t.Logf("child = %s", du.child)
	if du.child != nil {
		child := interp.Simplify(refs, du.child)
		t.Logf("simplified child = %s", child)
		if !child.Equal(expected) {
			t.Fatalf("expected %s\n, but got %s", expected, child)
		}
	} else if expected != nil {
		t.Fatalf("expected %s\n, but got <nil>", expected)
	}
}

func TestChild1(t *testing.T) {
	t.Skip("TODO")
	g := tests.MagazineFrameAFinanceJudo.Grammar()
	refs := relapse.NewRefsLookup(g)
	expected := refs["main"].Concat.LeftPattern.TreeNode.Pattern
	_, names := getAllReachableExprs(refs, refs["main"])
	child(t, refs, names[0], expected)
}

func TestChild2(t *testing.T) {
	t.Skip("TODO")
	g := tests.ContextPerson.Grammar()
	refs := relapse.NewRefsLookup(g)
	expected := refs["main"].Concat.RightPattern.Concat.LeftPattern.TreeNode.Pattern
	_, names := getAllReachableExprs(refs, refs["main"])
	child(t, refs, names[0], expected)
}

func TestChild3(t *testing.T) {
	g := tests.RecursiveSrcTree.Grammar()
	refs := relapse.NewRefsLookup(g)
	expected := refs["main"].Concat.RightPattern.Concat.LeftPattern.Or.LeftPattern.TreeNode.Pattern
	_, names := getAllReachableExprs(refs, refs["main"])
	child(t, refs, names[0], expected)
}

func TestChild4(t *testing.T) {
	g := tests.AndNameTelephonePerson.Grammar()
	refs := relapse.NewRefsLookup(g)
	expected := refs["main"].And.LeftPattern.Concat.RightPattern.Concat.LeftPattern.TreeNode.Pattern
	_, names := getAllReachableExprs(refs, refs["main"])
	child(t, refs, names[0], expected)
}

func TestChild5(t *testing.T) {
	p := combinator.MatchInOrder(combinator.Any(), combinator.MatchField("Telephone", funcs.Sprint(funcs.StringVarEq(funcs.StringConst("01234")))))
	refs := relapse.RefLookup{"main": p}
	name := relapse.NewName("David")
	child(t, refs, name, nil)
}

func DisabledTestChild6(t *testing.T) {
	t.Skip("TODO")
	g := tests.EmptyOrNilPerson.Grammar()
	refs := relapse.NewRefsLookup(g)
	expected := refs["empty"].Concat.RightPattern.Concat.LeftPattern.TreeNode.Pattern
	_, names := getAllReachableExprs(refs, refs["main"])
	child(t, refs, names[0], expected)
}

func TestChild7(t *testing.T) {
	t.Skip("TODO")
	g := tests.EmptyOrNilPerson.Grammar()
	refs := relapse.NewRefsLookup(g)
	_, names := getAllReachableExprs(refs, refs["main"])
	t.Logf("input = %s", refs["main"])
	dun := deriv(refs, refs["main"], nil, names[0])
	du := dun.(*callAndReturn)
	t.Logf("name = %s", names[0])
	t.Logf("null  = %s", du.retNull)
	t.Logf("else  = %s", du.retElse)
	snull := interp.Simplify(refs, du.retNull)
	selse := interp.Simplify(refs, du.retElse)
	t.Logf("snull = %s", snull)
	t.Logf("selse = %s", selse)
	if !snull.Equal(relapse.NewNot(relapse.NewEmptySet())) {
		t.Fatalf("for dst expected !EmptySet, but got %s", snull)
	}
	if !selse.Equal(relapse.NewReference("empty")) {
		t.Fatalf("for fail expected #empty, but got %s", selse)
	}
}
