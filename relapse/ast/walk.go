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

package relapse

type Visitor interface {
	Visit(node interface{}) Visitor
}

func (this *Grammar) Walk(v Visitor) {
	v = v.Visit(this)
	if this.TopPattern != nil {
		this.GetTopPattern().Walk(v)
	}
	for _, pdecl := range this.PatternDecls {
		pdecl.Walk(v)
	}
}

func (this *PatternDecl) Walk(v Visitor) {
	v = v.Visit(this)
	if this.Pattern != nil {
		this.GetPattern().Walk(v)
	}
}

func (this *Pattern) Walk(v Visitor) {
	v = v.Visit(this)
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

func (this *Empty) Walk(v Visitor) {
	v = v.Visit(this)
}

func (this *TreeNode) Walk(v Visitor) {
	v = v.Visit(this)
	if this.Name != nil {
		this.GetName().Walk(v)
	}
	if this.Pattern != nil {
		this.GetPattern().Walk(v)
	}
}

func (this *LeafNode) Walk(v Visitor) {
	v = v.Visit(this)
}

func (this *Concat) Walk(v Visitor) {
	v = v.Visit(this)
	if this.LeftPattern != nil {
		this.GetLeftPattern().Walk(v)
	}
	if this.RightPattern != nil {
		this.GetRightPattern().Walk(v)
	}
}

func (this *Or) Walk(v Visitor) {
	v = v.Visit(this)
	if this.LeftPattern != nil {
		this.GetLeftPattern().Walk(v)
	}
	if this.RightPattern != nil {
		this.GetRightPattern().Walk(v)
	}
}

func (this *And) Walk(v Visitor) {
	v = v.Visit(this)
	if this.LeftPattern != nil {
		this.GetLeftPattern().Walk(v)
	}
	if this.RightPattern != nil {
		this.GetRightPattern().Walk(v)
	}
}

func (this *ZeroOrMore) Walk(v Visitor) {
	v = v.Visit(this)
	if this.Pattern != nil {
		this.GetPattern().Walk(v)
	}
}

func (this *Reference) Walk(v Visitor) {
	v = v.Visit(this)
}

func (this *Not) Walk(v Visitor) {
	v = v.Visit(this)
	if this.Pattern != nil {
		this.GetPattern().Walk(v)
	}
}

func (this *ZAny) Walk(v Visitor) {
	v = v.Visit(this)
}

func (this *Contains) Walk(v Visitor) {
	v = v.Visit(this)
	if this.Pattern != nil {
		this.GetPattern().Walk(v)
	}
}

func (this *Optional) Walk(v Visitor) {
	v = v.Visit(this)
	if this.Pattern != nil {
		this.GetPattern().Walk(v)
	}
}

func (this *Interleave) Walk(v Visitor) {
	v = v.Visit(this)
	if this.LeftPattern != nil {
		this.GetLeftPattern().Walk(v)
	}
	if this.RightPattern != nil {
		this.GetRightPattern().Walk(v)
	}
}

func (this *NameExpr) Walk(v Visitor) {
	v = v.Visit(this)
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
	v = v.Visit(this)
}

func (this *AnyName) Walk(v Visitor) {
	v = v.Visit(this)
}

func (this *AnyNameExcept) Walk(v Visitor) {
	v = v.Visit(this)
	if this.Except != nil {
		this.GetExcept().Walk(v)
	}
}

func (this *NameChoice) Walk(v Visitor) {
	v = v.Visit(this)
	if this.Left != nil {
		this.GetLeft().Walk(v)
	}
	if this.Right != nil {
		this.GetRight().Walk(v)
	}
}

type vNot struct {
	has bool
}

func (this *vNot) Visit(node interface{}) Visitor {
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
