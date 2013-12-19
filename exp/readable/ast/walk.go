//  Copyright 2013 Walter Schulze
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

package ast

type Visitor interface {
	Visit(node interface{}) Visitor
}

func (this *Rules) Walk(v Visitor) {
	v = v.Visit(this)
	if this.Root != nil {
		this.GetRoot().Walk(v)
	}
	for _, init := range this.GetInit() {
		init.Walk(v)
	}
	for _, trans := range this.GetTransition() {
		trans.Walk(v)
	}
	for _, ifexpr := range this.GetIfExpr() {
		ifexpr.Walk(v)
	}
}

func (this *Init) Walk(v Visitor) {
	v.Visit(this)
}

func (this *Transition) Walk(v Visitor) {
	v.Visit(this)
}

func (this *IfExpr) Walk(v Visitor) {
	v = v.Visit(this)
	this.GetCondition().Walk(v)
	this.GetThen().Walk(v)
	this.GetElse().Walk(v)
}

func (this *StateExpr) Walk(v Visitor) {
	v = v.Visit(this)
	if this.IfExpr != nil {
		this.GetIfExpr().Walk(v)
	}
}

func (this *Expr) Walk(v Visitor) {
	v = v.Visit(this)
	if this.Terminal != nil {
		this.GetTerminal().Walk(v)
	}
	if this.Function != nil {
		this.GetFunction().Walk(v)
	}
}

func (this *Function) Walk(v Visitor) {
	v = v.Visit(this)
	for _, e := range this.GetParams() {
		e.Walk(v)
	}
}

func (this *Terminal) Walk(v Visitor) {
	v = v.Visit(this)
	if this.Variable != nil {
		this.GetVariable().Walk(v)
	}
}

func (this *Variable) Walk(v Visitor) {
	v = v.Visit(this)
}
