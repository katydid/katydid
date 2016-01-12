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

package convert

import (
	"fmt"
	"github.com/katydid/katydid/expr/ast"
	"github.com/katydid/katydid/expr/compose"
	"github.com/katydid/katydid/funcs"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/interp"
	"github.com/katydid/katydid/relapse/nameexpr"
	"sort"
	"strconv"
	"strings"
)

type auto struct {
	start    int
	states   map[int]*state
	patterns []*relapse.Pattern
}

func (this *auto) draw() {
	s := []string{}
	a := func(line string) {
		s = append(s, line)
	}
	a("digraph {")
	a("\trankdir=LR;")
	a("\tsize=\"8,5\"")
	accepts := []string{}
	other := []string{}
	for _, s := range this.states {
		current := this.patterns[s.current].String()
		scurrent := strconv.Quote(current)
		if s.final {
			accepts = append(accepts, scurrent)
		} else {
			other = append(other, scurrent)
		}
	}
	sort.Strings(accepts)
	sort.Strings(other)
	a("\tnode [shape = doublecircle]; " + strings.Join(accepts, " "))
	a("\tnode [shape = circle]; " + strings.Join(other, " "))
	for si, _ := range this.patterns {
		s := this.states[si]
		current := this.patterns[s.current].String()
		scurrent := strconv.Quote(current)
		if this.start == s.current {
			a("\tstart -> " + scurrent)
		}
		for _, t := range s.trans {
			v := funcs.Sprint(t.value)
			a("\t" + scurrent + " -> " + strconv.Quote(this.patterns[t.down].String()) + " [ label = " + strconv.Quote(v+"&uarr;"+current) + " ];")
			for _, u := range t.ups {
				a("\t" + strconv.Quote(this.patterns[u.bot].String()) + " -> " + strconv.Quote(this.patterns[u.top].String()) + " [ label = " + strconv.Quote(v+"&darr;"+current) + " ];")
			}
		}
	}
	a("}")
	fmt.Printf("%s\n", strings.Join(s, "\n"))
}

func Convert(refs relapse.RefLookup, p *relapse.Pattern) *auto {
	c := newConverter(refs)
	start := c.addPattern(p)
	c.sconvert(p)
	c.exhaust()
	finals := c.getReachable(p)
	for _, f := range finals {
		currentpattern := c.states[f].current
		if interp.Nullable(c.refs, c.getPattern(currentpattern)) {
			c.states[f].final = true
		}
	}
	a := &auto{start, c.states, c.patterns}
	a.draw()
	return a
}

type converter struct {
	states   map[int]*state
	patterns []*relapse.Pattern
	refs     relapse.RefLookup
}

func newConverter(refs relapse.RefLookup) *converter {
	return &converter{
		states:   make(map[int]*state),
		patterns: []*relapse.Pattern{relapse.NewZAny(), relapse.NewNot(relapse.NewZAny()), relapse.NewEmpty()},
		refs:     refs,
	}
}

func (this *converter) addPattern(p *relapse.Pattern) int {
	p = interp.Simplify(this.refs, p)
	for i, pat := range this.patterns {
		if pat.Equal(p) {
			return i
		}
	}
	this.patterns = append(this.patterns, p)
	return len(this.patterns) - 1
}

func (this *converter) getPattern(i int) *relapse.Pattern {
	return this.patterns[i]
}

func (this *converter) getAny() int {
	return 0
}

func (this *converter) getNone() int {
	return 1
}

func (this *converter) getEmpty() int {
	return 2
}

func (this *converter) exhaust() {
	changed := true
	for changed {
		changed = false
		for i, p := range this.patterns {
			if _, ok := this.states[i]; !ok {
				fmt.Printf("exhausting %v %d\n", p, len(this.states))
				this.sconvert(p)
				changed = true
			}
		}
	}
}

func setToList(ms map[int]struct{}) []int {
	is := make([]int, 0, len(ms))
	for i, _ := range ms {
		is = append(is, i)
	}
	sort.Ints(is)
	return is
}

