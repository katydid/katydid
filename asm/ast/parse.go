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

import (
	"code.google.com/p/gogoprotobuf/proto"
	"errors"
	"fmt"
	"github.com/awalterschulze/katydid/asm/token"
	"github.com/awalterschulze/katydid/asm/util"
	"github.com/awalterschulze/katydid/types"
	"strconv"
	"strings"
)

func (this *Rules) SetSpace(s *Space) *Rules {
	this.Final = s
	return this
}

func NewString(v interface{}) string {
	t := v.(*token.Token)
	return string(t.Lit)
}

func NewKeyword(space interface{}, v interface{}) *Keyword {
	t := v.(*token.Token)
	k := &Keyword{
		Value: string(t.Lit),
	}
	if space != nil {
		k.Before = space.(*Space)
	}
	return k
}

func NewSpace(s interface{}) *Space {
	t := s.(*token.Token)
	return &Space{Space: []string{string(t.Lit)}}
}

func AppendSpace(ss interface{}, s interface{}) *Space {
	space := ss.(*Space)
	t := s.(*token.Token)
	space.Space = append(space.Space, string(t.Lit))
	return space
}

func SetTerminalSpace(term interface{}, s interface{}) *Terminal {
	terminal := term.(*Terminal)
	terminal.Before = s.(*Space)
	return terminal
}

func SetExprComma(e interface{}, c interface{}) *Expr {
	expr := e.(*Expr)
	expr.Comma = c.(*Keyword)
	return expr
}

func Strip(slit string, sub string) []byte {
	slit = strings.Replace(slit, sub, "", -1)
	return []byte(slit[1 : len(slit)-1])
}

func NewVariableTerminal(variable interface{}, typ types.Type) (*Terminal, error) {
	varStr := string(variable.(*token.Token).Lit)
	vars := strings.Split(varStr, "(")
	name := vars[1][:len(vars[1])-1]
	return &Terminal{Variable: &Variable{Name: name, Type: typ}, Literal: varStr}, nil
}

func NewBoolTerminal(v interface{}) (*Terminal, error) {
	b := v.(bool)
	if b {
		return &Terminal{BoolValue: proto.Bool(b), Literal: "true"}, nil
	}
	return &Terminal{BoolValue: proto.Bool(b), Literal: "false"}, nil
}

func NewStringTerminal(v interface{}) (*Terminal, error) {
	slit := string(v.(*token.Token).Lit)
	return &Terminal{StringValue: proto.String(ToString(v)), Literal: slit}, nil
}

func ToString(v interface{}) string {
	t := v.(*token.Token)
	s1 := string(t.Lit)
	s, err := strconv.Unquote(s1)
	if err != nil {
		return s1
	}
	return s
}

func NewInt64Terminal(v interface{}) (*Terminal, error) {
	slit := string(v.(*token.Token).Lit)
	return &Terminal{Int64Value: ToInt64(Strip(slit, "int64")), Literal: slit}, nil
}

func ToInt64(tok []byte) *int64 {
	i, err := util.IntValue(tok)
	if err != nil {
		panic(err)
	}
	return &i
}

func NewInt32Terminal(v interface{}) (*Terminal, error) {
	slit := string(v.(*token.Token).Lit)
	return &Terminal{Int32Value: ToInt32(Strip(slit, "int32")), Literal: slit}, nil
}

func ToInt32(tok []byte) *int32 {
	i, err := util.IntValue(tok)
	if err != nil {
		panic(err)
	}
	ii := int32(i)
	return &ii
}

func NewUint64Terminal(v interface{}) (*Terminal, error) {
	slit := string(v.(*token.Token).Lit)
	return &Terminal{Uint64Value: ToUint64(Strip(slit, "uint64")), Literal: slit}, nil
}

func ToUint64(tok []byte) *uint64 {
	i, err := util.UintValue(tok)
	if err != nil {
		panic(err)
	}
	return &i
}

func NewUint32Terminal(v interface{}) (*Terminal, error) {
	slit := string(v.(*token.Token).Lit)
	return &Terminal{Uint32Value: ToUint32(Strip(slit, "uint32")), Literal: slit}, nil
}

func ToUint32(tok []byte) *uint32 {
	i, err := util.UintValue(tok)
	if err != nil {
		panic(err)
	}
	ii := uint32(i)
	return &ii
}

func NewDoubleTerminal(v interface{}) (*Terminal, error) {
	slit := string(v.(*token.Token).Lit)
	return &Terminal{DoubleValue: ToFloat64(Strip(slit, "double")), Literal: slit}, nil
}

