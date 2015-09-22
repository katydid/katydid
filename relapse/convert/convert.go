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

package convert

import (
	"fmt"
	"github.com/katydid/katydid/expr/ast"
	eparser "github.com/katydid/katydid/expr/parser"
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/interp"
	"github.com/katydid/katydid/viper/ast"
	"strconv"
)

func ToRules(g *relapse.Grammar) *viper.Rules {
	refs := relapse.NewRefsLookup(g)
	ts, states := toTrans(refs)
	rules := []*viper.Rule{}
	rules = append(rules, newStart(interp.Simplify(refs, refs["main"])))
	for _, state := range states {
		if interp.Nullable(refs, state) {
			rules = append(rules, newFinal(interp.Simplify(refs, state)))
		}
	}
	for _, t := range ts {
		switch v := t.(type) {
		case *internal:
			rules = append(rules, newInternal(v.src, v.expr, v.dst))
		case *callAndReturn:
			rules = append(rules,
				newCall(v.src, v.name, v.src, v.child),
			)
			for _, state := range states {
				if interp.Nullable(refs, state) {
					rules = append(rules, newReturn(v.src, state, v.name, v.retNull))
				} else {
					if v.retElse != nil {
						rules = append(rules, newReturn(v.src, state, v.name, v.retElse))
					} else {
						//TODO
					}
				}
			}
		}
	}
	return viper.NewRules(rules...)
}

func toTrans(refs relapse.RefLookup) (transitions transSet, states map[string]*relapse.Pattern) {
	fmt.Printf("ROOT = %s\n", refs["main"])
	p := interp.Simplify(refs, refs["main"])
	transitions = transForPattern(refs, p)
	states = make(map[string]*relapse.Pattern)
	states[p.String()] = interp.Simplify(refs, p)
	//fmt.Printf("newState = %s\n", p)
	//fmt.Printf("newState = %s\n", relapse.NewEmptySet())
	allVisited := false
	for !allVisited {
		allVisited = true
		for _, t := range transitions {
			switch v := t.(type) {
			case *internal:
				sdst := v.dst.String()
				if _, ok := states[sdst]; !ok {
					//tsdst := transForPattern(refs, v.dst)
					//transitions = transitions.add(tsdst...)
					states[sdst] = v.dst
					//fmt.Printf("newState %s\n", sdst)
					allVisited = false
				}
			case *callAndReturn:
				sdst := v.retNull.String()
				if _, ok := states[sdst]; !ok {
					tsnul := transForPattern(refs, v.retNull)
					transitions = transitions.add(tsnul...)
					states[sdst] = v.retNull
					//fmt.Printf("newState %s\n", sdst)
					allVisited = false
				}
				schild := v.child.String()
				if _, ok := states[schild]; !ok {
					tschild := transForPattern(refs, v.child)
					transitions = transitions.add(tschild...)
					states[schild] = v.child
					//fmt.Printf("newState %s\n", schild)
					allVisited = false
				}
			}
		}
	}
	return transitions, states
}

func newCall(src *relapse.Pattern, name *relapse.NameExpr, parentDst *relapse.Pattern, childDst *relapse.Pattern) *viper.Rule {
	return viper.NewCall(
		strconv.Quote(src.String()),
		nameToExpr(name),
		strconv.Quote(parentDst.String()),
		strconv.Quote(childDst.String()),
	)
}

func newReturn(parentSrc *relapse.Pattern, childSrc *relapse.Pattern, name *relapse.NameExpr, dst *relapse.Pattern) *viper.Rule {
	return viper.NewReturn(
		strconv.Quote(parentSrc.String()),
		strconv.Quote(childSrc.String()),
		nameToExpr(name),
		strconv.Quote(dst.String()),
	)
}

func newInternal(src *relapse.Pattern, expr *expr.Expr, dst *relapse.Pattern) *viper.Rule {
	return viper.NewInternal(
		strconv.Quote(src.String()),
		expr,
		strconv.Quote(dst.String()),
	)
}

