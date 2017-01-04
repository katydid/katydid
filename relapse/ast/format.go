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

import (
	"strings"
)

//Format formats the spacing in the Grammar.
func (this *Grammar) Format() {
	if this.TopPattern != nil {
		this.TopPattern.format(false)
	}
	for i, p := range this.PatternDecls {
		p.format(i != 0 || this.TopPattern != nil)
	}
	this.After.format(false)
}

func (this *PatternDecl) format(space bool) {
	this.Hash.format(space)
	this.Before.format(false)
	this.Eq.format(true)
	this.Pattern.format(true)
}

//Format formats the spacing in the Pattern.
func (this *Pattern) Format() {
	this.format(false)
}

func (this *Pattern) format(space bool) {
	v := this.GetValue()
	v.(interface {
		format(bool)
	}).format(space)
}

func (this *Empty) format(space bool) {
	this.Empty.format(space)
}

func (this *TreeNode) format(space bool) {
	this.Name.format(space)
	if this.Colon != nil {
		this.Colon.format(false)
	}
	this.Pattern.format(this.Colon == nil)
}

func (this *Contains) format(space bool) {
	this.Dot.format(space)
	this.Pattern.format(false)
}

func (this *LeafNode) format(space bool) {
	this.Expr.format(space)
}

func (this *Concat) format(space bool) {
	this.OpenBracket.format(space)
	this.LeftPattern.format(false)
	this.Comma.format(false)
	this.RightPattern.format(true)
	if this.ExtraComma != nil {
		this.ExtraComma.format(false)
	}
	this.CloseBracket.format(this.ExtraComma == nil)
}

func (this *Or) format(space bool) {
	this.OpenParen.format(space)
	this.LeftPattern.format(false)
	this.Pipe.format(true)
	this.RightPattern.format(true)
	this.CloseParen.format(false)
}

func (this *And) format(space bool) {
	this.OpenParen.format(space)
	this.LeftPattern.format(false)
	this.Ampersand.format(true)
	this.RightPattern.format(true)
	this.CloseParen.format(false)
}

func (this *ZeroOrMore) format(space bool) {
	this.OpenParen.format(space)
	this.Pattern.format(false)
	this.CloseParen.format(false)
	this.Star.format(false)
}

func (this *Reference) format(space bool) {
	this.At.format(space)
}

func (this *Not) format(space bool) {
	this.Exclamation.format(space)
	this.OpenParen.format(false)
	this.Pattern.format(false)
	this.CloseParen.format(false)
}

func (this *ZAny) format(space bool) {
	this.Star.format(space)
}

func (this *Optional) format(space bool) {
	this.OpenParen.format(space)
	this.Pattern.format(false)
	this.CloseParen.format(false)
	this.QuestionMark.format(false)
}

func (this *Interleave) format(space bool) {
	this.OpenCurly.format(space)
	this.LeftPattern.format(false)
	this.SemiColon.format(false)
	this.RightPattern.format(true)
	if this.ExtraSemiColon != nil {
		this.ExtraSemiColon.format(false)
	}
	this.CloseCurly.format(this.ExtraSemiColon == nil)
}

func (this *Expr) format(space bool) {
	if this.RightArrow != nil {
		this.RightArrow.format(space)
		space = false
	} else if this.Comma != nil {
		this.Comma.format(space)
		space = false
	}
	if this.Terminal != nil {
		this.Terminal.format(space)
	} else if this.List != nil {
		this.List.format(space)
	} else if this.Function != nil {
		this.Function.format(space)
	} else if this.BuiltIn != nil {
		this.BuiltIn.format(space)
	}
}

func (this *NameExpr) format(space bool) {
	v := this.GetValue()
	v.(interface {
		format(bool)
	}).format(space)
}

func (this *Name) format(space bool) {
	this.Before.format(space)
}

func (this *AnyName) format(space bool) {
	this.Underscore.format(space)
}

func (this *AnyNameExcept) format(space bool) {
	this.Exclamation.format(space)
	this.OpenParen.format(false)
	this.Except.format(false)
	this.CloseParen.format(false)
}

func (this *NameChoice) format(space bool) {
	this.OpenParen.format(space)
	this.Left.format(false)
	this.Pipe.format(false)
	this.Right.format(false)
	this.CloseParen.format(false)
}

func (this *List) format(space bool) {
	this.Before.format(space)
	this.OpenCurly.format(false)
	for i := range this.Elems {
		this.Elems[i].format(i != 0)
	}
	this.CloseCurly.format(false)
}

func (this *Function) format(space bool) {
	this.Before.format(space)
	this.OpenParen.format(false)
	for i := range this.Params {
		this.Params[i].format(i != 0)
	}
	this.CloseParen.format(false)
}

func (this *BuiltIn) format(space bool) {
	this.Symbol.format(space)
	this.Expr.format(true)
}

func (this *Terminal) format(space bool) {
	this.Before.format(space)
}

func (this *Keyword) format(space bool) {
	if this == nil {
		return
	}
	this.Before.format(space)
}

func (this *Space) Format() {
	this.format(false)
}

func (this *Space) format(space bool) {
	if this == nil {
		return
	}
	newSpace := formatSpace(this.Space)
	if space {
		if len(this.Space) > 0 && strings.Contains(this.Space[len(this.Space)-1], "\n") {
			newSpace = append(newSpace, "\n")
		} else {
			newSpace = append(newSpace, " ")
		}
	}
	this.Space = newSpace
}

type state int

var startState = state(0)
var lineCommentState = state(1)
var blockCommentState = state(2)

func formatSpace(spaces []string) []string {
	newlines := 0
	current := startState
	formatted := []string{}
	for _, space := range spaces {
		for _, c := range space {
			if c == '\n' {
				newlines++
			}
		}
		if isComment(space) {
			comment := strings.TrimSpace(space)
			if isLineComment(space) {
				comment = comment + "\n"
			}
			switch current {
			case startState:
				formatted = append(formatted, comment)
			case lineCommentState:
				if newlines >= 2 {
					formatted = append(formatted, "\n")
				}
				formatted = append(formatted, comment)
			case blockCommentState:
				if newlines >= 2 {
					formatted = append(formatted, "\n\n")
				}
				formatted = append(formatted, comment)
			}
			if isLineComment(space) {
				current = lineCommentState
				newlines = 1
			} else {
				current = blockCommentState
				newlines = 0
			}
		}
	}
	return formatted
}
