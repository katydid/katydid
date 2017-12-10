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
	"fmt"
	"io"

	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/compose"
)

//Interpret interprets the grammar given the parser and returns whether the parser is valid given the grammar.
//Interpret uses derivatives and simplification to recusively derive the resulting grammar.
//This resulting grammar's nullability then represents the result of the function.
//This implementation does not handle immediate recursion, see the HasRecursion function.
func Interpret(g *ast.Grammar, record bool, parser parser.Interface) (bool, error) {
	c := NewConstructor()
	if record {
		c = NewConstructorOptimizedForRecords()
	}
	main, err := c.AddGrammar(g)
	if err != nil {
		return false, err
	}
	finals, err := deriv(c, []*Pattern{main}, parser)
	if err != nil {
		return false, err
	}
	return finals[0].nullable, nil
}

func escapable(patterns []*Pattern) bool {
	for _, p := range patterns {
		if isZAny(p) {
			continue
		}
		if isNotZAny(p) {
			continue
		}
		return true
	}
	return false
}

func deriv(c Construct, patterns []*Pattern, tree parser.Interface) ([]*Pattern, error) {
	var resPatterns []*Pattern = patterns
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
		childPatterns, err := derivCalls(c, resPatterns, tree)
		if err != nil {
			return nil, err
		}
		if tree.IsLeaf() {
			//do nothing
		} else {
			tree.Down()
			zchild, zi := Zip(childPatterns)
			zchild, err = deriv(c, zchild, tree)
			if err != nil {
				return nil, err
			}
			childPatterns = Unzip(zchild, zi)
			tree.Up()
		}
		resPatterns, err = derivReturns(c, resPatterns, childPatterns)
		if err != nil {
			return nil, err
		}
	}
	return resPatterns, nil
}

func derivCalls(c Construct, patterns []*Pattern, label parser.Value) ([]*Pattern, error) {
	res := []*Pattern{}
	for _, pattern := range patterns {
		ps, err := derivCall(c, pattern, label)
		if err != nil {
			return nil, err
		}
		res = append(res, ps...)
	}
	return res, nil
}

func derivCall(c Construct, p *Pattern, label parser.Value) ([]*Pattern, error) {
	switch p.Type {
	case Empty:
		return []*Pattern{}, nil
	case ZAny:
		return []*Pattern{}, nil
	case Node:
		f, err := compose.NewBoolFunc(p.Func)
		if err != nil {
			return nil, err
		}
		eval, err := f.Eval(label)
		if err != nil {
			return nil, err
		}
		if eval {
			return []*Pattern{p.Patterns[0]}, nil
		}
		return []*Pattern{notzany}, nil
	case Concat:
		res := []*Pattern{}
		for i := range p.Patterns {
			ps, err := derivCall(c, p.Patterns[i], label)
			if err != nil {
				return nil, err
			}
			res = append(res, ps...)
			if !p.Patterns[i].nullable {
				break
			}
		}
		return res, nil
	case Or:
		return derivCallAll(c, p.Patterns, label)
	case And:
		return derivCallAll(c, p.Patterns, label)
	case Interleave:
		return derivCallAll(c, p.Patterns, label)
	case ZeroOrMore:
		return derivCall(c, p.Patterns[0], label)
	case Reference:
		return derivCall(c, c.Deref(p.Ref), label)
	case Not:
		return derivCall(c, p.Patterns[0], label)
	case Contains:
		return derivCall(c, p.Patterns[0], label)
	case Optional:
		return derivCall(c, p.Patterns[0], label)
	}
	panic(fmt.Sprintf("unknown pattern typ %d", p.Type))
}

func derivCallAll(c Construct, ps []*Pattern, label parser.Value) ([]*Pattern, error) {
	res := []*Pattern{}
	for i := range ps {
		ps, err := derivCall(c, ps[i], label)
		if err != nil {
			return nil, err
		}
		res = append(res, ps...)
	}
	return res, nil
}

func derivReturns(c Construct, originals []*Pattern, evaluated []*Pattern) ([]*Pattern, error) {
	res := make([]*Pattern, len(originals))
	rest := evaluated
	var err error
	for i, original := range originals {
		res[i], rest, err = derivReturn(c, original, rest)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func derivReturn(c Construct, p *Pattern, patterns []*Pattern) (*Pattern, []*Pattern, error) {
	switch p.Type {
	case Empty:
		return c.NewNot(c.NewZAny()), patterns, nil
	case ZAny:
		return c.NewZAny(), patterns, nil
	case Node:
		if patterns[0].nullable {
			return c.NewEmpty(), patterns[1:], nil
		}
		return c.NewNot(c.NewZAny()), patterns[1:], nil
	case Concat:
		rest := patterns
		orPatterns := make([]*Pattern, 0, len(p.Patterns))
		var err error
		for i := range p.Patterns {
			var ret *Pattern
			ret, rest, err = derivReturn(c, p.Patterns[i], rest)
			if err != nil {
				return nil, nil, err
			}
			concatPat := c.NewConcat(append([]*Pattern{ret}, p.Patterns[i+1:]...))
			orPatterns = append(orPatterns, concatPat)
			if !p.Patterns[i].nullable {
				break
			}
		}
		o, err := c.NewOr(orPatterns)
		return o, rest, err
	case Or:
		rest := patterns
		orPatterns := make([]*Pattern, len(p.Patterns))
		var err error
		for i := range p.Patterns {
			orPatterns[i], rest, err = derivReturn(c, p.Patterns[i], rest)
			if err != nil {
				return nil, nil, err
			}
		}
		o, err := c.NewOr(orPatterns)
		return o, rest, err
	case And:
		rest := patterns
		andPatterns := make([]*Pattern, len(p.Patterns))
		var err error
		for i := range p.Patterns {
			andPatterns[i], rest, err = derivReturn(c, p.Patterns[i], rest)
			if err != nil {
				return nil, nil, err
			}
		}
		a, err := c.NewAnd(andPatterns)
		return a, rest, err
	case Interleave:
		rest := patterns
		orPatterns := make([]*Pattern, len(p.Patterns))
		var err error
		for i := range p.Patterns {
			interleaves := make([]*Pattern, len(p.Patterns))
			copy(interleaves, p.Patterns)
			interleaves[i], rest, err = derivReturn(c, p.Patterns[i], rest)
			if err != nil {
				return nil, nil, err
			}
			orPatterns[i] = c.NewInterleave(interleaves)
		}
		o, err := c.NewOr(orPatterns)
		return o, rest, err
	case ZeroOrMore:
		pp, rest, err := derivReturn(c, p.Patterns[0], patterns)
		return c.NewConcat([]*Pattern{pp, p}), rest, err
	case Reference:
		return derivReturn(c, c.Deref(p.Ref), patterns)
	case Not:
		pp, rest, err := derivReturn(c, p.Patterns[0], patterns)
		return c.NewNot(pp), rest, err
	case Contains:
		orPatterns := make([]*Pattern, 0, 3)
		orPatterns = append(orPatterns, p)
		ret, rest, err := derivReturn(c, p.Patterns[0], patterns)
		if err != nil {
			return nil, nil, err
		}
		orPatterns = append(orPatterns, c.NewConcat([]*Pattern{ret, zany}))
		if p.Patterns[0].nullable {
			orPatterns = append(orPatterns, zany)
		}
		o, err := c.NewOr(orPatterns)
		return o, rest, err
	case Optional:
		return derivReturn(c, p.Patterns[0], patterns)
	}
	panic(fmt.Sprintf("unknown pattern typ %d", p.Type))
}
