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

package ast

//Visitor is used to do a top down walk through the expression tree using the Walk methods.
type Visitor interface {
	//Visit takes in an ast type and should return another type that implementes the Visitor interface.
	Visit(node interface{}) interface{}
}

//Walk visits the Grammar instance and all its possible child pattern and all its pattern declarations.
func (this *Grammar) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.TopPattern != nil {
		this.GetTopPattern().Walk(v)
	}
	for _, pdecl := range this.PatternDecls {
		pdecl.Walk(v)
	}
}

//Walk visits the PatternDecl instance and its child pattern.
func (this *PatternDecl) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.Pattern != nil {
		this.GetPattern().Walk(v)
	}
}

//Walk visits the Pattern instance, its non nil patterns.
func (this *Pattern) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.Empty != nil {
		this.GetEmpty().Walk(v)
	}
	if this.TreeNode != nil {
		this.GetTreeNode().Walk(v)
	}
	if this.LeafNode != nil {
		this.GetLeafNode().Walk(v)
	}
	if this.Concat != nil {
		this.GetConcat().Walk(v)
	}
	if this.Or != nil {
		this.GetOr().Walk(v)
	}
	if this.And != nil {
		this.GetAnd().Walk(v)
	}
	if this.ZeroOrMore != nil {
		this.GetZeroOrMore().Walk(v)
	}
	if this.Reference != nil {
		this.GetReference().Walk(v)
	}
	if this.Not != nil {
		this.GetNot().Walk(v)
	}
	if this.ZAny != nil {
		this.GetZAny().Walk(v)
	}
	if this.Contains != nil {
		this.GetContains().Walk(v)
	}
	if this.Optional != nil {
		this.GetOptional().Walk(v)
	}
	if this.Interleave != nil {
		this.GetInterleave().Walk(v)
	}
}

//Walk visits the Empty pattern.
func (this *Empty) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
}

//Walk visits the TreeNode pattern, its name and its pattern.
func (this *TreeNode) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.Name != nil {
		this.GetName().Walk(v)
	}
	if this.Pattern != nil {
		this.GetPattern().Walk(v)
	}
}

//Walk visits the LeafNode pattern.
func (this *LeafNode) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
}

//Walk visits the Concat pattern and its child patterns.
func (this *Concat) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.LeftPattern != nil {
		this.GetLeftPattern().Walk(v)
	}
	if this.RightPattern != nil {
		this.GetRightPattern().Walk(v)
	}
}

//Walk visits the Or pattern and its child patterns.
func (this *Or) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.LeftPattern != nil {
		this.GetLeftPattern().Walk(v)
	}
	if this.RightPattern != nil {
		this.GetRightPattern().Walk(v)
	}
}

//Walk visits the And pattern and its child patterns.
func (this *And) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.LeftPattern != nil {
		this.GetLeftPattern().Walk(v)
	}
	if this.RightPattern != nil {
		this.GetRightPattern().Walk(v)
	}
}

//Walk visits the ZeroOrMore pattern and its pattern.
func (this *ZeroOrMore) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.Pattern != nil {
		this.GetPattern().Walk(v)
	}
}

//Walk visits the Reference pattern.
func (this *Reference) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
}

//Walk visits the Not pattern and its pattern.
func (this *Not) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.Pattern != nil {
		this.GetPattern().Walk(v)
	}
}

//Walk visits the ZAny pattern.
func (this *ZAny) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
}

//Walk visits the Contains pattern and its pattern.
func (this *Contains) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.Pattern != nil {
		this.GetPattern().Walk(v)
	}
}

//Walk visits the Optional pattern and its pattern.
func (this *Optional) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.Pattern != nil {
		this.GetPattern().Walk(v)
	}
}

//Walk visits the Interleave pattern and its child patterns.
func (this *Interleave) Walk(v Visitor) {
	v = v.Visit(this).(Visitor)
	if this.LeftPattern != nil {
		this.GetLeftPattern().Walk(v)
	}
	if this.RightPattern != nil {
		this.GetRightPattern().Walk(v)
	}
}

//Walk visits every possible field that is not nil and not of type Keyword or Space.
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

//Walk visits every possible field that is not nil and not of type Keyword or Space.
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

type vNot struct {
	has bool
}

func (this *vNot) Visit(node interface{}) interface{} {
	n, ok := node.(*Not)
	if !ok {
		return this
	}
	if n.GetPattern().ZAny != nil {
		return this
	}
	this.has = true
	return this
}

func HasNot(g *Grammar) bool {
	vNot := &vNot{}
	g.Walk(vNot)
	return vNot.has
}