func newStart(p *relapse.Pattern) *viper.Rule {
	return viper.NewStart(strconv.Quote(p.String()))
}

func newFinal(p *relapse.Pattern) *viper.Rule {
	return viper.NewFinal(strconv.Quote(p.String()))
}

func transForPattern(refs relapse.RefLookup, p *relapse.Pattern) transSet {
	simplify := func(p *relapse.Pattern) *relapse.Pattern {
		return interp.Simplify(refs, p)
	}
	p = interp.Simplify(refs, p)
	es, ns := getAllReachableExprs(refs, p)
	ts := make(transSet, 0, len(ns)+len(es))
	for i, n := range ns {
		dn := deriv(refs, p, nil, n)
		d, ok := dn.(*callAndReturn)
		if ok {
			dNull := simplify(d.retNull)
			dElse := applyIfNotNil(simplify, d.retElse)
			dchild := applyIfNotNil(simplify, d.child)
			if dchild == nil {
				dchild = interp.Simplify(refs, relapse.NewNot(relapse.NewEmptySet()))
			}
			ts = ts.add(&callAndReturn{
				src:     p,
				name:    ns[i],
				child:   dchild,
				retNull: dNull,
				retElse: dElse,
			})
		} else {
			//TODO
		}
	}
	for i, e := range es {
		de := deriv(refs, p, e, nil)
		d, ok := de.(*internal)
		if ok {
			ddst := interp.Simplify(refs, d.dst)
			ts = ts.add(&internal{
				src:  p,
				expr: es[i],
				dst:  ddst,
			})
		} else {
			//TODO
		}
	}
	return ts
}

func newTrue() *expr.Expr {
	t := true
	return &expr.Expr{Terminal: &expr.Terminal{BoolValue: &t}}
}

func getAllReachableExprs(refs relapse.RefLookup, p *relapse.Pattern) ([]*expr.Expr, []*relapse.NameExpr) {
	all := getAllExprs(refs.Clone(), refs, p)
	catchAllExpr := newAllExpr(newTrue())
	catchAllName := newAllName(relapse.NewAnyName())
	all = all.add(catchAllExpr).add(catchAllName)
	return all.exprs, all.names
}

type allExprs struct {
	exprs    []*expr.Expr
	names    []*relapse.NameExpr
	setExprs map[string]struct{}
	setNames map[string]struct{}
}

func (this *allExprs) add(that *allExprs) *allExprs {
	for i, e := range that.exprs {
		if _, ok := this.setExprs[e.String()]; !ok {
			this.exprs = append(this.exprs, that.exprs[i])
			this.setExprs[e.String()] = struct{}{}
		}
	}
	for i, n := range that.names {
		if _, ok := this.setNames[n.String()]; !ok {
			this.names = append(this.names, that.names[i])
			this.setNames[n.String()] = struct{}{}
		}
	}
	return this
}

func newAllExpr(e *expr.Expr) *allExprs {
	return &allExprs{
		exprs: []*expr.Expr{e},
		setExprs: map[string]struct{}{
			e.String(): struct{}{},
		},
		names:    []*relapse.NameExpr{},
		setNames: map[string]struct{}{},
	}
}

func newAllNil() *allExprs {
	return &allExprs{
		exprs:    []*expr.Expr{},
		setExprs: map[string]struct{}{},
		names:    []*relapse.NameExpr{},
		setNames: map[string]struct{}{},
	}
}

func newAllName(n *relapse.NameExpr) *allExprs {
	return &allExprs{
		exprs:    []*expr.Expr{},
		setExprs: map[string]struct{}{},
		names:    []*relapse.NameExpr{n},
		setNames: map[string]struct{}{
			n.String(): struct{}{},
		},
	}
}

