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
					rules = append(rules, newReturn(v.src, state, v.name, v.dst))
				} else {
					rules = append(rules, newReturn(v.src, state, v.name, v.fail))
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
					tsdst := transForPattern(refs, v.dst)
					transitions = transitions.add(tsdst...)
					states[sdst] = v.dst
					//fmt.Printf("newState %s\n", sdst)
					allVisited = false
				}
			case *callAndReturn:
				sdst := v.dst.String()
				if _, ok := states[sdst]; !ok {
					tsdst := transForPattern(refs, v.dst)
					transitions = transitions.add(tsdst...)
					states[sdst] = v.dst
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
	p = interp.Simplify(refs, p)
	es, ns := getAllReachableExprs(refs, p)
	ts := make(transSet, 0, len(ns)+len(es))
	for i, n := range ns {
		d := deriv(refs, p, nil, n)
		ddst := interp.Simplify(refs, d.dst)
		dfail := interp.Simplify(refs, d.fail)
		dchild := d.child
		if dchild != nil {
			dchild = interp.Simplify(refs, dchild)
		} else {
			dchild = interp.Simplify(refs, relapse.NewNot(relapse.NewEmptySet()))
		}
		ts = ts.add(&callAndReturn{
			src:   p,
			name:  ns[i],
			child: dchild,
			dst:   ddst,
			fail:  dfail,
		})
	}
	for i, e := range es {
		d := deriv(refs, p, e, nil)
		ddst := interp.Simplify(refs, d.dst)
		ts = ts.add(&internal{
			src:  p,
			expr: es[i],
			dst:  ddst,
		})
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
	// given the deterministic nature in which child and dst are calculated.
	return true
}

type callAndReturn struct {
	src   *relapse.Pattern
	name  *relapse.NameExpr
	child *relapse.Pattern
	dst   *relapse.Pattern
	fail  *relapse.Pattern
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
	// given the deterministic nature in which child and dst are calculated.
	return true
}

type downup struct {
	child *relapse.Pattern
	dst   *relapse.Pattern
	fail  *relapse.Pattern
}

func concat(c *downup, p *relapse.Pattern) *downup {
	return &downup{
		child: c.child,
		dst:   relapse.NewConcat(c.dst, p),
		fail:  relapse.NewConcat(c.fail, p),
	}
}

func or(left *downup, right *downup) *downup {
	child := left.child
	if right.child != nil {
		if left.child != nil {
			child = relapse.NewOr(left.child, right.child)
		} else {
			child = right.child
		}
	}
	return &downup{
		child: child,
		dst:   relapse.NewOr(left.dst, right.dst),
		fail:  relapse.NewOr(left.fail, right.fail),
	}
}

func and(left *downup, right *downup) *downup {
	child := left.child
	if right.child != nil {
		if left.child != nil {
			child = relapse.NewAnd(left.child, right.child)
		} else {
			child = right.child
		}
	}
	return &downup{
		child: child,
		dst:   relapse.NewAnd(left.dst, right.dst),
		fail:  relapse.NewAnd(left.fail, right.fail),
	}
}

//TODO think about this again and write a few tests
func not(c *downup) *downup {
	child := c.child
	if child != nil {
		child = relapse.NewNot(child)
	}
	return &downup{
		child: child,
		dst:   relapse.NewNot(c.dst),
		fail:  relapse.NewNot(c.fail),
	}
}

func emptySet() *downup {
	return &downup{
		child: nil,
		dst:   relapse.NewEmptySet(),
		fail:  relapse.NewEmptySet(),
	}
}

func deriv(refs relapse.RefLookup, p *relapse.Pattern, expr *expr.Expr, name *relapse.NameExpr) *downup {
	d := func(pattern *relapse.Pattern) *downup {
		dp := deriv(refs, pattern, expr, name)
		fmt.Printf("dst  deriv %s -> %s\n", pattern.String(), dp.dst.String())
		fmt.Printf("fail deriv %s -> %s\n", pattern.String(), dp.fail.String())
		return dp
	}
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return emptySet()
	case *relapse.EmptySet:
		return emptySet()
	case *relapse.TreeNode:
		if name == nil {
			return emptySet()
		}
		if !v.GetName().Equal(name) {
			return emptySet()
		}
		return &downup{
			child: v.GetPattern(),
			dst:   relapse.NewEmpty(),
			fail:  relapse.NewEmptySet(),
		}
	case *relapse.LeafNode:
		if expr == nil {
			return emptySet()
		}
		if !v.GetExpr().Equal(expr) {
			return emptySet()
		}
		return &downup{
			child: nil,
			dst:   relapse.NewEmpty(),
			fail:  relapse.NewEmptySet(),
		}
	case *relapse.Concat:
		dl := d(v.GetLeftPattern())
		//fmt.Printf("Left: %s\n Right: %s\n dl.Child: %s\n dl.Dst: %s\n", v.GetLeftPattern(), v.GetRightPattern(), dl.child, dl.dst)
		c := concat(dl, v.GetRightPattern())
		//fmt.Printf(" cl.Child %s\n cl.Dst %s\n", c.child, c.dst)
		if !interp.Nullable(refs, v.GetLeftPattern()) {
			//fmt.Printf("NotNullable %s\n", v.GetLeftPattern())
			return c
		}
		o := or(c, d(v.GetRightPattern()))
		//fmt.Printf(" or.Child %s\n or.Dst %s\n", o.child, o.dst)
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
