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
	"github.com/katydid/katydid/expr/compose"
	nameexpr "github.com/katydid/katydid/expr/name"
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse"
	"io"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

//This is a naive implementation and it does not handle left recursion
func Interpret(g *relapse.Grammar, parser parser.Interface) bool {
	refs := relapse.NewRefsLookup(g)
	finals := deriv(refs, []*relapse.Pattern{refs["main"]}, parser)
	return Nullable(refs, finals[0])
}

func escapable(patterns []*relapse.Pattern) bool {
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

func deriv(refs map[string]*relapse.Pattern, patterns []*relapse.Pattern, tree parser.Interface) []*relapse.Pattern {
	var resPatterns []*relapse.Pattern = patterns
	for {
		if !escapable(resPatterns) {
			return resPatterns
		}
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		childPatterns := derivCalls(refs, resPatterns, tree)
		if tree.IsLeaf() {
			//do nothing
		} else {
			tree.Down()
			zchild, zi := zip(childPatterns)
			zchild = deriv(refs, zchild, tree)
			childPatterns = unzip(zchild, zi)
			tree.Up()
		}
		resPatterns = derivReturns(refs, resPatterns, childPatterns)
		resPatterns = simps(refs, resPatterns)
	}
	return resPatterns
}

func simps(refs map[string]*relapse.Pattern, patterns []*relapse.Pattern) []*relapse.Pattern {
	for i := range patterns {
		patterns[i] = Simplify(refs, patterns[i])
	}
	return patterns
}

func zip(patterns []*relapse.Pattern) ([]*relapse.Pattern, []int) {
	zipped := relapse.Set(patterns)
	relapse.Sort(zipped)
	indexes := make([]int, len(patterns))
	for i, pattern := range patterns {
		indexes[i] = relapse.Index(zipped, pattern)
	}
	return zipped, indexes
}

func unzip(patterns []*relapse.Pattern, indexes []int) []*relapse.Pattern {
	res := make([]*relapse.Pattern, len(indexes))
	for i, index := range indexes {
		res[i] = patterns[index]
	}
	return res
}

func derivCalls(refs map[string]*relapse.Pattern, patterns []*relapse.Pattern, label parser.Value) []*relapse.Pattern {
	res := []*relapse.Pattern{}
	for _, pattern := range patterns {
		ps := derivCall(refs, pattern, label)
		ps = simps(refs, ps)
		res = append(res, ps...)
	}
	return res
}

func derivCall(refs map[string]*relapse.Pattern, p *relapse.Pattern, label parser.Value) []*relapse.Pattern {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return []*relapse.Pattern{}
	case *relapse.ZAny:
		return []*relapse.Pattern{}
	case *relapse.TreeNode:
		b := nameexpr.NameToFunc(v.GetName())
		f, err := compose.NewBoolFunc(b)
		if err != nil {
			panic(err)
		}
		eval, err := f.Eval(label)
		if err != nil {
			panic(err)
		}
		if eval {
			return []*relapse.Pattern{v.GetPattern()}
		}
		return []*relapse.Pattern{relapse.NewNot(relapse.NewZAny())}
	case *relapse.LeafNode:
		b, err := compose.NewBool(v.GetExpr())
		if err != nil {
			panic(err)
		}
		f, err := compose.NewBoolFunc(b)
		if err != nil {
			panic(err)
		}
		eval, err := f.Eval(label)
		if err != nil {
			panic(err)
		}
		if eval {
			return []*relapse.Pattern{relapse.NewEmpty()}
		}
		return []*relapse.Pattern{relapse.NewNot(relapse.NewZAny())}
	case *relapse.Concat:
		l := derivCall(refs, v.GetLeftPattern(), label)
		if !Nullable(refs, v.GetLeftPattern()) {
			return l
		}
		r := derivCall(refs, v.GetRightPattern(), label)
		return append(l, r...)
	case *relapse.Or:
		return derivCall2(refs, v.GetLeftPattern(), v.GetRightPattern(), label)
	case *relapse.And:
		return derivCall2(refs, v.GetLeftPattern(), v.GetRightPattern(), label)
	case *relapse.Interleave:
		return derivCall2(refs, v.GetLeftPattern(), v.GetRightPattern(), label)
	case *relapse.ZeroOrMore:
		return derivCall(refs, v.GetPattern(), label)
	case *relapse.Reference:
		return derivCall(refs, refs[v.GetName()], label)
	case *relapse.Not:
		return derivCall(refs, v.GetPattern(), label)
	case *relapse.Contains:
		return derivCall(refs, relapse.NewConcat(relapse.NewZAny(), relapse.NewConcat(v.GetPattern(), relapse.NewZAny())), label)
	case *relapse.Optional:
		return derivCall(refs, relapse.NewOr(v.GetPattern(), relapse.NewEmpty()), label)
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func derivCall2(refs map[string]*relapse.Pattern, left, right *relapse.Pattern, label parser.Value) []*relapse.Pattern {
	l := derivCall(refs, left, label)
	r := derivCall(refs, right, label)
	return append(l, r...)
}

func derivReturns(refs map[string]*relapse.Pattern, originals []*relapse.Pattern, evaluated []*relapse.Pattern) []*relapse.Pattern {
	res := make([]*relapse.Pattern, len(originals))
	rest := evaluated
	for i, original := range originals {
		res[i], rest = derivReturn(refs, original, rest)
	}
	return res
}

func derivReturn(refs map[string]*relapse.Pattern, p *relapse.Pattern, patterns []*relapse.Pattern) (*relapse.Pattern, []*relapse.Pattern) {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return relapse.NewNot(relapse.NewZAny()), patterns
	case *relapse.ZAny:
		return relapse.NewZAny(), patterns
	case *relapse.TreeNode:
		if Nullable(refs, patterns[0]) {
			return relapse.NewEmpty(), patterns[1:]
		}
		return relapse.NewNot(relapse.NewZAny()), patterns[1:]
	case *relapse.LeafNode:
		if Nullable(refs, patterns[0]) {
			return relapse.NewEmpty(), patterns[1:]
		}
		return relapse.NewNot(relapse.NewZAny()), patterns[1:]
	case *relapse.Concat:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), patterns)
		leftConcat := relapse.NewConcat(l, v.GetRightPattern())
		if !Nullable(refs, v.GetLeftPattern()) {
			return leftConcat, leftRest
		}
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return relapse.NewOr(leftConcat, r), rightRest
	case *relapse.Or:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), patterns)
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return relapse.NewOr(l, r), rightRest
	case *relapse.And:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), patterns)
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return relapse.NewAnd(l, r), rightRest
	case *relapse.Interleave:
		l, leftRest := derivReturn(refs, v.GetLeftPattern(), patterns)
		r, rightRest := derivReturn(refs, v.GetRightPattern(), leftRest)
		return relapse.NewOr(relapse.NewInterleave(l, v.GetRightPattern()), relapse.NewInterleave(r, v.GetLeftPattern())), rightRest
	case *relapse.ZeroOrMore:
		c, rest := derivReturn(refs, v.GetPattern(), patterns)
		return relapse.NewConcat(c, p), rest
	case *relapse.Reference:
		return derivReturn(refs, refs[v.GetName()], patterns)
	case *relapse.Not:
		c, rest := derivReturn(refs, v.GetPattern(), patterns)
		return relapse.NewNot(c), rest
	case *relapse.Contains:
		return derivReturn(refs, relapse.NewConcat(relapse.NewZAny(), relapse.NewConcat(v.GetPattern(), relapse.NewZAny())), patterns)
	case *relapse.Optional:
		return derivReturn(refs, relapse.NewOr(v.GetPattern(), relapse.NewEmpty()), patterns)
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}
