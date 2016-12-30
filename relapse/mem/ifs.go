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
	cond funcs.Bool
	then *ifExprs
	els  *ifExprs
	ret  []*ast.Pattern
}

func newIfExprs(ifs []*ifExpr) *ifExprs {
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
	f, err := compose.NewBoolFunc(this.cond)
	if err != nil {
		return nil, err
	}
	cond, err := f.Eval(label)
	if err != nil {
		return nil, err
	}
	if cond {
		return this.then.eval(label)
	}
	return this.els.eval(label)
}

func (top *ifExprs) addReturn(ret *ast.Pattern) {
	if top.ret != nil {
		top.ret = append(top.ret, ret)
		return
	}
	top.then.addReturn(ret)
	top.els.addReturn(ret)
	return
}

func (top *ifExprs) addIfExpr(cond funcs.Bool, then, els *ast.Pattern) {
	if top.ret != nil {
		top.cond = cond
		thenterms := make([]*ast.Pattern, len(top.ret)+1)
		copy(thenterms, top.ret)
		thenterms[len(thenterms)-1] = then
		top.then = &ifExprs{ret: thenterms}
		top.els = &ifExprs{ret: append(top.ret, els)}
		top.ret = nil
		return
	}
	if funcs.Equal(top.cond, cond) {
		top.then.addReturn(then)
		top.els.addReturn(els)
		return
	}
	if funcs.IsFalse(funcs.Simplify(funcs.And(top.cond, cond))) {
		top.then.addReturn(els)
		top.els.addIfExpr(cond, then, els)
		return
	}
	if funcs.IsFalse(funcs.Simplify(funcs.And(top.cond, funcs.Not(cond)))) {
		top.then.addIfExpr(cond, then, els)
		top.els.addReturn(then)
		return
	}
	top.then.addIfExpr(cond, then, els)
	top.els.addIfExpr(cond, then, els)
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
