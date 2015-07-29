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
	. "github.com/katydid/katydid/relapse/interp"
	"github.com/katydid/katydid/tests"
	"testing"
)

func TestSimplify1(t *testing.T) {
	c := relapse.NewConcat(relapse.NewEmptySet(), relapse.NewNot(relapse.NewEmptySet()))
	refs := relapse.RefLookup{"main": c}
	s := Simplify(refs, c)
	if !s.Equal(relapse.NewEmptySet()) {
		t.Fatalf("Expected EmptySet, but got %s", s)
	}
}

func TestSimplify2(t *testing.T) {
	refs := relapse.NewRefsLookup(tests.AndNameTelephonePerson.Grammar())
	s := Simplify(refs, refs["main"])
	if s.Equal(relapse.NewEmptySet()) {
		t.Fatalf("Did not expected EmptySet")
	}
}