func tops(trans []tran) map[int]struct{} {
	tops := make(map[int]struct{})
	for _, t := range trans {
		for _, u := range t.ups {
			tops[u.top] = struct{}{}
		}
	}
	return tops
}

type Set map[int]struct{}

func newSet() Set {
	return make(Set)
}

func (this Set) add(i int) Set {
	this[i] = struct{}{}
	return this
}

func (this Set) union(that Set) Set {
	for j := range that {
		this = this.add(j)
	}
	return this
}

func (this *converter) reachableConcat(left, right Set) Set {
	s := newSet()
	for l := range left {
		for r := range right {
			lp := this.getPattern(l)
			rp := this.getPattern(r)
			c := relapse.NewConcat(lp, rp)
			s.add(this.addPattern(c))
		}
	}
	return s
}

func (this *converter) reachableIntersect(left, right Set) Set {
	s := newSet()
	for l := range left {
		for r := range right {
			lp := this.getPattern(l)
			rp := this.getPattern(r)
			c := relapse.NewAnd(lp, rp)
			s.add(this.addPattern(c))
		}
	}
	return s
}

func (this *converter) reachableInterleave(left, right Set) Set {
	s := newSet()
	for l := range left {
		for r := range right {
			lp := this.getPattern(l)
			rp := this.getPattern(r)
			c := relapse.NewInterleave(lp, rp)
			s.add(this.addPattern(c))
		}
	}
	return s
}

func (this *converter) getReachable2(p *relapse.Pattern) Set {
	fmt.Printf("getReachable2 %s\n", p.String())
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return newSet().add(this.getNone())
	case *relapse.TreeNode:
		return newSet().add(this.getEmpty()).add(this.getNone())
	case *relapse.LeafNode:
		return newSet().add(this.getEmpty()).add(this.getNone())
	case *relapse.Concat:
		lefts1 := this.getReachable2(v.GetLeftPattern())
		rights := this.getReachable2(v.GetRightPattern())
		lefts2 := this.reachableConcat(lefts1, rights)
		if !interp.Nullable(this.refs, v.GetLeftPattern()) {
			return lefts2
		}
		return lefts2.union(rights)
	case *relapse.Or:
		lefts := this.getReachable2(v.GetLeftPattern())
		rights := this.getReachable2(v.GetRightPattern())
		return lefts.union(rights)
	case *relapse.And:
		left := this.getReachable2(v.GetLeftPattern())
		right := this.getReachable2(v.GetRightPattern())
		return this.reachableIntersect(left, right)
	case *relapse.ZeroOrMore:
		pat := this.getReachable2(v.GetPattern())
		here := newSet().add(this.addPattern(p))
		return this.reachableConcat(pat, here)
	case *relapse.Reference:
		r := this.refs[v.GetName()]
		return this.getReachable2(r)
	case *relapse.Not:
		rs := this.getReachable2(v.GetPattern())
		set := newSet()
		for r := range rs {
			set.add(this.addPattern(relapse.NewNot(this.getPattern(r))))
		}
		return set
	case *relapse.ZAny:
		return newSet().add(this.getAny())
	case *relapse.Contains:
		return this.getReachable2(relapse.NewConcat(relapse.NewZAny(), relapse.NewConcat(v.GetPattern(), relapse.NewZAny())))
	case *relapse.Optional:
		return this.getReachable2(relapse.NewOr(v.GetPattern(), relapse.NewEmpty()))
	case *relapse.Interleave:
		left := this.getReachable2(v.GetLeftPattern())
		right := this.getReachable2(v.GetRightPattern())
		lefts := this.reachableInterleave(left, newSet().add(this.addPattern(v.GetRightPattern())))
		rights := this.reachableInterleave(right, newSet().add(this.addPattern(v.GetLeftPattern())))
		return lefts.union(rights)
	}
	panic(fmt.Sprintf("unknown pattern typ %T %v", typ, p))
}

