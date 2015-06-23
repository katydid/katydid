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

package asm

import (
	"github.com/katydid/katydid/expr/ast"
)

func (this *Rules) Walk(v expr.Visitor) {
	v = v.Visit(this)
	for i := range this.Rules {
		this.Rules[i].Walk(v)
	}
}

func (this *Rule) Walk(v expr.Visitor) {
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

func (this *Root) Walk(v expr.Visitor) {
	v.Visit(this)
}

func (this *Init) Walk(v expr.Visitor) {
	v.Visit(this)
}

func (this *Final) Walk(v expr.Visitor) {
	v.Visit(this)
}

func (this *Transition) Walk(v expr.Visitor) {
	v.Visit(this)
}

func (this *FunctionDecl) Walk(v expr.Visitor) {
	v.Visit(this)
	this.GetFunction().Walk(v)
}
