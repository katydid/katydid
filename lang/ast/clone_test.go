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

package lang

import (
	"github.com/katydid/katydid/expr/ast"
	"reflect"
	"testing"
)

func TestClone(t *testing.T) {
	in := NewGrammar(map[string]*Pattern{
		"main": NewAnd(NewEmpty(), NewTreeNode(NewName("a"), NewLeafNode(&expr.Expr{}))),
	})
	out := in.Clone()
	if !reflect.DeepEqual(in, out) {
		t.Fatalf("not cloned")
	}
	treeNode := in.GetPatternDecls()[0].GetPattern().GetAnd().GetRightPattern().GetTreeNode()
	treeNode.Pattern = NewNot(NewEmptySet())
	if reflect.DeepEqual(in, out) {
		t.Fatalf("not cloned")
	}
}
