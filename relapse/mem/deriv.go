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
	"io"

	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/intern"
)

func deriv(mem *Mem, patterns int, tree parser.Interface) (int, error) {
	for {
		if !mem.states.Get(patterns).Escapable {
			return patterns, nil
		}
		if err := tree.Next(); err != nil {
			if err == io.EOF {
				break
			} else {
				return 0, err
			}
		}
		callTree, err := mem.getCallTree(patterns)
		if err != nil {
			return 0, err
		}
		childPatterns, stackElm, err := mem.eval(callTree, tree)
		if err != nil {
			return 0, err
		}
		if !tree.IsLeaf() {
			tree.Down()
			childPatterns, err = deriv(mem, childPatterns, tree)
			if err != nil {
				return 0, err
			}
			tree.Up()
		}
		nullIndex := mem.getNullable(childPatterns)
		patterns, err = mem.getReturn(stackElm, nullIndex)
		if err != nil {
			return 0, err
		}
	}
	return patterns, nil
}

func derivCalls(construct intern.Construct, patterns []*intern.Pattern) []*ifExpr {
	res := []*ifExpr{}
	for _, pattern := range patterns {
		cs := derivCall(construct, pattern)
		res = append(res, cs...)
	}
	return res
}

func derivCall(c intern.Construct, p *intern.Pattern) []*ifExpr {
	switch p.Type {
	case intern.Empty:
		return []*ifExpr{}
	case intern.ZAny:
		return []*ifExpr{}
	case intern.Node:
		return []*ifExpr{{p.Func, p.Patterns[0], c.NewNotZAny()}}
	case intern.Concat:
		res := []*ifExpr{}
		for i := range p.Patterns {
			ps := derivCall(c, p.Patterns[i])
			res = append(res, ps...)
			if !p.Patterns[i].Nullable() {
				break
			}
		}
		return res
	case intern.Or:
		return derivCallAll(c, p.Patterns)
	case intern.And:
		return derivCallAll(c, p.Patterns)
	case intern.Interleave:
		return derivCallAll(c, p.Patterns)
	case intern.ZeroOrMore:
		return derivCall(c, p.Patterns[0])
	case intern.Reference:
		return derivCall(c, c.Deref(p.Ref))
	case intern.Not:
		return derivCall(c, p.Patterns[0])
	case intern.Contains:
		return derivCall(c, p.Patterns[0])
	case intern.Optional:
		return derivCall(c, p.Patterns[0])
	}
	panic(fmt.Sprintf("unknown pattern typ %d", p.Type))
}

func derivCallAll(c intern.Construct, ps []*intern.Pattern) []*ifExpr {
	res := []*ifExpr{}
	for i := range ps {
		pss := derivCall(c, ps[i])
		res = append(res, pss...)
	}
	return res
}

func derivReturns(c intern.Construct, originals []*intern.Pattern, nullable []bool) ([]*intern.Pattern, error) {
	res := make([]*intern.Pattern, len(originals))
	rest := nullable
	var err error
	for i, original := range originals {
		res[i], rest, err = derivReturn(c, original, rest)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func derivReturn(c intern.Construct, p *intern.Pattern, nullable []bool) (*intern.Pattern, []bool, error) {
	switch p.Type {
	case intern.Empty:
		return c.NewNotZAny(), nullable, nil
	case intern.ZAny:
		return c.NewZAny(), nullable, nil
	case intern.Node:
		if nullable[0] {
			return c.NewEmpty(), nullable[1:], nil
		}
		return c.NewNotZAny(), nullable[1:], nil
	case intern.Concat:
		rest := nullable
		orPatterns := make([]*intern.Pattern, 0, len(p.Patterns))
		var err error
		for i := range p.Patterns {
			var ret *intern.Pattern
			ret, rest, err = derivReturn(c, p.Patterns[i], rest)
			if err != nil {
				return nil, nil, err
			}
			concatPat, err := c.NewConcat(append([]*intern.Pattern{ret}, p.Patterns[i+1:]...))
			if err != nil {
				return nil, nil, err
			}
			orPatterns = append(orPatterns, concatPat)
			if !p.Patterns[i].Nullable() {
				break
			}
		}
		o, err := c.NewOr(orPatterns)
		return o, rest, err
	case intern.Or:
		rest := nullable
		orPatterns := make([]*intern.Pattern, len(p.Patterns))
		var err error
		for i := range p.Patterns {
			orPatterns[i], rest, err = derivReturn(c, p.Patterns[i], rest)
			if err != nil {
				return nil, nil, err
			}
		}
		o, err := c.NewOr(orPatterns)
		return o, rest, err
	case intern.And:
		rest := nullable
		andPatterns := make([]*intern.Pattern, len(p.Patterns))
		var err error
		for i := range p.Patterns {
			andPatterns[i], rest, err = derivReturn(c, p.Patterns[i], rest)
			if err != nil {
				return nil, nil, err
			}
		}
		a, err := c.NewAnd(andPatterns)
		return a, rest, err
	case intern.Interleave:
		rest := nullable
		orPatterns := make([]*intern.Pattern, len(p.Patterns))
		var err error
		for i := range p.Patterns {
			interleaves := make([]*intern.Pattern, len(p.Patterns))
			copy(interleaves, p.Patterns)
			interleaves[i], rest, err = derivReturn(c, p.Patterns[i], rest)
			if err != nil {
				return nil, nil, err
			}
			orPatterns[i], err = c.NewInterleave(interleaves)
			if err != nil {
				return nil, nil, err
			}
		}
		o, err := c.NewOr(orPatterns)
		return o, rest, err
	case intern.ZeroOrMore:
		pp, rest, err := derivReturn(c, p.Patterns[0], nullable)
		if err != nil {
			return nil, nil, err
		}
		ppp, err := c.NewConcat([]*intern.Pattern{pp, p})
		return ppp, rest, err
	case intern.Reference:
		return derivReturn(c, c.Deref(p.Ref), nullable)
	case intern.Not:
		pp, rest, err := derivReturn(c, p.Patterns[0], nullable)
		if err != nil {
			return nil, nil, err
		}
		ppp, err := c.NewNot(pp)
		return ppp, rest, err
	case intern.Contains:
		orPatterns := make([]*intern.Pattern, 0, 3)
		orPatterns = append(orPatterns, p)
		ret, rest, err := derivReturn(c, p.Patterns[0], nullable)
		if err != nil {
			return nil, nil, err
		}
		cp, err := c.NewConcat([]*intern.Pattern{ret, c.NewZAny()})
		if err != nil {
			return nil, nil, err
		}
		orPatterns = append(orPatterns, cp)
		if p.Patterns[0].Nullable() {
			orPatterns = append(orPatterns, c.NewZAny())
		}
		o, err := c.NewOr(orPatterns)
		return o, rest, err
	case intern.Optional:
		return derivReturn(c, p.Patterns[0], nullable)
	}
	panic(fmt.Sprintf("unknown pattern typ %d", p.Type))
}
