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

package ast

import (
	"fmt"
	"github.com/katydid/katydid/relapse/types"
	"regexp"
	"strconv"
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
	return this.Expr.String()
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

//String returns the parsable expression string.
func (this *Expr) String() string {
	space := this.RightArrow.String() + this.Comma.String()
	if this.Terminal != nil {
		return space + this.GetTerminal().String()
	}
	if this.List != nil {
		return space + this.GetList().String()
	}
	if this.Function != nil {
		return space + this.GetFunction().String()
	}
	return this.BuiltIn.String()
}

//String returns the parsable name expression string.
func (this *NameExpr) String() string {
	v := this.GetValue()
	return v.(interface {
		String() string
	}).String()
}

// _decimal_digit : '0' - '9' ;
// _upcase : 'A'-'Z' ;
// _lowcase : 'a'-'z' ;
// _id_char : _upcase | _lowcase | '_' | _decimal_digit ;
// _id : (_upcase | _lowcase | '_' ) {_id_char} ;
// id : _id ;
var idRegexp = regexp.MustCompile("^[a-zA-Z_][_a-zA-Z0-9]*$")

func isId(s string) bool {
	return idRegexp.MatchString(s)
}

//String returns the parsable name string.
func (this *Name) String() string {
	if this.DoubleValue != nil {
		return this.Before.String() + strconv.FormatFloat(this.GetDoubleValue(), 'f', -1, 10)
	}
	if this.IntValue != nil {
		return this.Before.String() + strconv.FormatInt(this.GetIntValue(), 10)
	}
	if this.UintValue != nil {
		return this.Before.String() + "uint(" + strconv.FormatUint(this.GetUintValue(), 10) + ")"
	}
	if this.BoolValue != nil {
		return this.Before.String() + strconv.FormatBool(this.GetBoolValue())
	}
	if this.StringValue != nil {
		if isId(this.GetStringValue()) {
			return this.Before.String() + this.GetStringValue()
		}
		return this.Before.String() + strconv.Quote(this.GetStringValue())
	}
	if this.BytesValue != nil {
		return this.Before.String() + fmt.Sprintf("%#v", this.GetBytesValue())
	}
	panic("unreachable")
}

//String returns the parsable any name string.
func (this *AnyName) String() string {
	return this.Underscore.String()
}

//String returns the parsable any name except string.
func (this *AnyNameExcept) String() string {
	return this.Exclamation.String() + this.OpenParen.String() +
		this.Except.String() + this.CloseParen.String()
}

//String returns the parsable name choice string.
func (this *NameChoice) String() string {
	return this.OpenParen.String() + this.Left.String() +
		this.Pipe.String() + this.Right.String() +
		this.CloseParen.String()
}

//String returns the parsable typed list string.
func (this *List) String() string {
	es := make([]string, len(this.GetElems()))
	for i, v := range this.GetElems() {
		es[i] = v.String()
	}
	return this.Before.String() + listTypeToString(this.Type) + this.OpenCurly.String() + strings.Join(es, "") + this.CloseCurly.String()
}

func listTypeToString(typ types.Type) string {
	switch typ {
	case types.LIST_DOUBLE:
		return "[]double"
	case types.LIST_INT:
		return "[]int"
	case types.LIST_UINT:
		return "[]uint"
	case types.LIST_BOOL:
		return "[]bool"
	case types.LIST_STRING:
		return "[]string"
	case types.LIST_BYTES:
		return "[][]byte"
	}
	panic("unreachable")
}

//String returns the parsable function string.
func (this *Function) String() string {
	ps := make([]string, len(this.GetParams()))
	for i, v := range this.GetParams() {
		ps[i] = v.String()
	}
	return this.Before.String() + this.GetName() + this.OpenParen.String() + strings.Join(ps, "") + this.CloseParen.String()
}

//String returns the parsable builtin string.
func (this *BuiltIn) String() string {
	return this.Symbol.String() + this.Expr.String()
}

//String returns the parsable terminal string.
func (this *Terminal) String() string {
	if this.DoubleValue != nil {
		return this.Before.String() + strconv.FormatFloat(this.GetDoubleValue(), 'f', -1, 64)
	}
	if this.IntValue != nil {
		return this.Before.String() + strconv.FormatInt(this.GetIntValue(), 10)
	}
	if this.UintValue != nil {
		return this.Before.String() + "uint(" + strconv.FormatUint(this.GetUintValue(), 10) + ")"
	}
	if this.BoolValue != nil {
		return this.Before.String() + strconv.FormatBool(this.GetBoolValue())
	}
	if this.StringValue != nil {
		return this.Before.String() + strconv.Quote(this.GetStringValue())
	}
	if this.BytesValue != nil {
		return this.Before.String() + fmt.Sprintf("%#v", this.GetBytesValue())
	}
	if this.Variable != nil {
		return this.Before.String() + this.Variable.String()
	}
	panic("unreachable")
}

//String returns the parsable variable string.
func (this *Variable) String() string {
	typ := this.GetType()
	if types.IsList(typ) {
		types.ListToSingle(typ)
	}
	switch typ {
	case types.SINGLE_DOUBLE:
		return "$double"
	case types.SINGLE_INT:
		return "$int"
	case types.SINGLE_UINT:
		return "$uint"
	case types.SINGLE_BOOL:
		return "$bool"
	case types.SINGLE_STRING:
		return "$string"
	case types.SINGLE_BYTES:
		return "$[]byte"
	}
	panic(fmt.Errorf("unknown type %s", this.GetType()))
}

//String returns the parsable keyword string.
func (this *Keyword) String() string {
	if this == nil {
		return ""
	}
	return this.Before.String() + this.Value
}

//String returns the parsable space string.
func (this *Space) String() string {
	if this == nil {
		return ""
	}
	return strings.Join(this.Space, "")
}
