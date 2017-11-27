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

package interp

import (
	"fmt"
	"io"

	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/compose"
	nameexpr "github.com/katydid/katydid/relapse/name"
	"github.com/katydid/katydid/relapse/sets"
)

//Interpret interprets the grammar given the parser and returns whether the parser is valid given the grammar.
//Interpret uses derivatives and simplification to recusively derive the resulting grammar.
//This resulting grammar's nullability then represents the result of the function.
//This implementation does not handle immediate recursion, see the HasRecursion function.
func Interpret(g *ast.Grammar, parser parser.Interface) (bool, error) {
	refs := ast.NewRefLookup(g)
	finals, err := deriv(refs, []*ast.Pattern{refs["main"]}, parser)
	if err != nil {
		return false, err
	}
	return Nullable(refs, finals[0]), nil
}

func escapable(patterns []*ast.Pattern) bool {
	for _, pattern := range patterns {
		if pattern.ZAny != nil {
			continue
		}
		if pattern.Not != nil {
			if pattern.GetNot().GetPattern().ZAny != nil {
				continue
			}
		}
		return true
	}
	return false
}

func deriv(refs ast.RefLookup, patterns []*ast.Pattern, tree parser.Interface) ([]*ast.Pattern, error) {
	var resPatterns []*ast.Pattern = patterns
	for {
		if !escapable(resPatterns) {
			return resPatterns, nil
		}
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		childPatterns, err := derivCalls(refs, resPatterns, tree)
		if err != nil {
			return nil, err
		}
		if tree.IsLeaf() {
			//do nothing
		} else {
			tree.Down()
			zchild, zi := sets.Zip(childPatterns)
			zchild, err = deriv(refs, zchild, tree)
			if err != nil {
				return nil, err
			}
			childPatterns = sets.Unzip(zchild, zi)
			tree.Up()
		}
		resPatterns = derivReturns(refs, resPatterns, childPatterns)
		resPatterns = simps(refs, resPatterns)
	}
	return resPatterns, nil
}

func simps(refs ast.RefLookup, patterns []*ast.Pattern) []*ast.Pattern {
	for i := range patterns {
		patterns[i] = NewSimplifier(ast.NewGrammar(refs)).Simplify(patterns[i])
	}
	return patterns
}

func derivCalls(refs ast.RefLookup, patterns []*ast.Pattern, label parser.Value) ([]*ast.Pattern, error) {
	res := []*ast.Pattern{}
	for _, pattern := range patterns {
		ps, err := derivCall(refs, pattern, label)
		if err != nil {
			return nil, err
		}
		ps = simps(refs, ps)
		res = append(res, ps...)
	}
	return res, nil
}

