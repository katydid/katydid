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
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/combinator"
	. "github.com/katydid/katydid/relapse/interp"
	"github.com/katydid/katydid/tests"
	"testing"
)

func TestSimplify1(t *testing.T) {
	c := relapse.NewConcat(relapse.NewNot(relapse.NewZAny()), relapse.NewZAny())
	refs := relapse.RefLookup{"main": c}
	s := Simplify(refs, c)
	if !s.Equal(relapse.NewNot(relapse.NewZAny())) {
		t.Fatalf("Expected EmptySet, but got %s", s)
	}
}

func TestSimplify2(t *testing.T) {
	refs := relapse.NewRefsLookup(tests.AndNameTelephonePerson.Grammar())
	s := Simplify(refs, refs["main"])
	if s.Equal(relapse.NewNot(relapse.NewZAny())) {
		t.Fatalf("Did not expected EmptySet")
	}
}

func newT(s string) *relapse.Pattern {
	return relapse.NewTreeNode(relapse.NewStringName(s), relapse.NewEmpty())
}

func TestSimplifyOr1(t *testing.T) {
	input := relapse.NewOr(newT("B"), relapse.NewOr(newT("C"), relapse.NewOr(newT("A"), newT("B"))))
	refs := relapse.RefLookup{"main": input}
	output := Simplify(refs, input)
	expected := relapse.NewOr(relapse.NewOr(newT("A"), newT("B")), newT("C"))
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}

func TestSimplifyOr2(t *testing.T) {
	input := relapse.NewOr(relapse.NewOr(newT("A"), newT("B")), relapse.NewOr(newT("B"), newT("C")))
	refs := relapse.RefLookup{"main": input}
	output := Simplify(refs, input)
	expected := relapse.NewOr(relapse.NewOr(newT("A"), newT("B")), newT("C"))
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}

func TestSimplifyTree(t *testing.T) {
	left := relapse.NewTreeNode(relapse.NewStringName("A"),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewContains(
			relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewZAny()),
		)),
	)
	right := relapse.NewTreeNode(relapse.NewStringName("A"),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewContains(
			relapse.NewTreeNode(relapse.NewStringName("D"), relapse.NewZAny()),
		)),
	)
	input := relapse.NewAnd(left, right)
	expected := relapse.NewTreeNode(relapse.NewStringName("A"),
		relapse.NewTreeNode(relapse.NewStringName("B"), relapse.NewAnd(
			relapse.NewContains(
				relapse.NewTreeNode(relapse.NewStringName("C"), relapse.NewZAny()),
			),
			relapse.NewContains(
				relapse.NewTreeNode(relapse.NewStringName("D"), relapse.NewZAny()),
			),
		)),
	)
	refs := relapse.RefLookup{"main": input}
	output := Simplify(refs, input)
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}

func TestSimplifyFalseLeaf(t *testing.T) {
	input := combinator.Value(funcs.And(funcs.StringEq(funcs.StringVar(), funcs.StringConst("a")), funcs.StringEq(funcs.StringVar(), funcs.StringConst("b"))))
	expected := relapse.NewNot(relapse.NewZAny())
	refs := relapse.RefLookup{"main": input}
	output := Simplify(refs, input)
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}

func TestSimplifyFalseTreeNode(t *testing.T) {
	input := relapse.NewTreeNode(relapse.NewAnyNameExcept(relapse.NewAnyName()), relapse.NewZAny())
	expected := relapse.NewNot(relapse.NewZAny())
	refs := relapse.RefLookup{"main": input}
	output := Simplify(refs, input)
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}

func TestSimplifyTreeNodeWithNotZanyChild(t *testing.T) {
	input := relapse.NewTreeNode(relapse.NewAnyName(), relapse.NewNot(relapse.NewZAny()))
	expected := relapse.NewNot(relapse.NewZAny())
	refs := relapse.RefLookup{"main": input}
	output := Simplify(refs, input)
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}

func TestSimplifyContainsFalseTreeNode(t *testing.T) {
	input := relapse.NewContains(relapse.NewTreeNode(relapse.NewAnyNameExcept(relapse.NewAnyName()), relapse.NewZAny()))
	expected := relapse.NewNot(relapse.NewZAny())
	refs := relapse.RefLookup{"main": input}
	output := Simplify(refs, input)
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}
