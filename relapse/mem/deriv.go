//  Copyright 2016 Walter Schulze
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

package mem

import (
	"fmt"
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/funcs"
	"github.com/katydid/katydid/relapse/interp"
	nameexpr "github.com/katydid/katydid/relapse/name"
	"io"
)

func deriv(mem *Mem, current int, tree parser.Interface) (int, error) {
	for {
		if !mem.escapable(current) {
			return current, nil
		}
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				return 0, err
			}
		}
		callTree, err := mem.getCallTree(current)
		if err != nil {
			return 0, err
		}
		childState, stackElm, err := callTree.Eval(tree)
		if err != nil {
			return 0, err
		}
		if !tree.IsLeaf() {
			tree.Down()
			childState, err = deriv(mem, childState, tree)
			if err != nil {
				return 0, err
			}
			tree.Up()
		}
		current = mem.getReturn(stackElm, childState)
	}
	return current, nil
}

func derivCalls(refs map[string]*ast.Pattern, getFunc func(*ast.Expr) funcs.Bool, patterns []*ast.Pattern) []*callable {
	res := []*callable{}
	for _, pattern := range patterns {
		cs := derivCall(refs, getFunc, pattern)
		res = append(res, cs...)
	}
	return res
}

func derivCall(refs map[string]*ast.Pattern, getFunc func(*ast.Expr) funcs.Bool, p *ast.Pattern) []*callable {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *ast.Empty:
		return []*callable{}
	case *ast.ZAny:
		return []*callable{}
	case *ast.TreeNode:
		b := nameexpr.NameToFunc(v.GetName())
		return []*callable{&callable{b, v.GetPattern(), ast.NewNot(ast.NewZAny())}}
	case *ast.LeafNode:
		b := getFunc(v.GetExpr())
		return []*callable{&callable{b, ast.NewEmpty(), ast.NewNot(ast.NewZAny())}}
	case *ast.Concat:
		l := derivCall(refs, getFunc, v.GetLeftPattern())
		if !interp.Nullable(refs, v.GetLeftPattern()) {
			return l
		}
		r := derivCall(refs, getFunc, v.GetRightPattern())
		return append(l, r...)
	case *ast.Or:
		return derivCall2(refs, getFunc, v.GetLeftPattern(), v.GetRightPattern())
	case *ast.And:
		return derivCall2(refs, getFunc, v.GetLeftPattern(), v.GetRightPattern())
	case *ast.Interleave:
		return derivCall2(refs, getFunc, v.GetLeftPattern(), v.GetRightPattern())
	case *ast.ZeroOrMore:
		return derivCall(refs, getFunc, v.GetPattern())
	case *ast.Reference:
		return derivCall(refs, getFunc, refs[v.GetName()])
	case *ast.Not:
		return derivCall(refs, getFunc, v.GetPattern())
	case *ast.Contains:
		return derivCall(refs, getFunc, ast.NewConcat(ast.NewZAny(), ast.NewConcat(v.GetPattern(), ast.NewZAny())))
	case *ast.Optional:
		return derivCall(refs, getFunc, ast.NewOr(v.GetPattern(), ast.NewEmpty()))
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func derivCall2(refs map[string]*ast.Pattern, getFunc func(*ast.Expr) funcs.Bool, left, right *ast.Pattern) []*callable {
	l := derivCall(refs, getFunc, left)
	r := derivCall(refs, getFunc, right)
	return append(l, r...)
}

func derivReturns(refs map[string]*ast.Pattern, originals []*ast.Pattern, nullable []bool) []*ast.Pattern {
	res := make([]*ast.Pattern, len(originals))
	rest := nullable
	for i, original := range originals {
		res[i], rest = derivReturn(refs, original, rest)
	}
	return res
}

func derivReturn(refs map[string]*ast.Pattern, p *ast.Pattern, nullable []bool) (*ast.Pattern, []bool) {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *ast.Empty:
		return ast.NewNot(ast.NewZAny()), nullable
	case *ast.ZAny:
		return ast.NewZAny(), nullable
	case *ast.TreeNode:
		if nullable[0] {
			return ast.NewEmpty(), nullable[1:]
		}
		return ast.NewNot(ast.NewZAny()), nullable[1:]
	case *ast.LeafNode:
		if nullable[0] {
			return ast.NewEmpty(), nullable[1:]
		}
		return ast.NewNot(ast.NewZAny()), nullable[1:]
	case *ast.Concat:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), nullable)
		leftConcat := ast.NewConcat(l, v.GetRightPattern())
		if !interp.Nullable(refs, v.GetLeftPattern()) {
			return leftConcat, leftRest
		}
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return ast.NewOr(leftConcat, r), rightRest
	case *ast.Or:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), nullable)
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return ast.NewOr(l, r), rightRest
	case *ast.And:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), nullable)
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return ast.NewAnd(l, r), rightRest
	case *ast.Interleave:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), nullable)
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return ast.NewOr(ast.NewInterleave(l, v.GetRightPattern()), ast.NewInterleave(r, v.GetLeftPattern())), rightRest
	case *ast.ZeroOrMore:
		c, rest := derivReturn(refs, v.GetPattern(), nullable)
		return ast.NewConcat(c, p), rest
	case *ast.Reference:
		return derivReturn(refs, refs[v.GetName()], nullable)
	case *ast.Not:
		c, rest := derivReturn(refs, v.GetPattern(), nullable)
		return ast.NewNot(c), rest
	case *ast.Contains:
		return derivReturn(refs, ast.NewConcat(ast.NewZAny(), ast.NewConcat(v.GetPattern(), ast.NewZAny())), nullable)
	case *ast.Optional:
		return derivReturn(refs, ast.NewOr(v.GetPattern(), ast.NewEmpty()), nullable)
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}
