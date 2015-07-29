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
	"github.com/katydid/katydid/tests"
	"testing"
)

func TestConvert1(t *testing.T) {
	g := tests.MagazineFrameAFinanceJudo.Grammar()
	refs := relapse.NewRefsLookup(g)
	t.Logf("input = %s", refs["main"])
	childPattern := refs["main"].Concat.LeftPattern.TreeNode.Pattern
	_, names := getAllReachableExprs(refs, refs["main"])
	name := names[0]
	t.Logf("name = %s", name)
	du := deriv(refs, refs["main"], nil, name)
	child := interp.Simplify(refs, du.child)
	t.Logf("child = %s", child)
	if !child.Equal(childPattern) {
		t.Fatalf("expected %s\n, but got %s", childPattern, child)
	}
}

func TestConvert2(t *testing.T) {
	g := tests.ContextPerson.Grammar()
	refs := relapse.NewRefsLookup(g)
	t.Logf("input = %s", refs["main"])
	childPattern := refs["main"].Concat.RightPattern.Concat.LeftPattern.TreeNode.Pattern
	_, names := getAllReachableExprs(refs, refs["main"])
	name := names[0]
	t.Logf("name = %s", name)
	du := deriv(refs, refs["main"], nil, name)
	child := interp.Simplify(refs, du.child)
	t.Logf("child = %s", child)
	if !child.Equal(childPattern) {
		t.Fatalf("expected %s\n, but got %s", childPattern, child)
	}
}

func TestConvert3(t *testing.T) {
	g := tests.RecursiveSrcTree.Grammar()
	refs := relapse.NewRefsLookup(g)
	t.Logf("input = %s", refs["main"])
	childPattern := refs["main"].Concat.RightPattern.Concat.LeftPattern.Or.LeftPattern.TreeNode.Pattern
	_, names := getAllReachableExprs(refs, refs["main"])
	name := names[0]
	t.Logf("name = %s", name)
	du := deriv(refs, refs["main"], nil, name)
	child := interp.Simplify(refs, du.child)
	t.Logf("child = %s", child)
	if !child.Equal(childPattern) {
		t.Fatalf("expected %s\n, but got %s", childPattern, child)
	}
}

func DisabledTestConvert4(t *testing.T) {
	g := tests.AndNameTelephonePerson.Grammar()
	refs := relapse.NewRefsLookup(g)
	t.Logf("input = %s", refs["main"])
	childPattern := refs["main"].And.LeftPattern.Concat.RightPattern.Concat.LeftPattern.TreeNode.Pattern
	_, names := getAllReachableExprs(refs, refs["main"])
	name := names[0]
	t.Logf("name = %s", name)
	du := deriv(refs, refs["main"], nil, name)
	child := interp.Simplify(refs, du.child)
	t.Logf("child = %s", child)
	if !child.Equal(childPattern) {
		t.Fatalf("expected %s\n, but got %s", childPattern, child)
	}
}
