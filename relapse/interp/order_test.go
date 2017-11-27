package interp

import (
	"reflect"
	"testing"

	"github.com/katydid/katydid/relapse/ast"
)

func TestOrderedSet(t *testing.T) {
	email := ast.NewZeroOrMore(ast.NewTreeNode(ast.NewStringName("Email"), ast.NewZAny()))
	addr := ast.NewOptional(ast.NewTreeNode(ast.NewStringName("Addr"), ast.NewZAny()))
	o1 := orderedSet([]*ast.Pattern{email, addr})
	o2 := orderedSet([]*ast.Pattern{addr, email})
	if !reflect.DeepEqual(o1, o2) {
		t.Fatalf("%s != %s", o1, o2)
	}
}

func TestSimplifyOrderedSet1(t *testing.T) {
	email := ast.NewZeroOrMore(ast.NewTreeNode(ast.NewStringName("Email"), ast.NewZAny()))
	addr := ast.NewOptional(ast.NewTreeNode(ast.NewStringName("Addr"), ast.NewZAny()))
	ea := ast.NewInterleave(email, addr)
	ae := ast.NewInterleave(addr, email)

	ors := ast.NewOr(ea, ae, ea)
	got := NewSimplifier(ors.Grammar()).Simplify(ors)

	want := ae

	if !want.Equal(got) {
		t.Fatalf("want %s got %s", want, got)
	}
}
