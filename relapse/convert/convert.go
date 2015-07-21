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

func toTrans(refs relapse.RefLookup) (transitions transitions, finals map[string]*relapse.Pattern) {
	es := getAllExprs(refs)
	p := refs["main"]
	transitions = transForPattern(refs, es, p)
	visited := make(map[string]struct{})
	visited[p.String()] = struct{}{}
	finals = make(map[string]*relapse.Pattern)
	if interp.Nullable(refs, p) {
		finals[p.String()] = p
	}
	allVisited := false
	for !allVisited {
		allVisited = true
		for _, t := range transitions {
			sdst := t.dst.String()
			if _, ok := visited[sdst]; !ok {
				if interp.Nullable(refs, t.dst) {
					finals[sdst] = t.dst
				}
				tsdst := transForPattern(refs, es, t.dst)
				transitions.add(tsdst...)
				visited[sdst] = struct{}{}
				allVisited = false
			}
			schild := t.child.String()
			if _, ok := visited[schild]; !ok {
				if interp.Nullable(refs, t.child) {
					finals[schild] = t.child
				}
				tschild := transForPattern(refs, es, t.child)
				transitions.add(tschild...)
				visited[schild] = struct{}{}
				allVisited = false
			}
		}
	}
	return transitions, finals
}

func newCall(src *relapse.Pattern, expr *expr.Expr, parentDst *relapse.Pattern, childDst *relapse.Pattern) *viper.Rule {
	return viper.NewCall(
		strconv.Quote(src.String()),
		expr,
		strconv.Quote(parentDst.String()),
		strconv.Quote(childDst.String()),
	)
}

func newReturn(parentSrc *relapse.Pattern, childSrc *relapse.Pattern, expr *expr.Expr, dst *relapse.Pattern) *viper.Rule {
	return viper.NewReturn(
		strconv.Quote(parentSrc.String()),
		strconv.Quote(childSrc.String()),
		expr,
		strconv.Quote(dst.String()),
	)
}

func ToRules(g *relapse.Grammar) *viper.Rules {
	refs := relapse.NewRefsLookup(g)
	ts, fs := toTrans(refs)
	rules := []*viper.Rule{}
	for _, t := range ts {
		for _, f := range fs {
			rules = append(rules,
				newCall(t.src, t.expr, t.src, t.child),
				newReturn(t.src, f, t.expr, t.dst),
			)
		}
	}
	return viper.NewRules(rules...)
}

func transForPattern(refs relapse.RefLookup, es []*expr.Expr, p *relapse.Pattern) transitions {
	ts := make(transitions, 0, len(es))
	for i, e := range es {
		d := deriv(refs, p, e)
		ddst := interp.Simplify(refs, d.dst)
		dchild := interp.Simplify(refs, d.child)
		ts.add(&trans{
			src:   p,
			expr:  es[i],
			child: dchild,
			dst:   ddst,
		})
	}
	return ts
}

func getAllExprs(refs relapse.RefLookup) []*expr.Expr {
	es := []*expr.Expr{}
	for _, p := range refs {
		es = append(es, getExprs(p)...)
	}
	return es
}

func getExprs(p *relapse.Pattern) []*expr.Expr {
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return nil
	case *relapse.EmptySet:
		return nil
	case *relapse.TreeNode:
		e := nameToExpr(v.GetName())
		es := getExprs(v.GetPattern())
		return append([]*expr.Expr{e}, es...)
	case *relapse.LeafNode:
		return []*expr.Expr{v.GetExpr()}
	case *relapse.Concat:
		left := getExprs(v.GetLeftPattern())
		right := getExprs(v.GetRightPattern())
		return append(left, right...)
	case *relapse.Or:
		left := getExprs(v.GetLeftPattern())
		right := getExprs(v.GetRightPattern())
		return append(left, right...)
	case *relapse.And:
		left := getExprs(v.GetLeftPattern())
		right := getExprs(v.GetRightPattern())
		return append(left, right...)
	case *relapse.ZeroOrMore:
		return getExprs(v.GetPattern())
	case *relapse.Reference:
		return nil
	case *relapse.Not:
		return getExprs(v.GetPattern())
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

type transitions []*trans

func (as transitions) add(bs ...*trans) {
	for _, b := range bs {
		if as.contains(b) {
			continue
		}
		as = append(as, b)
	}
}

func (as transitions) contains(b *trans) bool {
	for _, a := range as {
		if a.Equal(b) {
			return true
		}
	}
	return false
}

type trans struct {
	src   *relapse.Pattern
	expr  *expr.Expr
	child *relapse.Pattern
	dst   *relapse.Pattern
}

func (a *trans) Equal(b *trans) bool {
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

type downup struct {
	child *relapse.Pattern
	dst   *relapse.Pattern
}

func concat(c *downup, p *relapse.Pattern) *downup {
	return &downup{
		child: c.child,
		dst:   relapse.NewConcat(c.dst, p),
	}
}

func or(left *downup, right *downup) *downup {
	return &downup{
		child: relapse.NewOr(left.child, right.child),
		dst:   relapse.NewOr(left.child, right.child),
	}
}

func and(left *downup, right *downup) *downup {
	return &downup{
		child: relapse.NewAnd(left.child, right.child),
		dst:   relapse.NewAnd(left.child, right.child),
	}
}

//TODO think about this again and write a few tests
func not(c *downup) *downup {
	return &downup{
		child: c.child,
		dst:   relapse.NewNot(c.dst),
	}
}

func emptySet() *downup {
	return &downup{
		child: relapse.NewEmptySet(),
		dst:   relapse.NewEmptySet(),
	}
}

func deriv(refs relapse.RefLookup, p *relapse.Pattern, expr *expr.Expr) *downup {
	d := func(pattern *relapse.Pattern) *downup {
		return deriv(refs, pattern, expr)
	}
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return emptySet()
	case *relapse.EmptySet:
		return emptySet()
	case *relapse.TreeNode:
		if !expr.Equal(nameToExpr(v.GetName())) {
			return emptySet()
		}
		return &downup{
			child: v.GetPattern(),
			dst:   relapse.NewEmpty(),
		}
	case *relapse.LeafNode:
		if !expr.Equal(v.GetExpr()) {
			return emptySet()
		}
		return &downup{
			child: relapse.NewEmpty(), //should not have any children
			dst:   relapse.NewEmpty(),
		}
	case *relapse.Concat:
		c := concat(d(v.GetLeftPattern()), v.GetRightPattern())
		if !interp.Nullable(refs, v.GetLeftPattern()) {
			return c
		}
		return or(c, d(v.GetRightPattern()))
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
