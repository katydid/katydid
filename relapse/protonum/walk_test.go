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

package protonum

import (
	"fmt"
	"github.com/katydid/katydid/relapse"
)

func topDown(pattern *relapse.Pattern, f func(*relapse.Pattern)) {
	f(pattern)
	typ := pattern.GetValue()
	switch v := typ.(type) {
	case *relapse.TreeNode:
		topDown(v.GetPattern(), f)
	case *relapse.Concat:
		topDown(v.GetLeftPattern(), f)
		topDown(v.GetRightPattern(), f)
	case *relapse.Or:
		topDown(v.GetLeftPattern(), f)
		topDown(v.GetRightPattern(), f)
	case *relapse.And:
		topDown(v.GetLeftPattern(), f)
		topDown(v.GetRightPattern(), f)
	case *relapse.ZeroOrMore:
		topDown(v.GetPattern(), f)
	case *relapse.Not:
		topDown(v.GetPattern(), f)
	case *relapse.Contains:
		topDown(v.GetPattern(), f)
	case *relapse.Optional:
		topDown(v.GetPattern(), f)
	case *relapse.Interleave:
		topDown(v.GetLeftPattern(), f)
		topDown(v.GetRightPattern(), f)
	case *relapse.Empty, *relapse.LeafNode, *relapse.Reference, *relapse.ZAny:
		// do nothing
	default:
		panic(fmt.Sprintf("unknown pattern typ %T", typ))
	}
}

func topDownName(n *relapse.NameExpr, f func(*relapse.NameExpr)) {
	f(n)
	typ := n.GetValue()
	switch v := typ.(type) {
	case *relapse.Name, *relapse.AnyName:
		//do nothing
	case *relapse.AnyNameExcept:
		topDownName(v.GetExcept(), f)
	case *relapse.NameChoice:
		topDownName(v.GetLeft(), f)
		topDownName(v.GetRight(), f)
	default:
		panic(fmt.Sprintf("unknown name typ %T", typ))
	}
}

func allNames(p *relapse.Pattern, f func(*relapse.NameExpr) bool) bool {
	ret := true
	topDown(p, func(p *relapse.Pattern) {
		if p.TreeNode == nil {
			return
		}
		topDownName(p.TreeNode.GetName(), func(n *relapse.NameExpr) {
			ret = ret && f(n)
		})
	})
	return ret
}

func anyNames(p *relapse.Pattern, f func(*relapse.NameExpr) bool) bool {
	ret := false
	topDown(p, func(p *relapse.Pattern) {
		if p.TreeNode == nil {
			return
		}
		topDownName(p.TreeNode.GetName(), func(n *relapse.NameExpr) {
			ret = ret || f(n)
		})
	})
	return ret
}

func anyStringNames(p *relapse.Pattern) bool {
	return anyNames(p, func(n *relapse.NameExpr) bool {
		if n.GetName() == nil {
			return false
		}
		return n.GetName().StringValue != nil
	})
}

func onlyUintNames(p *relapse.Pattern) bool {
	return allNames(p, func(n *relapse.NameExpr) bool {
		if n.GetName() == nil {
			return true
		}
		return n.GetName().UintValue != nil
	})
}