func (this *converter) getTops(p *relapse.Pattern) Set {
	d := this.addPattern(p)
	sd := this.states[d]
	if sd != nil {
		return tops(sd.trans)
	}
	return this.getReachable2(p)
}

func (this *converter) getReachable(p *relapse.Pattern) []int {
	current := this.addPattern(p)
	allDsts := this.getTops(p)
	allDsts[current] = struct{}{}
	prevDsts := allDsts
	nextDsts := make(map[int]struct{})
	for {
		for _, d := range setToList(prevDsts) {
			dsts := this.getTops(this.getPattern(d))
			for n, _ := range dsts {
				if _, ok := allDsts[n]; !ok {
					nextDsts[n] = struct{}{}
					allDsts[n] = struct{}{}
				}
			}
		}
		if len(nextDsts) == 0 {
			break
		}
		prevDsts = make(map[int]struct{})
		for n, _ := range nextDsts {
			prevDsts[n] = struct{}{}
		}
		nextDsts = make(map[int]struct{})
	}
	return setToList(allDsts)
}

type state struct {
	current int
	final   bool
	trans   []tran
}

type tran struct {
	value funcs.Bool
	down  int
	ups   []up
}

func (this tran) Equal(that tran) bool {
	if !funcs.Equal(this.value, that.value) {
		return false
	}
	if this.down != that.down {
		return false
	}
	if len(this.ups) != len(that.ups) {
		return false
	}
	for i := range this.ups {
		if this.ups[i].bot != that.ups[i].bot {
			return false
		}
		if this.ups[i].top != that.ups[i].top {
			return false
		}
	}
	return true
}

type up struct {
	bot int
	top int
}

func (this *converter) toStr(s *state) string {
	current := this.getPattern(s.current)
	ts := []string{}
	for _, t := range s.trans {
		us := []string{}
		for _, u := range t.ups {
			us = append(us, "( "+this.getPattern(u.bot).String()+" ^ "+this.getPattern(u.top).String()+" )")
		}
		ts = append(ts, funcs.Sprint(t.value)+" -> "+this.getPattern(t.down).String()+" [ "+strings.Join(us, ", ")+" ]")
	}
	return current.String() + " { " + strings.Join(ts, ", ") + " }"
}

func (this *converter) newState(current int, trans []tran) *state {
	// fmt.Printf("deduping %v\n", this.toStr(&state{current, false, trans}))
	// for i := range trans {
	// 	trans[i].ups = this.dedup2(trans[i].ups)
	// }
	//mtrans := make(map[string]int, len(trans))
	is := make([]int, 0, len(trans))
	for trani := range trans {
		sf := funcs.Sprint(trans[trani].value)
		if sf == "false" {
			continue
		}
		found := false
		for _, i := range is {
			if trans[i].Equal(trans[trani]) {
				found = true
			}
		}
		if found {
			continue
		}
		// 	if j, ok := mtrans[sf]; ok {
		// 		if !trans[i].Equal(trans[j]) {
		// 			fmt.Printf("deduped! %v\n", this.toStr(&state{current, false, trans}))
		// 			panic("wtf")
		// 		}
		// 	}
		// 	mtrans[funcs.Sprint(trans[i].value)] = i
		is = append(is, trani)
	}
	trans2 := make([]tran, 0, len(is))
	for _, i := range is {
		trans2 = append(trans2, trans[i])
	}
	s := &state{
		current: current,
		trans:   trans2,
	}
	this.states[current] = s
	// fmt.Printf("deduped  %v\n", this.toStr(s))
	return s
}

func (this *converter) dedup2(ups []up) []up {
	rups := make(map[int]int)
	for _, up := range ups {
		if d, ok := rups[up.bot]; !ok {
			rups[up.bot] = up.top
		} else if d != up.top {
			fmt.Printf("wtf\n")
			rups[up.bot] = this.addPattern(relapse.NewOr(this.getPattern(up.top), this.getPattern(d)))
		}
	}
	iups := []int{}
	for i := range rups {
		iups = append(iups, i)
	}
	sort.Ints(iups)
	nups := []up{}
	for _, i := range iups {
		nups = append(nups, up{i, rups[i]})
	}
	return nups
}