func ToFloat64(tok []byte) *float64 {
	f, err := strconv.ParseFloat(string(tok), 64)
	if err != nil {
		panic(err)
	}
	return &f
}

func NewFloatTerminal(v interface{}) (*Terminal, error) {
	slit := string(v.(*token.Token).Lit)
	return &Terminal{FloatValue: ToFloat32(Strip(slit, "float")), Literal: slit}, nil
}

func ToFloat32(tok []byte) *float32 {
	f, err := strconv.ParseFloat(string(tok), 32)
	if err != nil {
		panic(err)
	}
	ff := float32(f)
	return &ff
}

func NewBytesTerminal(v interface{}) (*Terminal, error) {
	lit := v.(*token.Token).Lit
	stringLit := string(lit)
	data, err := parseBytes(stringLit)
	if err != nil {
		return nil, err
	}
	return &Terminal{BytesValue: data, Literal: stringLit}, nil
}

func parseBytes(s string) ([]byte, error) {
	byteElems := strings.Split(s[strings.Index(s, "{")+1:strings.LastIndex(s, "}")], ",")
	data := make([]byte, 0, len(byteElems))
	for _, b := range byteElems {
		s := strings.TrimSpace(b)
		if len(s) == 0 {
			continue
		}
		d, err := parseByte(s)
		if err != nil {
			return nil, err
		}
		data = append(data, d)
	}
	return data, nil
}

func hexToByte(c byte) byte {
	if 'a' <= c && c <= 'f' {
		return c - 'a' + 10
	}
	if 'A' <= c && c <= 'F' {
		return c - 'A' + 10
	}
	return c - '0'
}

func hexesToByte(a byte, b byte) byte {
	aa, bb := hexToByte(a), hexToByte(b)
	return byte(aa*16 + bb)
}

func parseByte(s string) (byte, error) {
	if s[0] == '\'' {
		r := util.RuneValue([]byte(s))
		if r <= 255 {
			return byte(r), nil
		}
		return 0, fmt.Errorf("rune too large %v", r)
	} else if s[0] == '0' {
		if len(s) == 1 {
			return 0, nil
		}
		if s[1] == 'x' || s[1] == 'X' {
			if len(s) == 4 {
				return hexesToByte(s[2], s[3]), nil
			} else if len(s) == 3 {
				return hexToByte(s[2]), nil
			}
			return 0, fmt.Errorf("not a hex digit %v", s)
		} else {
			switch len(s) {
			case 4:
				o := (int(s[1]-'0') * 64) + (int(s[2]-'0') * 8) + int(s[3]-'0')
				if o >= 255 {
					return 0, fmt.Errorf("octal too large %d", o)
				}
				return byte(o), nil
			case 3:
				return byte((s[1]-'0')*8 + (s[2] - '0')), nil
			case 2:
				return byte(s[1] - '0'), nil
			}
			return 0, nil
		}
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	if i >= 0 && i <= 255 {
		return byte(i), nil
	}
	return 0, fmt.Errorf("int too large %d", i)
}

func newRules(r *Rule) *Rules {
	return &Rules{Rules: []*Rule{r}}
}

func NewRule(r interface{}) (*Rules, error) {
	switch rule := r.(type) {
	case *Root:
		return newRules(&Rule{Root: rule}), nil
	case *Init:
		return newRules(&Rule{Init: rule}), nil
	case *Transition:
		return newRules(&Rule{Transition: rule}), nil
	case *IfExpr:
		return newRules(&Rule{IfExpr: rule}), nil
	}
	panic("unreachable")
}

func AppendRule(rs, r interface{}) (*Rules, error) {
	rules := rs.(*Rules)
	switch rule := r.(type) {
	case *Root:
		if rules.HasRoot() {
			return nil, errors.New("only one root is allowed")
		}
		root := &Rule{Root: rule}
		for _, r := range rules.Rules {
			if r.Init != nil {
				if r.Init.Message == rule.Message &&
					r.Init.Package == rule.Package {
					root.Root.State = r.Init.State
				}
			}
		}
		rules.Rules = append(rules.Rules, root)
		return rules, nil
	case *Init:
		if root := rules.GetRoot(); root != nil {
			if rule.Package == root.Package &&
				rule.Message == root.Message {
				if root.State != "root" {
					return nil, errors.New("root can only start in a single state")
				}
				root.State = rule.State
			}
		}
		rules.Rules = append(rules.Rules, &Rule{Init: rule})
		return rules, nil
	case *Transition:
		rules.Rules = append(rules.Rules, &Rule{Transition: rule})
		return rules, nil
	case *IfExpr:
		rules.Rules = append(rules.Rules, &Rule{IfExpr: rule})
		return rules, nil
	}
	panic("unreachable")
}
