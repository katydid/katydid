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
	"github.com/katydid/katydid/relapse/combinator"
	"github.com/katydid/katydid/relapse/funcs"
	. "github.com/katydid/katydid/relapse/interp"
	"github.com/katydid/katydid/relapse/tests"
	"testing"
)

func TestSimplify1(t *testing.T) {
	c := ast.NewConcat(ast.NewNot(ast.NewZAny()), ast.NewZAny())
	s := NewSimplifier(c.Grammar()).Simplify(c)
	if !s.Equal(ast.NewNot(ast.NewZAny())) {
		t.Fatalf("Expected EmptySet, but got %s", s)
	}
}

func TestSimplify2(t *testing.T) {
	s := NewSimplifier(tests.AndNameTelephonePerson.Grammar()).Simplify(tests.AndNameTelephonePerson["main"])
	if s.Equal(ast.NewNot(ast.NewZAny())) {
		t.Fatalf("Did not expected EmptySet")
	}
}

func newT(s string) *ast.Pattern {
	return ast.NewTreeNode(ast.NewStringName(s), ast.NewEmpty())
}

func TestSimplifyOr1(t *testing.T) {
	input := ast.NewOr(newT("B"), ast.NewOr(newT("C"), ast.NewOr(newT("A"), newT("B"))))
	output := NewSimplifier(input.Grammar()).Simplify(input)
	expected := ast.NewOr(ast.NewOr(newT("A"), newT("B")), newT("C"))
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}

func TestSimplifyOr2(t *testing.T) {
	input := ast.NewOr(ast.NewOr(newT("A"), newT("B")), ast.NewOr(newT("B"), newT("C")))
	output := NewSimplifier(input.Grammar()).Simplify(input)
	expected := ast.NewOr(ast.NewOr(newT("A"), newT("B")), newT("C"))
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}

func TestSimplifyTree(t *testing.T) {
	left := ast.NewTreeNode(ast.NewStringName("A"),
		ast.NewTreeNode(ast.NewStringName("B"), ast.NewContains(
			ast.NewTreeNode(ast.NewStringName("C"), ast.NewZAny()),
		)),
	)
	right := ast.NewTreeNode(ast.NewStringName("A"),
		ast.NewTreeNode(ast.NewStringName("B"), ast.NewContains(
			ast.NewTreeNode(ast.NewStringName("D"), ast.NewZAny()),
		)),
	)
	input := ast.NewAnd(left, right)
	expected := ast.NewTreeNode(ast.NewStringName("A"),
		ast.NewTreeNode(ast.NewStringName("B"), ast.NewAnd(
			ast.NewContains(
				ast.NewTreeNode(ast.NewStringName("C"), ast.NewZAny()),
			),
			ast.NewContains(
				ast.NewTreeNode(ast.NewStringName("D"), ast.NewZAny()),
			),
		)),
	)
	output := NewSimplifier(input.Grammar()).Simplify(input)
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}

func TestSimplifyFalseLeaf(t *testing.T) {
	input := combinator.Value(funcs.And(funcs.StringEq(funcs.StringVar(), funcs.StringConst("a")), funcs.StringEq(funcs.StringVar(), funcs.StringConst("b"))))
	expected := ast.NewNot(ast.NewZAny())
	output := NewSimplifier(input.Grammar()).Simplify(input)
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}

func TestSimplifyFalseTreeNode(t *testing.T) {
	input := ast.NewTreeNode(ast.NewAnyNameExcept(ast.NewAnyName()), ast.NewZAny())
	expected := ast.NewNot(ast.NewZAny())
	output := NewSimplifier(input.Grammar()).Simplify(input)
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}

func TestSimplifyTreeNodeWithNotZanyChild(t *testing.T) {
	input := ast.NewTreeNode(ast.NewAnyName(), ast.NewNot(ast.NewZAny()))
	expected := ast.NewNot(ast.NewZAny())
	output := NewSimplifier(input.Grammar()).Simplify(input)
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}

func TestSimplifyContainsFalseTreeNode(t *testing.T) {
	input := ast.NewContains(ast.NewTreeNode(ast.NewAnyNameExcept(ast.NewAnyName()), ast.NewZAny()))
	expected := ast.NewNot(ast.NewZAny())
	output := NewSimplifier(input.Grammar()).Simplify(input)
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}

func TestSimplifyRecordLeaf1(t *testing.T) {
	input := ast.NewAnd(
		ast.NewContains(ast.NewTreeNode(ast.NewStringName("A"), combinator.Value(funcs.Contains(funcs.StringVar(), funcs.StringConst("a"))))),
		ast.NewContains(ast.NewTreeNode(ast.NewStringName("A"), combinator.Value(funcs.Contains(funcs.StringVar(), funcs.StringConst("b"))))),
	)
	t.Logf("input: %v", input)
	expected := ast.NewContains(ast.NewTreeNode(ast.NewStringName("A"), combinator.Value(funcs.And(
		funcs.Contains(funcs.StringVar(), funcs.StringConst("a")),
		funcs.Contains(funcs.StringVar(), funcs.StringConst("b")),
	))))
	output := NewSimplifier(input.Grammar()).OptimizeForRecord().Simplify(input)
	expected.Format()
	output.Format()
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}

func TestSimplifyRecordLeaf2(t *testing.T) {
	input := ast.NewAnd(
		ast.NewContains(ast.NewTreeNode(ast.NewStringName("A"), ast.NewContains(ast.NewTreeNode(ast.NewStringName("B"), combinator.Value(funcs.Contains(funcs.StringVar(), funcs.StringConst("a"))))))),
		ast.NewContains(ast.NewTreeNode(ast.NewStringName("A"), ast.NewContains(ast.NewTreeNode(ast.NewStringName("B"), combinator.Value(funcs.Contains(funcs.StringVar(), funcs.StringConst("b"))))))),
	)
	t.Logf("input: %v", input)
	expected := ast.NewContains(ast.NewTreeNode(ast.NewStringName("A"), ast.NewContains(ast.NewTreeNode(ast.NewStringName("B"), combinator.Value(funcs.And(
		funcs.Contains(funcs.StringVar(), funcs.StringConst("a")),
		funcs.Contains(funcs.StringVar(), funcs.StringConst("b")),
	))))))
	output := NewSimplifier(input.Grammar()).OptimizeForRecord().Simplify(input)
	expected.Format()
	output.Format()
	t.Logf("%v", output)
	if !expected.Equal(output) {
		t.Fatalf("expected %v, but got %v", expected, output)
	}
}
