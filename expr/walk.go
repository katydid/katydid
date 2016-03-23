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

package expr

type Visitor interface {
	Visit(node interface{}) interface{}
}

func (this *Expr) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.Terminal != nil {
		this.GetTerminal().Walk(v)
	}
	if this.List != nil {
		this.GetList().Walk(v)
	}
	if this.Function != nil {
		this.GetFunction().Walk(v)
	}
	if this.BuiltIn != nil {
		this.GetBuiltIn().Walk(v)
	}
}

func (this *NameExpr) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.Name != nil {
		this.GetName().Walk(v)
	}
	if this.AnyName != nil {
		this.GetAnyName().Walk(v)
	}
	if this.AnyNameExcept != nil {
		this.GetAnyNameExcept().Walk(v)
	}
	if this.NameChoice != nil {
		this.GetNameChoice().Walk(v)
	}
}

func (this *Name) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
}

func (this *AnyName) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
}

func (this *AnyNameExcept) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.Except != nil {
		this.GetExcept().Walk(v)
	}
}

func (this *NameChoice) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.Left != nil {
		this.GetLeft().Walk(v)
	}
	if this.Right != nil {
		this.GetRight().Walk(v)
	}
}

func (this *List) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	for _, e := range this.GetElems() {
		e.Walk(v)
	}
}

func (this *Function) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	for _, e := range this.GetParams() {
		e.Walk(v)
	}
}

func (this *BuiltIn) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	this.Expr.Walk(v)
}

func (this *Terminal) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.Variable != nil {
		this.GetVariable().Walk(v)
	}
}

func (this *Variable) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
}
