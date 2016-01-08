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
	c.convert(p)
	finals := c.getReachable(c.states[start])
	c.exhaust()
	for _, f := range finals {
		if interp.Nullable(c.refs, c.getPattern(c.states[f].current)) {
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
				this.convert(p)
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

func (this *converter) getReachable(state *state) []int {
	allDsts := this.states[state.current].tops()
	allDsts[state.current] = struct{}{}
	prevDsts := allDsts
	nextDsts := make(map[int]struct{})
	for {
		for _, d := range setToList(prevDsts) {
			if _, ok := this.states[d]; !ok {
				this.convert(this.getPattern(d))
			}
			if this.states[d] == nil {
				fmt.Printf("yeah: %v\n", this.getPattern(d))
				panic("yeah")
			}
			for n, _ := range this.states[d].tops() {
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

func (this *state) tops() map[int]struct{} {
	tops := make(map[int]struct{})
	for _, t := range this.trans {
		for _, u := range t.ups {
			tops[u.top] = struct{}{}
		}
	}
	return tops
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

func (this *converter) union(left, right *state) []tran {
	ts := []tran{}
	for i, _ := range left.trans {
		ts = append(ts, left.trans[i])
	}
	for i, _ := range right.trans {
		ts = append(ts, right.trans[i])
	}
	return ts
}

func (this *converter) intersect(left, right *state) []tran {
	ts := []tran{}
	for _, lt := range left.trans {
		for _, rt := range right.trans {
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

func (this *converter) leftConcat(left *state, rightCurrent int) []tran {
	ts := this.copy(left.trans)
	for i, lt := range left.trans {
		for j, lu := range lt.ups {
			ts[i].ups[j].top = this.addPattern(relapse.NewConcat(this.getPattern(lu.top), this.getPattern(rightCurrent)))
		}
	}
	return ts
}

func (this *converter) concat(current int, left, right *state) []tran {
	ts := this.leftConcat(left, right.current)
	if !interp.Nullable(this.refs, this.getPattern(left.current)) {
		return ts
	}
	newleft := &state{current, false, ts}
	fmt.Printf("concating %v\n", this.getPattern(current))
	fmt.Printf("newleft %v\n", this.toStr(newleft))
	fmt.Printf("right %v\n", this.toStr(right))
	trans := this.union(newleft, right)
	return trans
}

func (this *converter) interleave(left *state, right *relapse.Pattern) []tran {
	ts := this.copy(left.trans)
	for i, lt := range left.trans {
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

func (this *converter) convert(p *relapse.Pattern) *state {
	p = interp.Simplify(this.refs, p)
	c := this.addPattern(p)
	if this.states[c] != nil {
		return this.states[c]
	}
	this.states[c] = nil
	typ := p.GetValue()
	switch v := typ.(type) {
	case *relapse.Empty:
		s := this.newState(
			this.getEmpty(),
			[]tran{
				this.newDeadEnd(funcs.BoolConst(true)),
			},
		)
		this.states[s.current] = s
		return s
	case *relapse.TreeNode:
		current := this.addPattern(p)
		f := nameToFunc(v.GetName())
		below := this.convert(v.GetPattern())
		tops := this.getReachable(below)
		ups := make([]up, 0, len(tops))
		for _, d := range tops {
			if interp.Nullable(this.refs, this.getPattern(d)) {
				ups = append(ups, up{d, this.getEmpty()})
			} else {
				ups = append(ups, up{d, this.getNone()})
			}
		}
		s := this.newState(
			current,
			[]tran{
				{f, this.addPattern(v.Pattern), ups},
				this.newDeadEnd(not(f)),
			},
		)
		this.states[s.current] = s
		return s
	case *relapse.LeafNode:
		f := exprToFunc(v.GetExpr())
		s := this.newState(
			this.addPattern(p),
			[]tran{
				{
					f, this.getEmpty(), []up{
						{this.getEmpty(), this.getEmpty()},
						{this.getNone(), this.getNone()},
					},
				},
				this.newDeadEnd(not(f)),
			},
		)
		this.states[s.current] = s
		return s
	case *relapse.Concat:
		left := this.convert(v.GetLeftPattern())
		right := this.convert(v.GetRightPattern())
		current := this.addPattern(p)
		trans := this.concat(current, left, right)
		s := this.newState(current, trans)
		this.states[s.current] = s
		return s
	case *relapse.Or:
		left := this.convert(v.GetLeftPattern())
		right := this.convert(v.GetRightPattern())
		current := this.addPattern(p)
		ts := this.union(left, right)
		s := this.newState(current, ts)
		this.states[s.current] = s
		return s
	case *relapse.And:
		left := this.convert(v.GetLeftPattern())
		right := this.convert(v.GetRightPattern())
		current := this.addPattern(p)
		ts := this.intersect(left, right)
		s := this.newState(current, ts)
		this.states[s.current] = s
		return s
	case *relapse.ZeroOrMore:
		elem := this.convert(v.GetPattern())
		current := this.addPattern(p)
		ts := this.leftConcat(elem, current)
		// if interp.Nullable(this.refs, v.GetPattern()) {
		// 	panic("not implemented")
		// } else {
		s := this.newState(current, ts)
		this.states[s.current] = s
		return s
		//}
	case *relapse.Reference:

	case *relapse.Not:
		n := this.convert(v.GetPattern())
		trans := make([]tran, len(n.trans))
		for i := range n.trans {
			ups := []up{}
			for j := range n.trans[i].ups {
				d := n.trans[i].ups[j].top
				dpat := this.getPattern(d)
				ndpat := relapse.NewNot(dpat)
				ups = append(ups, up{
					top: this.addPattern(ndpat),
					bot: n.trans[i].ups[j].bot,
				})
			}
			trans[i] = tran{
				value: n.trans[i].value,
				down:  n.trans[i].down,
				ups:   ups,
			}
		}
		s := this.newState(this.addPattern(p), trans)
		this.states[s.current] = s
		return s
	case *relapse.ZAny:
		s := this.newState(
			this.getAny(),
			[]tran{
				this.newAnyEnd(funcs.BoolConst(true)),
			},
		)
		this.states[s.current] = s
		return s
	case *relapse.Contains:
		left := this.convert(relapse.NewZAny())
		right := this.convert(relapse.NewConcat(v.GetPattern(), relapse.NewZAny()))
		current := this.addPattern(p)
		trans := this.concat(current, left, right)
		s := this.newState(current, trans)
		this.states[s.current] = s
		return s
	case *relapse.Optional:
		left := this.convert(v.GetPattern())
		right := this.convert(relapse.NewEmpty())
		current := this.addPattern(p)
		ts := this.union(left, right)
		s := this.newState(current, ts)
		this.states[s.current] = s
		return s
	case *relapse.Interleave:
		left := this.convert(v.GetLeftPattern())
		right := this.convert(v.GetRightPattern())
		current := this.addPattern(p)
		//leftfirst := this.addPattern(relapse.NewConcat(v.GetLeftPattern(), v.GetRightPattern()))
		//rightfirst := this.addPattern(relapse.NewConcat(v.GetRightPattern(), v.GetLeftPattern()))
		lefttrans := this.interleave(left, v.GetRightPattern())
		righttrans := this.interleave(right, v.GetLeftPattern())
		ts := this.union(&state{current, false, lefttrans}, &state{current, false, righttrans})
		s := this.newState(current, ts)
		this.states[s.current] = s
		return s
	}
	panic(fmt.Sprintf("unknown pattern typ %T %v", typ, p))
}
