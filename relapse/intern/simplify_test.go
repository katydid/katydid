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

package intern_test

import (
	"testing"

	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/combinator"
	. "github.com/katydid/katydid/relapse/intern"
)

func TestSimplify1(t *testing.T) {
	c := NewConstructor()
	n, err := c.NewNot(c.NewZAny())
	if err != nil {
		t.Fatal(err)
	}
	p, err := c.NewConcat([]*Pattern{n, c.NewZAny()})
	if err != nil {
		t.Fatal(err)
	}
	if !p.Equal(c.NewNotZAny()) {
		t.Fatalf("Expected EmptySet, but got %v", p)
	}
}

var andNameTelephonePerson = combinator.G{
	"main": combinator.InOrder(
		combinator.AllOf(
			combinator.InOrder(
				combinator.Any(),
				combinator.In("Name", combinator.Value(
					combinator.Eq(combinator.StringVar(), combinator.StringConst("David"))),
				),
				combinator.Any(),
			),
			combinator.InOrder(
				combinator.Any(),
				combinator.In("Telephone", combinator.Value(
					combinator.Eq(combinator.StringVar(), combinator.StringConst("0123456789"))),
				),
				combinator.Any(),
			),
		),
	),
}

func TestSimplify2(t *testing.T) {
	c := NewConstructor()
	p, err := c.AddGrammar(andNameTelephonePerson.Grammar())
	if err != nil {
		t.Fatal(err)
	}
	if p.Equal(c.NewNotZAny()) {
		t.Fatalf("Did not expected EmptySet")
	}
}

func newT(s string) *ast.Pattern {
	return ast.NewTreeNode(ast.NewStringName(s), ast.NewEmpty())
}

func TestSimplifyOr1(t *testing.T) {
	c := NewConstructor()
	input := ast.NewOr(newT("B"), ast.NewOr(newT("C"), ast.NewOr(newT("A"), newT("B"))))
	got, err := c.NewPattern(input)
	if err != nil {
		t.Fatal(err)
	}
	expected := ast.NewOr(ast.NewOr(newT("A"), newT("B")), newT("C"))
	want, err := c.NewPattern(expected)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", got)
	if !want.Equal(got) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestSimplifyOr2(t *testing.T) {
	c := NewConstructor()
	input := ast.NewOr(ast.NewOr(newT("A"), newT("B")), ast.NewOr(newT("B"), newT("C")))
	got, err := c.NewPattern(input)
	if err != nil {
		t.Fatal(err)
	}
	expected := ast.NewOr(ast.NewOr(newT("A"), newT("B")), newT("C"))
	want, err := c.NewPattern(expected)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", got)
	if !want.Equal(got) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestSimplifyTree(t *testing.T) {
	c := NewConstructor()

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
	got, err := c.NewPattern(input)
	if err != nil {
		t.Fatal(err)
	}

	expected := ast.NewTreeNode(ast.NewStringName("A"),
		ast.NewTreeNode(ast.NewStringName("B"), ast.NewAnd(
			ast.NewContains(
				ast.NewTreeNode(ast.NewStringName("D"), ast.NewZAny()),
			),
			ast.NewContains(
				ast.NewTreeNode(ast.NewStringName("C"), ast.NewZAny()),
			),
		)),
	)
	want, err := c.NewPattern(expected)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", got)
	if !want.Equal(got) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestSimplifyFalseLeaf(t *testing.T) {
	c := NewConstructor()

	input := combinator.Value(combinator.And(combinator.Eq(combinator.StringVar(), combinator.StringConst("a")), combinator.Eq(combinator.StringVar(), combinator.StringConst("b"))))
	got, err := c.NewPattern(input)
	if err != nil {
		t.Fatal(err)
	}

	expected := ast.NewNot(ast.NewZAny())
	want, err := c.NewPattern(expected)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", got)
	if !want.Equal(got) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestSimplifyFalseTreeNode(t *testing.T) {
	c := NewConstructor()

	input := ast.NewTreeNode(ast.NewAnyNameExcept(ast.NewAnyName()), ast.NewZAny())
	got, err := c.NewPattern(input)
	if err != nil {
		t.Fatal(err)
	}

	expected := ast.NewNot(ast.NewZAny())
	want, err := c.NewPattern(expected)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", got)
	if !want.Equal(got) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestSimplifyTreeNodeWithNotZanyChild(t *testing.T) {
	c := NewConstructor()

	input := ast.NewTreeNode(ast.NewAnyName(), ast.NewNot(ast.NewZAny()))
	got, err := c.NewPattern(input)
	if err != nil {
		t.Fatal(err)
	}

	expected := ast.NewNot(ast.NewZAny())
	want, err := c.NewPattern(expected)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", got)
	if !want.Equal(got) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestSimplifyContainsFalseTreeNode(t *testing.T) {
	c := NewConstructor()

	input := ast.NewContains(ast.NewTreeNode(ast.NewAnyNameExcept(ast.NewAnyName()), ast.NewZAny()))
	got, err := c.NewPattern(input)
	if err != nil {
		t.Fatal(err)
	}

	expected := ast.NewNot(ast.NewZAny())
	want, err := c.NewPattern(expected)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", got)
	if !want.Equal(got) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestSimplifyRecordLeaf1(t *testing.T) {
	c := NewConstructorOptimizedForRecords()

	input := ast.NewAnd(
		ast.NewContains(ast.NewTreeNode(ast.NewStringName("A"), combinator.Value(combinator.Contains(combinator.StringVar(), combinator.StringConst("a"))))),
		ast.NewContains(ast.NewTreeNode(ast.NewStringName("A"), combinator.Value(combinator.Contains(combinator.StringVar(), combinator.StringConst("b"))))),
	)
	t.Logf("input: %v", input)
	got, err := c.NewPattern(input)
	if err != nil {
		t.Fatal(err)
	}

	expected := ast.NewContains(ast.NewTreeNode(ast.NewStringName("A"), combinator.Value(combinator.And(
		combinator.Contains(combinator.StringVar(), combinator.StringConst("a")),
		combinator.Contains(combinator.StringVar(), combinator.StringConst("b")),
	))))
	want, err := c.NewPattern(expected)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", got)
	if !want.Equal(got) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}

func TestSimplifyRecordLeaf2(t *testing.T) {
	c := NewConstructorOptimizedForRecords()

	input := ast.NewAnd(
		ast.NewContains(ast.NewTreeNode(ast.NewStringName("A"), ast.NewContains(ast.NewTreeNode(ast.NewStringName("B"), combinator.Value(combinator.Contains(combinator.StringVar(), combinator.StringConst("a"))))))),
		ast.NewContains(ast.NewTreeNode(ast.NewStringName("A"), ast.NewContains(ast.NewTreeNode(ast.NewStringName("B"), combinator.Value(combinator.Contains(combinator.StringVar(), combinator.StringConst("b"))))))),
	)
	t.Logf("input: %v", input)
	got, err := c.NewPattern(input)
	if err != nil {
		t.Fatal(err)
	}

	expected := ast.NewContains(ast.NewTreeNode(ast.NewStringName("A"), ast.NewContains(ast.NewTreeNode(ast.NewStringName("B"), combinator.Value(combinator.And(
		combinator.Contains(combinator.StringVar(), combinator.StringConst("a")),
		combinator.Contains(combinator.StringVar(), combinator.StringConst("b")),
	))))))
	want, err := c.NewPattern(expected)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", got)
	if !want.Equal(got) {
		t.Fatalf("want %v, but got %v", want, got)
	}
}