func getAllExprs(bookKeepingRefs relapse.RefLookup, allRefs relapse.RefLookup, p *relapse.Pattern) *allExprs {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return newAllNil()
	case *relapse.EmptySet:
		return newAllNil()
	case *relapse.TreeNode:
		return newAllName(v.GetName())
	case *relapse.LeafNode:
		return newAllExpr(v.GetExpr())
	case *relapse.Concat:
		res := getAllExprs(bookKeepingRefs, allRefs, v.GetLeftPattern())
		if interp.Nullable(allRefs, v.GetLeftPattern()) {
			right := getAllExprs(bookKeepingRefs, allRefs, v.GetRightPattern())
			res = res.add(right)
		}
		return res
	case *relapse.Or:
		left := getAllExprs(bookKeepingRefs, allRefs, v.GetLeftPattern())
		right := getAllExprs(bookKeepingRefs, allRefs, v.GetRightPattern())
		return left.add(right)
	case *relapse.And:
		left := getAllExprs(bookKeepingRefs, allRefs, v.GetLeftPattern())
		right := getAllExprs(bookKeepingRefs, allRefs, v.GetRightPattern())
		return left.add(right)
	case *relapse.ZeroOrMore:
		return getAllExprs(bookKeepingRefs, allRefs, v.GetPattern())
	case *relapse.Reference:
		r, ok := bookKeepingRefs[v.GetName()]
		if !ok {
			return newAllNil()
		}
		delete(bookKeepingRefs, v.GetName())
		return getAllExprs(bookKeepingRefs, allRefs, r)
	case *relapse.Not:
		return getAllExprs(bookKeepingRefs, allRefs, v.GetPattern())
	case *relapse.ZAny:
		return getAllExprs(bookKeepingRefs, allRefs, relapse.NewNot(relapse.NewEmptySet()))
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}

func nameToExpr(n *relapse.NameExpr) *expr.Expr {
	f := nameToFunc(n)
	fs := funcs.Sprint(f)
	ex, err := eparser.NewParser().ParseExpr(fs)
	if err != nil {
		panic(err)
	}
	return ex
}

func nameToFunc(n *relapse.NameExpr) funcs.Bool {
	typ := n.GetValue()
	switch v := typ.(type) {
	case *relapse.Name:
		return funcs.StringEq(funcs.StringVar(), funcs.StringConst(v.GetName()))
	case *relapse.AnyName:
		return funcs.BoolConst(true)
	case *relapse.AnyNameExcept:
		return funcs.Not(nameToFunc(v.GetExcept()))
	case *relapse.NameChoice:
		return funcs.Or(nameToFunc(v.GetLeft()), nameToFunc(v.GetRight()))
	}
	panic(fmt.Sprintf("unknown name expr typ %T", typ))
}

type transSet []trans

func (as transSet) add(bs ...trans) transSet {
	for _, b := range bs {
		if !as.contains(b) {
			as = append(as, b)
		}
	}
	return as
}

func (as transSet) contains(b trans) bool {
	for _, a := range as {
		if a.Equal(b) {
			return true
		}
	}
	return false
}

type trans interface {
	Equal(trans) bool
	isTrans()
}

func (this *internal) isTrans()      {}
func (this *callAndReturn) isTrans() {}

type internal struct {
	src  *relapse.Pattern
	expr *expr.Expr
	dst  *relapse.Pattern
}

func (a *internal) Equal(that trans) bool {
	b, ok := that.(*internal)
	if !ok {
		return false
	}
	if !a.src.Equal(b.src) {
		return false
	}
	if !a.expr.Equal(b.expr) {
		return false
	}
	//The rest of the fields do not need to be checked,
	// given the deterministic nature in which the rest are calculated.
	return true
}

func (a *internal) String() string {
	return fmt.Sprintf("internal{\n\tsrc: %s,\n\texpr: %s,\n\tdst: %s\n}", a.src, a.expr, a.dst)
}

