//  Copyright 2017 Walter Schulze
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