func (this *converter) dedup(ups []up) []up {
	rups := make(map[int]map[int]struct{})
	nups := []up{}
	for i, up := range ups {
		if _, ok := rups[up.bot]; !ok {
			nups = append(nups, ups[i])
			rups[up.bot] = make(map[int]struct{})
			rups[up.bot][up.top] = struct{}{}
		} else if _, ok1 := rups[up.bot][up.top]; !ok1 {
			nups = append(nups, ups[i])
			rups[up.bot][up.top] = struct{}{}
			fmt.Printf("wtf %v -> %v \n", this.patterns[up.bot], this.patterns[up.top])
			//panic("wtf different returns")
		}
	}
	return ups
}

func (this *converter) union(left, right []tran) []tran {
	ts := []tran{}
	for i, _ := range left {
		ts = append(ts, left[i])
	}
	for i, _ := range right {
		ts = append(ts, right[i])
	}
	return ts
}

func (this *converter) intersect(left, right []tran) []tran {
	ts := []tran{}
	for _, lt := range left {
		for _, rt := range right {
			pdl := this.getPattern(lt.down)
			pdr := this.getPattern(rt.down)
			ups := []up{}
			for _, lu := range lt.ups {
				for _, ru := range rt.ups {
					ups = append(ups, up{
						bot: this.addPattern(relapse.NewAnd(this.getPattern(lu.bot), this.getPattern(ru.bot))),
						top: this.addPattern(relapse.NewAnd(this.getPattern(lu.top), this.getPattern(ru.top))),
					})
				}
			}
			tt := tran{
				value: funcs.Simplify(funcs.And(lt.value, rt.value)),
				down:  this.addPattern(relapse.NewAnd(pdl, pdr)),
				ups:   ups,
			}
			ts = append(ts, tt)
		}
	}
	return ts
}

func (this *converter) copy(trans []tran) []tran {
	ts := []tran{}
	for _, t := range trans {
		ups := []up{}
		for _, u := range t.ups {
			ups = append(ups, up{
				bot: u.bot,
				top: u.top,
			})
		}
		tt := tran{
			value: t.value,
			down:  t.down,
			ups:   ups,
		}
		ts = append(ts, tt)
	}
	return ts
}

func (this *converter) leftConcat(left []tran, rightCurrent int) []tran {
	ts := this.copy(left)
	for i, lt := range left {
		for j, lu := range lt.ups {
			ts[i].ups[j].top = this.addPattern(relapse.NewConcat(this.getPattern(lu.top), this.getPattern(rightCurrent)))
		}
	}
	return ts
}

func (this *converter) interleave(left []tran, right *relapse.Pattern) []tran {
	ts := this.copy(left)
	for i, lt := range left {
		for j, lu := range lt.ups {
			ts[i].ups[j].top = this.addPattern(relapse.NewInterleave(this.getPattern(lu.top), right))
		}
	}
	return ts
}

func (this *converter) newDeadEnd(f funcs.Bool) tran {
	return tran{f, this.getNone(), []up{{this.getNone(), this.getNone()}}}
}

func (this *converter) newAnyEnd(f funcs.Bool) tran {
	return tran{f, this.getAny(), []up{{this.getAny(), this.getAny()}}}
}

func nameToFunc(name *relapse.NameExpr) funcs.Bool {
	return funcs.Simplify(nameexpr.NameToFunc(name))
}

func not(f funcs.Bool) funcs.Bool {
	return funcs.Simplify(funcs.Not(f))
}

func exprToFunc(expr *expr.Expr) funcs.Bool {
	f, err := compose.NewBool(expr)
	if err != nil {
		panic(err)
	}
	return funcs.Simplify(f)
}

