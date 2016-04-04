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
	"github.com/katydid/katydid/relapse/ast"
)

func topDown(pattern *ast.Pattern, f func(*ast.Pattern)) {
	f(pattern)
	typ := pattern.GetValue()
	switch v := typ.(type) {
	case *ast.TreeNode:
		topDown(v.GetPattern(), f)
	case *ast.Concat:
		topDown(v.GetLeftPattern(), f)
		topDown(v.GetRightPattern(), f)
	case *ast.Or:
		topDown(v.GetLeftPattern(), f)
		topDown(v.GetRightPattern(), f)
	case *ast.And:
		topDown(v.GetLeftPattern(), f)
		topDown(v.GetRightPattern(), f)
	case *ast.ZeroOrMore:
		topDown(v.GetPattern(), f)
	case *ast.Not:
		topDown(v.GetPattern(), f)
	case *ast.Contains:
		topDown(v.GetPattern(), f)
	case *ast.Optional:
		topDown(v.GetPattern(), f)
	case *ast.Interleave:
		topDown(v.GetLeftPattern(), f)
		topDown(v.GetRightPattern(), f)
	case *ast.Empty, *ast.LeafNode, *ast.Reference, *ast.ZAny:
		// do nothing
	default:
		panic(fmt.Sprintf("unknown pattern typ %T", typ))
	}
}

func topDownName(n *ast.NameExpr, f func(*ast.NameExpr)) {
	f(n)
	typ := n.GetValue()
	switch v := typ.(type) {
	case *ast.Name, *ast.AnyName:
		//do nothing
	case *ast.AnyNameExcept:
		topDownName(v.GetExcept(), f)
	case *ast.NameChoice:
		topDownName(v.GetLeft(), f)
		topDownName(v.GetRight(), f)
	default:
		panic(fmt.Sprintf("unknown name typ %T", typ))
	}
}

func allNames(p *ast.Pattern, f func(*ast.NameExpr) bool) bool {
	ret := true
	topDown(p, func(p *ast.Pattern) {
		if p.TreeNode == nil {
			return
		}
		topDownName(p.TreeNode.GetName(), func(n *ast.NameExpr) {
			ret = ret && f(n)
		})
	})
	return ret
}

func anyNames(p *ast.Pattern, f func(*ast.NameExpr) bool) bool {
	ret := false
	topDown(p, func(p *ast.Pattern) {
		if p.TreeNode == nil {
			return
		}
		topDownName(p.TreeNode.GetName(), func(n *ast.NameExpr) {
			ret = ret || f(n)
		})
	})
	return ret
}

func anyStringNames(p *ast.Pattern) bool {
	return anyNames(p, func(n *ast.NameExpr) bool {
		if n.GetName() == nil {
			return false
		}
		return n.GetName().StringValue != nil
	})
}

func onlyUintNames(p *ast.Pattern) bool {
	return allNames(p, func(n *ast.NameExpr) bool {
		if n.GetName() == nil {
			return true
		}
		return n.GetName().UintValue != nil
	})
}