type callAndReturn struct {
	src     *relapse.Pattern
	name    *relapse.NameExpr
	child   *relapse.Pattern
	retNull *relapse.Pattern
	retElse *relapse.Pattern
}

func (a *callAndReturn) String() string {
	s := fmt.Sprintf("callAndReturn{\n\tsrc: %s,\n\tname: %s", a.src, a.name)
	if a.child != nil {
		s += fmt.Sprintf(",\n\tchild: %s", a.child)
	}
	s += fmt.Sprintf(",\n\tretNull: %s", a.retNull)
	if a.retElse != nil {
		s += fmt.Sprintf(",\n\tretElse: %s", a.retElse)
	}
	s += "\n}"
	return s
}

func (a *callAndReturn) Equal(that trans) bool {
	b, ok := that.(*callAndReturn)
	if !ok {
		return false
	}
	if !a.src.Equal(b.src) {
		return false
	}
	if !a.name.Equal(b.name) {
		return false
	}
	//The rest of the fields do not need to be checked,
	// given the deterministic nature in which the rest are calculated.
	return true
}

func applyIfNotNil(f func(*relapse.Pattern) *relapse.Pattern, p *relapse.Pattern) *relapse.Pattern {
	if p == nil {
		return p
	}
	return f(p)
}

func combineIfNotNil(f func(...*relapse.Pattern) *relapse.Pattern, left, right *relapse.Pattern) *relapse.Pattern {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return f(left, right)
}

func concat(t trans, p *relapse.Pattern) trans {
	switch v := t.(type) {
	case *internal:
		return &internal{
			src:  relapse.NewConcat(v.src, p),
			expr: v.expr,
			dst:  relapse.NewConcat(v.dst, p),
		}
	case *callAndReturn:
		return &callAndReturn{
			src:     relapse.NewConcat(v.src, p),
			name:    v.name,
			child:   v.child,
			retNull: relapse.NewConcat(v.retNull, p),
			retElse: combineIfNotNil(relapse.NewConcat, v.retElse, p),
		}
	}
	panic(fmt.Sprintf("unknown trans type %T", t))
}

func or(leftt, rightt trans) trans {
	switch left := leftt.(type) {
	case *internal:
		switch right := rightt.(type) {
		case *internal:
			return &internal{
				src:  relapse.NewOr(left.src, right.src),
				expr: left.expr,
				dst:  relapse.NewOr(left.dst, right.dst),
			}
		case *callAndReturn:
			panic(fmt.Sprintf("todo: %s or %s", leftt, rightt))
		}
		panic(fmt.Sprintf("unknown trans type %T", rightt))
	case *callAndReturn:
		switch right := rightt.(type) {
		case *internal:
			panic(fmt.Sprintf("todo: %s or %s", leftt, rightt))
		case *callAndReturn:
			return &callAndReturn{
				src:     relapse.NewOr(left.src, right.src),
				name:    left.name,
				child:   combineIfNotNil(relapse.NewOr, left.child, right.child),
				retNull: relapse.NewOr(left.retNull, right.retNull),
				retElse: combineIfNotNil(relapse.NewOr, left.retElse, right.retElse),
			}
		}
		panic(fmt.Sprintf("unknown trans type %T", rightt))
	}
	panic(fmt.Sprintf("unknown trans type %T", leftt))
}

func and(leftt, rightt trans) trans {
	switch left := leftt.(type) {
	case *internal:
		switch right := rightt.(type) {
		case *internal:
			return &internal{
				src:  relapse.NewAnd(left.src, right.src),
				expr: left.expr,
				dst:  relapse.NewAnd(left.dst, right.dst),
			}
		case *callAndReturn:
			panic("todo")
		}
		panic(fmt.Sprintf("unknown trans type %T", rightt))
	case *callAndReturn:
		switch right := rightt.(type) {
		case *internal:
			panic("todo")
		case *callAndReturn:
			return &callAndReturn{
				src:     relapse.NewAnd(left.src, right.src),
				name:    left.name,
				child:   combineIfNotNil(relapse.NewAnd, left.child, right.child),
				retNull: relapse.NewAnd(left.retNull, right.retNull),
				retElse: combineIfNotNil(relapse.NewAnd, left.retElse, right.retElse),
			}
		}
		panic(fmt.Sprintf("unknown trans type %T", rightt))
	}
	panic(fmt.Sprintf("unknown trans type %T", leftt))
}

