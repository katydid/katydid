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
	"github.com/katydid/katydid/parser"
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/compose"
	"github.com/katydid/katydid/relapse/funcs"
)

type ifExprs struct {
	cond     funcs.Bool
	composed compose.Bool
	then     *ifExprs
	els      *ifExprs
	ret      []*ast.Pattern
}

//compileIfExprs combines several if expressions into one nested if expression with a list of return values.
//While combining these if expressions, duplicate and impossible (always false) conditions are removed for efficiency.
func compileIfExprs(ifs []*ifExpr) *ifExprs {
	if len(ifs) == 0 {
		return &ifExprs{
			ret: []*ast.Pattern{},
		}
	}
	root := &ifExprs{}
	if ifs[0].els == nil || ifs[0].then.Equal(ifs[0].els) {
		root.ret = []*ast.Pattern{ifs[0].then}
	} else {
		root.cond = ifs[0].cond
		root.then = &ifExprs{ret: []*ast.Pattern{ifs[0].then}}
		root.els = &ifExprs{ret: []*ast.Pattern{ifs[0].els}}
	}
	for _, ifexpr := range ifs[1:] {
		if ifexpr.cond == nil {
			root.addReturn(ifexpr.then)
		} else {
			root.addIfExpr(ifexpr.cond, ifexpr.then, ifexpr.els)
		}
	}
	return root
}

func (this *ifExprs) eval(label parser.Value) ([]*ast.Pattern, error) {
	if this.ret != nil {
		return this.ret, nil
	}
	if this.composed == nil {
		composed, err := compose.NewBoolFunc(this.cond)
		if err != nil {
			return nil, err
		}
		this.composed = composed
	}
	cond, err := this.composed.Eval(label)
	if err != nil {
		return nil, err
	}
	if cond {
		return this.then.eval(label)
	}
	return this.els.eval(label)
}

//addReturn finds the leafs and appends a return to each.
func (this *ifExprs) addReturn(ret *ast.Pattern) {
	if this.ret != nil {
		this.ret = append(this.ret, ret)
		return
	}
	this.then.addReturn(ret)
	this.els.addReturn(ret)
	return
}

func (this *ifExprs) addIfExpr(cond funcs.Bool, then, els *ast.Pattern) {
	// efficienctly append the then and else return to two copies of the current returns.
	if this.ret != nil {
		this.cond = cond
		thenterms := make([]*ast.Pattern, len(this.ret)+1)
		copy(thenterms, this.ret)
		thenterms[len(thenterms)-1] = then
		this.then = &ifExprs{ret: thenterms}
		this.els = &ifExprs{ret: append(this.ret, els)}
		this.ret = nil
		return
	}
	// remove duplicate condition
	if funcs.Equal(this.cond, cond) {
		this.then.addReturn(then)
		this.els.addReturn(els)
		return
	}
	// remove impossible (always false) then condition
	if funcs.IsFalse(funcs.Simplify(funcs.And(this.cond, cond))) {
		this.then.addReturn(els)
		this.els.addIfExpr(cond, then, els)
		return
	}
	// remove impossible (always false) else condition
	if funcs.IsFalse(funcs.Simplify(funcs.And(this.cond, funcs.Not(cond)))) {
		this.then.addIfExpr(cond, then, els)
		this.els.addReturn(then)
		return
	}
	this.then.addIfExpr(cond, then, els)
	this.els.addIfExpr(cond, then, els)
	return
}

type ifExpr struct {
	cond funcs.Bool
	then *ast.Pattern
	els  *ast.Pattern
}

func (this *ifExpr) eval(label parser.Value) (*ast.Pattern, error) {
	if this.els == nil {
		return this.then, nil
	}
	f, err := compose.NewBoolFunc(this.cond)
	if err != nil {
		return nil, err
	}
	cond, err := f.Eval(label)
	if err != nil {
		return nil, err
	}
	if cond {
		return this.then, nil
	}
	return this.els, nil
}
