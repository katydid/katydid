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

func (this *auto) reachable(s int) bool {
	return this.reachable2(make(map[int]bool), s)
}

func (this *auto) reachable2(visited map[int]bool, s int) bool {
	if s == this.start {
		return true
	}
	if v, ok := visited[s]; ok {
		return v
	}
	visited[s] = false
	for _, state := range this.states {
		for _, tran := range state.trans {
			if tran.down == s {
				if this.reachable2(visited, state.current) {
					visited[s] = true
					return true
				}
			}
			for _, up := range tran.ups {
				if up.top == s {
					if this.reachable2(visited, state.current) && this.reachable2(visited, up.bot) {
						visited[s] = true
						return true
					}
				}
			}
		}
	}
	return false
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
		if !this.reachable(s.current) {
			continue
		}
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
		if !this.reachable(s.current) {
			continue
		}
		current := this.patterns[s.current].String()
		scurrent := strconv.Quote(current)
		if this.start == s.current {
			a("\tstart -> " + scurrent)
		}
		for _, t := range s.trans {
			v := funcs.Sprint(t.value)
			a("\t" + scurrent + " -> " + strconv.Quote(this.patterns[t.down].String()) + " [ label = " + strconv.Quote(v+"&uarr;"+current) + " ];")
			for _, u := range t.ups {
				if !this.reachable(u.bot) {
					continue
				}
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
		patterns: []*relapse.Pattern{},
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
	return this.addPattern(relapse.NewZAny())
}

func (this *converter) getNone() int {
	return this.addPattern(relapse.NewNot(relapse.NewZAny()))
}

func (this *converter) getEmpty() int {
	return this.addPattern(relapse.NewEmpty())
}

var dontcare = relapse.NewTreeNode(relapse.NewStringName("DontCare"), relapse.NewEmpty())

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

func (this Set) list() []int {
	is := make([]int, 0, len(this))
	for i, _ := range this {
		is = append(is, i)
	}
	sort.Ints(is)
	return is
}

func (this *converter) getTops(p *relapse.Pattern) Set {
	d := this.addPattern(p)
	sd := this.states[d]
	if sd != nil {
		return tops(sd.trans)
	}
	ts := this.convert(p, false)
	return tops(ts)
}

func (this *converter) getReachable(p *relapse.Pattern) []int {
	current := this.addPattern(p)
	allDsts := this.getTops(p).add(current)
	prevDsts := allDsts
	nextDsts := newSet()
	for {
		for _, d := range prevDsts.list() {
			dsts := this.getTops(this.getPattern(d))
			for _, n := range dsts.list() {
				fmt.Printf("getReachable from %v -> %d %v\n", p, n, this.getPattern(n))
				if _, ok := allDsts[n]; !ok {
					nextDsts = nextDsts.add(n)
					allDsts = allDsts.add(n)
				}
			}
		}
		if len(nextDsts) == 0 {
			break
		}
		prevDsts = newSet().union(nextDsts)
		nextDsts = newSet()
	}
	return allDsts.list()
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

func toStr(s *state, ps []*relapse.Pattern) string {
	current := ps[s.current]
	ts := []string{}
	for _, t := range s.trans {
		us := []string{}
		for _, u := range t.ups {
			us = append(us, "\t\t"+ps[u.bot].String()+" ^ "+ps[u.top].String())
		}
		ts = append(ts, "\t"+funcs.Sprint(t.value)+" -> "+ps[t.down].String()+" [\n"+strings.Join(us, ",\n")+"\n\t]")
	}
	return current.String() + " {\n" + strings.Join(ts, ",\n") + "\n}"
}

func (this *converter) toStr(s *state) string {
	return toStr(s, this.patterns)
}

func (this *converter) newState(current int, trans []tran) *state {
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
	return s
}

func (this *converter) orups(as, bs []up) []up {
	newups := make(map[int]int)
	newbots := []int{}
	for _, a := range as {
		for _, b := range bs {
			newbot := this.addPattern(relapse.NewOr(this.getPattern(a.bot), this.getPattern(b.bot)))
			newtop := this.addPattern(relapse.NewOr(this.getPattern(a.top), this.getPattern(b.top)))
			oldtop, ok := newups[newbot]
			if ok {
				newups[newbot] = this.addPattern(relapse.NewOr(this.getPattern(oldtop), this.getPattern(newtop)))
			} else {
				newbots = append(newbots, newbot)
				newups[newbot] = newtop
			}
		}
	}
	sort.Ints(newbots)
	ups := make([]up, 0, len(newbots))
	for i := range newbots {
		ups = append(ups, up{bot: newbots[i], top: newups[newbots[i]]})
	}
	return ups
}

func (this *converter) addion(ts []tran, t tran) []tran {
	tvaluestr := funcs.Sprint(t.value)
	if tvaluestr == "false" {
		return ts
	}
	for i := range ts {
		if funcs.Sprint(ts[i].value) == tvaluestr {
			pdts := this.getPattern(ts[i].down)
			pdt := this.getPattern(t.down)
			ups := this.orups(t.ups, ts[i].ups)
			ts[i] = tran{
				value: ts[i].value,
				down:  this.addPattern(relapse.NewOr(pdts, pdt)),
				ups:   ups,
			}
			return ts
		}
	}
	return append(ts, t)
}

func (this *converter) union(left, right []tran) []tran {
	ts := []tran{}
	for _, lt := range left {
		for _, rt := range right {
			pdl := this.getPattern(lt.down)
			pdr := this.getPattern(rt.down)
			ups := this.orups(lt.ups, rt.ups)

			truetrue := funcs.Simplify(funcs.And(lt.value, rt.value))
			ts = this.addion(ts, tran{
				value: truetrue,
				down:  this.addPattern(relapse.NewOr(pdl, pdr)),
				ups:   ups,
			})

			// truefalse := funcs.Simplify(funcs.And(lt.value, not(rt.value)))
			// ts = this.addion(ts, tran{
			// 	value: truefalse,
			// 	down:  lt.down,
			// 	ups:   lt.ups,
			// })

			// falsetrue := funcs.Simplify(funcs.And(not(lt.value), rt.value))
			// ts = this.addion(ts, tran{
			// 	value: falsetrue,
			// 	down:  rt.down,
			// 	ups:   rt.ups,
			// })

		}
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
	p = interp.Simplify(this.refs, p)
	fmt.Printf("sconvert %v\n", p)
	current := this.addPattern(p)
	if this.states[current] != nil {
		return this.states[current]
	}
	trans := this.convert(p, true)
	return this.newState(current, trans)
}

func (this *converter) convert(p *relapse.Pattern, down bool) []tran {
	fmt.Printf("converting %s\n", p.String())
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		return []tran{
			this.newDeadEnd(funcs.BoolConst(true)),
		}
	case *relapse.TreeNode:
		f := nameToFunc(v.GetName())
		if !down {
			return []tran{
				{f, this.addPattern(v.Pattern), []up{
					up{this.getEmpty(), this.getEmpty()},
					up{this.getNone(), this.getNone()},
				}},
				this.newDeadEnd(not(f)),
			}
		}
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
		left := this.convert(v.GetLeftPattern(), down)
		leftts := this.leftConcat(left, this.addPattern(v.GetRightPattern()))
		if !interp.Nullable(this.refs, v.GetLeftPattern()) {
			return leftts
		}
		rightts := this.convert(v.GetRightPattern(), down)
		trans := this.union(leftts, rightts)
		return trans
	case *relapse.Or:
		left := this.convert(v.GetLeftPattern(), down)
		right := this.convert(v.GetRightPattern(), down)
		fmt.Printf("ORLEFT = %v\n", this.toStr(&state{this.addPattern(v.GetLeftPattern()), false, left}))
		fmt.Printf("ORRIGHT = %v\n", this.toStr(&state{this.addPattern(v.GetRightPattern()), false, right}))
		ts := this.union(left, right)
		fmt.Printf("ORRES = %v\n", this.toStr(&state{this.addPattern(p), false, ts}))
		return ts
	case *relapse.And:
		left := this.convert(v.GetLeftPattern(), down)
		right := this.convert(v.GetRightPattern(), down)
		ts := this.intersect(left, right)
		return ts
	case *relapse.ZeroOrMore:
		ts := this.convert(v.GetPattern(), down)
		current := this.addPattern(p)
		trans := this.leftConcat(ts, current)
		return trans
	case *relapse.Reference:
		p := this.refs[v.GetName()]
		trans := this.convert(p, down)
		return trans
	case *relapse.Not:
		ntrans := this.convert(v.GetPattern(), down)
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
		return this.convert(relapse.NewConcat(relapse.NewZAny(), relapse.NewConcat(v.GetPattern(), relapse.NewZAny())), down)
	case *relapse.Optional:
		return this.convert(relapse.NewOr(v.GetPattern(), relapse.NewEmpty()), down)
	case *relapse.Interleave:
		left := this.convert(v.GetLeftPattern(), down)
		right := this.convert(v.GetRightPattern(), down)
		lefttrans := this.interleave(left, v.GetRightPattern())
		righttrans := this.interleave(right, v.GetLeftPattern())
		ts := this.union(lefttrans, righttrans)
		return ts
	}
	panic(fmt.Sprintf("unknown pattern typ %T %v", typ, p))
}