func not(t trans) trans {
	switch v := t.(type) {
	case *internal:
		return &internal{
			src:  relapse.NewNot(v.src),
			expr: v.expr,
			dst:  relapse.NewNot(v.dst),
		}
	case *callAndReturn:
		//TODO think about this, why is this a special case
		//     this should probably have been an internal
		return &callAndReturn{
			src:     relapse.NewNot(v.src),
			name:    v.name,
			child:   applyIfNotNil(relapse.NewNot, v.child),
			retNull: relapse.NewNot(v.retNull),
			retElse: applyIfNotNil(relapse.NewNot, v.retElse),
		}
	}
	panic(fmt.Sprintf("unknown trans type %T", t))
}

func deriv(refs relapse.RefLookup, p *relapse.Pattern, expr *expr.Expr, name *relapse.NameExpr) trans {
	d := func(pattern *relapse.Pattern) trans {
		dp := deriv(refs, pattern, expr, name)
		fmt.Printf("deriv %s\n", dp)
		return dp
	}
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty, *relapse.EmptySet:
		if expr != nil {
			return &internal{
				src:  p,
				expr: expr,
				dst:  relapse.NewEmptySet(),
			}
		} else {
			//TODO this should probably be an internal
			return &callAndReturn{
				src:     p,
				name:    name,
				child:   nil,
				retNull: relapse.NewEmptySet(),
				retElse: relapse.NewEmptySet(),
			}
		}
	case *relapse.TreeNode:
		if name != nil {
			nul := relapse.NewEmptySet()
			var els *relapse.Pattern = relapse.NewEmptySet() // TODO this should probably be an internal
			var child *relapse.Pattern = nil
			if v.GetName().Equal(name) {
				nul = relapse.NewEmpty()
				els = relapse.NewEmptySet()
				child = v.GetPattern()
			}
			return &callAndReturn{
				src:     p,
				name:    name,
				child:   child,
				retNull: nul,
				retElse: els,
			}
		}
		return &internal{
			src:  p,
			expr: expr,
			dst:  relapse.NewEmptySet(),
		}
	case *relapse.LeafNode:
		if expr != nil {
			dst := relapse.NewEmptySet()
			if v.GetExpr().Equal(expr) {
				dst = relapse.NewEmpty()
			}
			return &internal{
				src:  p,
				expr: expr,
				dst:  dst,
			}
		}
		return &callAndReturn{
			src:     p,
			name:    name,
			child:   nil,
			retNull: relapse.NewEmptySet(),
			retElse: relapse.NewEmptySet(),
		}
	case *relapse.Concat:
		dl := d(v.GetLeftPattern())
		c := concat(dl, v.GetRightPattern())
		if !interp.Nullable(refs, v.GetLeftPattern()) {
			return c
		}
		o := or(c, d(v.GetRightPattern()))
		return o
	case *relapse.Or:
		return or(d(v.GetLeftPattern()), d(v.GetRightPattern()))
	case *relapse.And:
		return and(d(v.GetLeftPattern()), d(v.GetRightPattern()))
	case *relapse.ZeroOrMore:
		return concat(d(v.GetPattern()), relapse.NewZeroOrMore(v.GetPattern()))
	case *relapse.Reference:
		return d(refs[v.GetName()])
	case *relapse.Not:
		return not(d(v.GetPattern()))
	}
	panic(fmt.Sprintf("unknown pattern typ %T", typ))
}
