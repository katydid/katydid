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

//Visitor is used to do a top down walk through the expression tree using the Walk methods.
type Visitor interface {
	//Visit takes in an ast type and should return another type that implementes the Visitor interface.
	Visit(node interface{}) interface{}
}

//Walk visits every possible Expr field that is not nil and not of type Space.
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

//Walk visits every possible NameExpr field that is not nil and not of type Space.
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

//Walk visits Name.
func (this *Name) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
}

//Walk visits AnyName.
func (this *AnyName) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
}

//Walk visits AnyNameExcept and its NameExpr.
func (this *AnyNameExcept) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.Except != nil {
		this.GetExcept().Walk(v)
	}
}

//Walk visits NameChoice and its NameExpr types.
func (this *NameChoice) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.Left != nil {
		this.GetLeft().Walk(v)
	}
	if this.Right != nil {
		this.GetRight().Walk(v)
	}
}

//Walk visits List and each element in the list.
func (this *List) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	for _, e := range this.GetElems() {
		e.Walk(v)
	}
}

//Walk visits Function and every parameter.
func (this *Function) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	for _, e := range this.GetParams() {
		e.Walk(v)
	}
}

//Walk visits BuiltIn and its Expr.
func (this *BuiltIn) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	this.Expr.Walk(v)
}

//Walk visits Terminal and its possible Variable.
func (this *Terminal) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.Variable != nil {
		this.GetVariable().Walk(v)
	}
}

//Walk visits Variable.
func (this *Variable) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
}
