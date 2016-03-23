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
	"github.com/katydid/katydid/relapse"
)

//TODO improve nullable for left recursion using fix points
// https://github.com/kennknowles/go-yid/blob/master/src/yid/nullable.go
//This is a naive implementation and it does not handle left recursion
func Nullable(refs relapse.RefLookup, p *relapse.Pattern) bool {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return true
	case *relapse.TreeNode:
		return false
	case *relapse.LeafNode:
		return false
	case *relapse.Concat:
		return Nullable(refs, v.GetLeftPattern()) && Nullable(refs, v.GetRightPattern())
	case *relapse.Or:
		return Nullable(refs, v.GetLeftPattern()) || Nullable(refs, v.GetRightPattern())
	case *relapse.And:
		return Nullable(refs, v.GetLeftPattern()) && Nullable(refs, v.GetRightPattern())
	case *relapse.ZeroOrMore:
		return true
	case *relapse.Reference:
		return Nullable(refs, refs[v.GetName()])
	case *relapse.Not:
		return !(Nullable(refs, v.GetPattern()))
	case *relapse.ZAny:
		return true
	case *relapse.Contains:
		return Nullable(refs, v.GetPattern())
	case *relapse.Optional:
		return true
	case *relapse.Interleave:
		return Nullable(refs, v.GetLeftPattern()) && Nullable(refs, v.GetRightPattern())
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}
