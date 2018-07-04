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

package intern

import (
	"testing"

	"github.com/katydid/katydid/relapse/ast"
)

func TestMergeAndRecord(t *testing.T) {
	inAny := func(p *ast.Pattern) *ast.Pattern {
		return ast.NewContains(ast.NewTreeNode(ast.NewAnyName(), p))
	}
	in := func(name string, p *ast.Pattern) *ast.Pattern {
		return ast.NewContains(ast.NewTreeNode(ast.NewStringName(name), p))
	}
	not := ast.NewNot

	eqA := ast.NewLeafNode(ast.NewEqual(ast.NewStringConst("aaaaaaaa@mm~")))
	eqB := ast.NewLeafNode(ast.NewEqual(ast.NewStringConst("bbbbbbb@~")))

	// .FountainTarget:!(._:== "aaaaaaaa@mm~")
	fa := in("FountainTarget", not(inAny(eqA)))
	// .FountainTarget:!(._:== "bbbbbbb@~")
	fb := in("FountainTarget", not(inAny(eqB)))
	// .BridgePepper:!(._:== "bbbbbbb@~")
	bb := in("BridgePepper", not(inAny(eqB)))

	c := NewConstructorOptimizedForRecords().(*construct)

	pattern_fb, err := c.NewPattern(fb)
	if err != nil {
		t.Fatal(err)
	}
	pattern_fa, err := c.NewPattern(fa)
	if err != nil {
		t.Fatal(err)
	}
	pattern_bb, err := c.NewPattern(bb)
	if err != nil {
		t.Fatal(err)
	}
	ps, err := c.mergeContainsAnd([]*Pattern{pattern_fa, pattern_fb, pattern_bb})
	if err != nil {
		t.Fatal(err)
	}
	if !ps[1].Equal(pattern_bb) {
		t.Fatalf("incorrectly merged")
	}
}
