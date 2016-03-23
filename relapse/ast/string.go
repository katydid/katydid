//  Copyright 2015 Walter Schulze
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

import (
	"strings"
)

func (this *Grammar) String() string {
	ss := make([]string, len(this.PatternDecls)+1)
	for i, p := range this.PatternDecls {
		ss[i] = p.String()
	}
	ss[len(ss)-1] = this.After.String()
	if this.TopPattern != nil {
		return this.TopPattern.String() + strings.Join(ss, "")
	}
	return strings.Join(ss, "")
}

func (this *PatternDecl) String() string {
	return this.Hash.String() + this.Before.String() + this.Name +
		this.Eq.String() + this.Pattern.String()
}

func (this *Pattern) String() string {
	v := this.GetValue()
	return v.(interface {
		String() string
	}).String()
}

func (this *Empty) String() string {
	return this.Empty.String()
}

func (this *TreeNode) String() string {
	return this.Name.String() + this.Colon.String() +
		this.Pattern.String()
}

func (this *Contains) String() string {
	return this.Dot.String() + this.Pattern.String()
}

func (this *LeafNode) String() string {
	return this.RightArrow.String() + this.Expr.String()
}

func (this *Concat) String() string {
	return this.OpenBracket.String() + this.LeftPattern.String() +
		this.Comma.String() + this.RightPattern.String() +
		this.ExtraComma.String() + this.CloseBracket.String()
}

func (this *Or) String() string {
	return this.OpenParen.String() + this.LeftPattern.String() +
		this.Pipe.String() + this.RightPattern.String() +
		this.CloseParen.String()
}

func (this *And) String() string {
	return this.OpenParen.String() + this.LeftPattern.String() +
		this.Ampersand.String() + this.RightPattern.String() +
		this.CloseParen.String()
}

func (this *ZeroOrMore) String() string {
	return this.OpenParen.String() + this.Pattern.String() +
		this.CloseParen.String() + this.Star.String()
}

func (this *Reference) String() string {
	return this.At.String() + this.Name
}

func (this *Not) String() string {
	return this.Exclamation.String() +
		this.OpenParen.String() + this.Pattern.String() +
		this.CloseParen.String()
}

func (this *ZAny) String() string {
	return this.Star.String()
}

func (this *Optional) String() string {
	return this.OpenParen.String() + this.Pattern.String() +
		this.CloseParen.String() + this.QuestionMark.String()
}

func (this *Interleave) String() string {
	return this.OpenCurly.String() + this.LeftPattern.String() +
		this.SemiColon.String() + this.RightPattern.String() +
		this.ExtraSemiColon.String() + this.CloseCurly.String()
}