func derivCall(refs ast.RefLookup, p *ast.Pattern, label parser.Value) ([]*ast.Pattern, error) {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *ast.Empty:
		return []*ast.Pattern{}, nil
	case *ast.ZAny:
		return []*ast.Pattern{}, nil
	case *ast.TreeNode:
		b := nameexpr.NameToFunc(v.GetName())
		f, err := compose.NewBoolFunc(b)
		if err != nil {
			return nil, err
		}
		eval, err := f.Eval(label)
		if err != nil {
			return nil, err
		}
		if eval {
			return []*ast.Pattern{v.GetPattern()}, nil
		}
		return []*ast.Pattern{ast.NewNot(ast.NewZAny())}, nil
	case *ast.LeafNode:
		b, err := compose.NewBool(v.GetExpr())
		if err != nil {
			return nil, err
		}
		f, err := compose.NewBoolFunc(b)
		if err != nil {
			return nil, err
		}
		eval, err := f.Eval(label)
		if err != nil {
			return nil, err
		}
		if eval {
			return []*ast.Pattern{ast.NewEmpty()}, nil
		}
		return []*ast.Pattern{ast.NewNot(ast.NewZAny())}, nil
	case *ast.Concat:
		l, err := derivCall(refs, v.GetLeftPattern(), label)
		if err != nil {
			return nil, err
		}
		if !Nullable(refs, v.GetLeftPattern()) {
			return l, nil
		}
		r, err := derivCall(refs, v.GetRightPattern(), label)
		if err != nil {
			return nil, err
		}
		return append(l, r...), nil
	case *ast.Or:
		return derivCall2(refs, v.GetLeftPattern(), v.GetRightPattern(), label)
	case *ast.And:
		return derivCall2(refs, v.GetLeftPattern(), v.GetRightPattern(), label)
	case *ast.Interleave:
		return derivCall2(refs, v.GetLeftPattern(), v.GetRightPattern(), label)
	case *ast.ZeroOrMore:
		return derivCall(refs, v.GetPattern(), label)
	case *ast.Reference:
		return derivCall(refs, refs[v.GetName()], label)
	case *ast.Not:
		return derivCall(refs, v.GetPattern(), label)
	case *ast.Contains:
		return derivCall(refs, ast.NewConcat(ast.NewZAny(), ast.NewConcat(v.GetPattern(), ast.NewZAny())), label)
	case *ast.Optional:
		return derivCall(refs, ast.NewOr(v.GetPattern(), ast.NewEmpty()), label)
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func derivCall2(refs ast.RefLookup, left, right *ast.Pattern, label parser.Value) ([]*ast.Pattern, error) {
	l, err := derivCall(refs, left, label)
	if err != nil {
		return nil, err
	}
	r, err := derivCall(refs, right, label)
	if err != nil {
		return nil, err
	}
	return append(l, r...), nil
}

func derivReturns(refs ast.RefLookup, originals []*ast.Pattern, evaluated []*ast.Pattern) []*ast.Pattern {
	res := make([]*ast.Pattern, len(originals))
	rest := evaluated
	for i, original := range originals {
		res[i], rest = derivReturn(refs, original, rest)
	}
	return res
}

func derivReturn(refs ast.RefLookup, p *ast.Pattern, patterns []*ast.Pattern) (*ast.Pattern, []*ast.Pattern) {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *ast.Empty:
		return ast.NewNot(ast.NewZAny()), patterns
	case *ast.ZAny:
		return ast.NewZAny(), patterns
	case *ast.TreeNode:
		if Nullable(refs, patterns[0]) {
			return ast.NewEmpty(), patterns[1:]
		}
		return ast.NewNot(ast.NewZAny()), patterns[1:]
	case *ast.LeafNode:
		if Nullable(refs, patterns[0]) {
			return ast.NewEmpty(), patterns[1:]
		}
		return ast.NewNot(ast.NewZAny()), patterns[1:]
	case *ast.Concat:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), patterns)
		leftConcat := ast.NewConcat(l, v.GetRightPattern())
		if !Nullable(refs, v.GetLeftPattern()) {
			return leftConcat, leftRest
		}
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return ast.NewOr(leftConcat, r), rightRest
	case *ast.Or:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), patterns)
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return ast.NewOr(l, r), rightRest
	case *ast.And:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), patterns)
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return ast.NewAnd(l, r), rightRest
	case *ast.Interleave:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), patterns)
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return ast.NewOr(ast.NewInterleave(l, v.GetRightPattern()), ast.NewInterleave(r, v.GetLeftPattern())), rightRest
	case *ast.ZeroOrMore:
		c, rest := derivReturn(refs, v.GetPattern(), patterns)
		return ast.NewConcat(c, p), rest
	case *ast.Reference:
		return derivReturn(refs, refs[v.GetName()], patterns)
	case *ast.Not:
		c, rest := derivReturn(refs, v.GetPattern(), patterns)
		return ast.NewNot(c), rest
	case *ast.Contains:
		return derivReturn(refs, ast.NewConcat(ast.NewZAny(), ast.NewConcat(v.GetPattern(), ast.NewZAny())), patterns)
	case *ast.Optional:
		return derivReturn(refs, ast.NewOr(v.GetPattern(), ast.NewEmpty()), patterns)
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}
