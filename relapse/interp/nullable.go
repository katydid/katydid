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

package interp

import (
	"fmt"
	"github.com/katydid/katydid/relapse/ast"
)

//TODO improve nullable for left recursion using fix points
// https://github.com/kennknowles/go-yid/blob/master/src/yid/nullable.go
//This is a naive implementation and it does not handle left recursion
func Nullable(refs ast.RefLookup, p *ast.Pattern) bool {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *ast.Empty:
		return true
	case *ast.TreeNode:
		return false
	case *ast.LeafNode:
		return false
	case *ast.Concat:
		return Nullable(refs, v.GetLeftPattern()) && Nullable(refs, v.GetRightPattern())
	case *ast.Or:
		return Nullable(refs, v.GetLeftPattern()) || Nullable(refs, v.GetRightPattern())
	case *ast.And:
		return Nullable(refs, v.GetLeftPattern()) && Nullable(refs, v.GetRightPattern())
	case *ast.ZeroOrMore:
		return true
	case *ast.Reference:
		return Nullable(refs, refs[v.GetName()])
	case *ast.Not:
		return !(Nullable(refs, v.GetPattern()))
	case *ast.ZAny:
		return true
	case *ast.Contains:
		return Nullable(refs, v.GetPattern())
	case *ast.Optional:
		return true
	case *ast.Interleave:
		return Nullable(refs, v.GetLeftPattern()) && Nullable(refs, v.GetRightPattern())
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}
