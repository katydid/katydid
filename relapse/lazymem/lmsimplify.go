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

package lazymem

import (
	"fmt"
)

func Simplify(p *Pattern) *Pattern {
	return simplify(make(map[*Pattern]struct{}), p)
}

func simplify(visited map[*Pattern]struct{}, p *Pattern) *Pattern {
	if !p.HasHead() {
		return p
	}
	if _, ok := visited[p]; ok {
		return p
	}
	visited[p] = struct{}{}
	typ := p.Head().GetValue()
	switch v := typ.(type) {
	case *Empty:
		return p
	case *Node:
		newp := simplify(visited, v.Pattern)
		return NewNode(v.F, v.Name, v.Expr, newp)
	case *Concat:
		return simplifyConcat(
			simplify(visited, v.Left),
			simplify(visited, v.Right),
		)
	case *Or:
		return simplifyOr(
			simplify(visited, v.Left),
			simplify(visited, v.Right),
		)
	case *And:
		return simplifyAnd(
			simplify(visited, v.Left),
			simplify(visited, v.Right),
		)
	case *ZeroOrMore:
		return NewZeroOrMore(simplify(visited, v.Pattern))
	case *Not:
		return simplifyNot(simplify(visited, v.Pattern))
	case *ZAny:
		return p
	case *Contains:
		return simplifyContains(simplify(visited, v.Pattern))
	case *Optional:
		return simplifyOptional(simplify(visited, v.Pattern))
	case *Interleave:
		return simplifyInterleave(
			simplify(visited, v.Left),
			simplify(visited, v.Right),
		)
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func isNotZany(p *Pattern) bool {
	return p.Head().Not != nil && p.Head().Not.Pattern.HasHead() && p.Head().Not.Pattern.Head().ZAny != nil
}

func isEmpty(p *Pattern) bool {
	return p.Head().Empty != nil
}

func isZany(p *Pattern) bool {
	return p.Head().ZAny != nil
}

func simplifyConcat(p1, p2 *Pattern) *Pattern {
	if !p1.HasHead() || !p2.HasHead() {
		return NewConcat(p1, p2)
	}
	if isNotZany(p1) || isNotZany(p2) {
		return NewNot(NewZAny())
	}
	if p1.Head().Concat != nil {
		return simplifyConcat(
			p1.Head().Concat.Left,
			NewConcat(p1.Head().Concat.Right, p2),
		)
	}
	if isEmpty(p1) {
		return p2
	}
	if isEmpty(p2) {
		return p1
	}
	if isZany(p1) && p2.Head().Concat != nil {
		if l := p2.Head().Concat.Left; isZany(p2.Head().Concat.Right) {
			return NewContains(l)
		}
	}
	return NewConcat(p1, p2)
}

func simplifyOr(p1, p2 *Pattern) *Pattern {
	if !p1.HasHead() || !p2.HasHead() {
		return NewOr(p1, p2)
	}
	if isNotZany(p1) {
		return p2
	}
	if isNotZany(p2) {
		return p1
	}
	if isZany(p1) || isZany(p2) {
		return NewZAny()
	}
	/* TODO
	if isEmpty(p1) && Nullable(p2) {
		return p2
	}
	if isEmpty(p2) && Nullable(p1) {
		return p1
	}
	if p1.Head().Equal(p2.Head()) {
		return p1
	}
	if p2.Head().Or != nil {
		if p1.Equal(p2.Head().Or.Left) || p1.Equal(p2.Head().Or.Right) {
			return simplifyOr(p2.Head().Or.Left, p2.Head().Or.Right)
		}
	}
	*/
	return NewOr(p1, p2)
}

func simplifyAnd(p1, p2 *Pattern) *Pattern {
	if !p1.HasHead() || !p2.HasHead() {
		return NewAnd(p1, p2)
	}
	if isNotZany(p1) || isNotZany(p2) {
		return NewNot(NewZAny())
	}
	if isZany(p1) {
		return p2
	}
	if isZany(p2) {
		return p1
	}
	/*TODO
	if isEmpty(p1) {
		if Nullable(p2) {
			return NewEmpty()
		} else {
			return NewNot(NewZAny())
		}
	}
	if isEmpty(p2) {
		if Nullable(p1) {
			return NewEmpty()
		} else {
			return NewNot(NewZAny())
		}
	}
	if p1.Equal(p2) {
		return p1
	}
	*/
	return NewAnd(p1, p2)
}

func simplifyNot(p *Pattern) *Pattern {
	if !p.HasHead() {
		return NewNot(p)
	}
	if p.Head().Not != nil {
		return p.Head().Not.Pattern
	}
	return NewNot(p)
}

//TODO we can do better
func simplifyOptional(p *Pattern) *Pattern {
	if !p.HasHead() {
		return NewOptional(p)
	}
	if isEmpty(p) {
		return NewEmpty()
	}
	return NewOptional(p)
}

func simplifyInterleave(p1, p2 *Pattern) *Pattern {
	if !p1.HasHead() || !p2.HasHead() {
		return NewInterleave(p1, p2)
	}
	if isNotZany(p1) || isNotZany(p2) {
		return NewNot(NewZAny())
	}
	if isEmpty(p1) {
		return p2
	}
	if isEmpty(p2) {
		return p1
	}
	if isZany(p1) && isZany(p2) {
		return p1
	}
	return NewInterleave(p1, p2)
}

func simplifyContains(p *Pattern) *Pattern {
	if !p.HasHead() {
		return NewContains(p)
	}
	if isEmpty(p) || isZany(p) {
		return NewZAny()
	}
	if isNotZany(p) {
		return p
	}
	return NewContains(p)
}