func (this *converter) sconvert(p *relapse.Pattern) *state {
	fmt.Printf("sconvert %v\n", p)
	p = interp.Simplify(this.refs, p)
	current := this.addPattern(p)
	if this.states[current] != nil {
		return this.states[current]
	}
	trans := this.convert(p)
	return this.newState(current, trans)
}

func (this *converter) convert(p *relapse.Pattern) []tran {
	fmt.Printf("converting %s\n", p.String())
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return []tran{
			this.newDeadEnd(funcs.BoolConst(true)),
		}
	case *relapse.TreeNode:
		f := nameToFunc(v.GetName())
		tops := this.getReachable(v.GetPattern())
		ups := make([]up, 0, len(tops))
		for _, d := range tops {
			if interp.Nullable(this.refs, this.getPattern(d)) {
				ups = append(ups, up{d, this.getEmpty()})
			} else {
				ups = append(ups, up{d, this.getNone()})
			}
		}
		return []tran{
			{f, this.addPattern(v.Pattern), ups},
			this.newDeadEnd(not(f)),
		}
	case *relapse.LeafNode:
		f := exprToFunc(v.GetExpr())
		return []tran{
			{
				f, this.getEmpty(), []up{
					{this.getEmpty(), this.getEmpty()},
					{this.getNone(), this.getNone()},
				},
			},
			this.newDeadEnd(not(f)),
		}
	case *relapse.Concat:
		fmt.Printf("concating %v\n", p)
		left := this.convert(v.GetLeftPattern())
		leftts := this.leftConcat(left, this.addPattern(v.GetRightPattern()))
		if !interp.Nullable(this.refs, v.GetLeftPattern()) {
			return leftts
		}
		rightts := this.convert(v.GetRightPattern())
		trans := this.union(leftts, rightts)
		return trans
	case *relapse.Or:
		left := this.convert(v.GetLeftPattern())
		right := this.convert(v.GetRightPattern())
		ts := this.union(left, right)
		return ts
	case *relapse.And:
		left := this.convert(v.GetLeftPattern())
		right := this.convert(v.GetRightPattern())
		ts := this.intersect(left, right)
		return ts
	case *relapse.ZeroOrMore:
		ts := this.convert(v.GetPattern())
		current := this.addPattern(p)
		trans := this.leftConcat(ts, current)
		return trans
	case *relapse.Reference:
		p := this.refs[v.GetName()]
		trans := this.convert(p)
		return trans
	case *relapse.Not:
		ntrans := this.convert(v.GetPattern())
		trans := make([]tran, len(ntrans))
		for i := range ntrans {
			ups := []up{}
			for j := range ntrans[i].ups {
				d := ntrans[i].ups[j].top
				dpat := this.getPattern(d)
				ndpat := relapse.NewNot(dpat)
				ups = append(ups, up{
					top: this.addPattern(ndpat),
					bot: ntrans[i].ups[j].bot,
				})
			}
			trans[i] = tran{
				value: ntrans[i].value,
				down:  ntrans[i].down,
				ups:   ups,
			}
		}
		return trans
	case *relapse.ZAny:
		return []tran{
			this.newAnyEnd(funcs.BoolConst(true)),
		}
	case *relapse.Contains:
		return this.convert(relapse.NewConcat(relapse.NewZAny(), relapse.NewConcat(v.GetPattern(), relapse.NewZAny())))
	case *relapse.Optional:
		return this.convert(relapse.NewOr(v.GetPattern(), relapse.NewEmpty()))
	case *relapse.Interleave:
		left := this.convert(v.GetLeftPattern())
		right := this.convert(v.GetRightPattern())
		lefttrans := this.interleave(left, v.GetRightPattern())
		righttrans := this.interleave(right, v.GetLeftPattern())
		ts := this.union(lefttrans, righttrans)
		return ts
	}
	panic(fmt.Sprintf("unknown pattern typ %T %v", typ, p))
}
