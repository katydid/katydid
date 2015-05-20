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
	for i := range this.Rules {
		this.Rules[i].Walk(v)
	}
}

func (this *Rule) Walk(v Visitor) {
	if this.Root != nil {
		this.Root.Walk(v)
	}
	if this.Init != nil {
		this.Init.Walk(v)
	}
	if this.Final != nil {
		this.Final.Walk(v)
	}
	if this.Transition != nil {
		this.Transition.Walk(v)
	}
	if this.FunctionDecl != nil {
		this.FunctionDecl.Walk(v)
	}
}

func (this *Root) Walk(v Visitor) {
	v.Visit(this)
}

func (this *Init) Walk(v Visitor) {
	v.Visit(this)
}

func (this *Final) Walk(v Visitor) {
	v.Visit(this)
}

func (this *Transition) Walk(v Visitor) {
	v.Visit(this)
}

func (this *FunctionDecl) Walk(v Visitor) {
	v.Visit(this)
	this.GetFunction().Walk(v)
}

func (this *Expr) Walk(v Visitor) {
	v = v.Visit(this)
	if this.Terminal != nil {
		this.GetTerminal().Walk(v)
	}
	if this.List != nil {
		this.GetList().Walk(v)
	}
	if this.Function != nil {
		this.GetFunction().Walk(v)
	}
}

func (this *List) Walk(v Visitor) {
	v = v.Visit(this)
	for _, e := range this.GetElems() {
		e.Walk(v)
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
