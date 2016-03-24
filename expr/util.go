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

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/katydid/katydid/expr/token"
	"github.com/katydid/katydid/expr/types"
	"strconv"
	"strings"
	"unicode/utf8"
)

//NewKeyword is a parser utility function that returns a Keyword given a space and a token.
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

//NewSpace is a parser utility function that returns a Space given a token
func NewSpace(s interface{}) *Space {
	t := s.(*token.Token)
	return &Space{Space: []string{string(t.Lit)}}
}

//NewAppendSpace is a parser utility function that returns a Space by append the given string to the given Space's Space field, which is a list of strings.
func AppendSpace(ss interface{}, s string) *Space {
	space := ss.(*Space)
	space.Space = append(space.Space, s)
	return space
}

//SetTerminalSpace is a parser utility function that takes a Terminal and a Space and places the Space inside the returned Terminal.
func SetTerminalSpace(term interface{}, s interface{}) *Terminal {
	terminal := term.(*Terminal)
	terminal.Before = s.(*Space)
	return terminal
}

//SetExpComma is a parser utility function that takes an expression and a comma Keyword and places the comma inside the returned Expr.
func SetExprComma(e interface{}, c interface{}) *Expr {
	expr := e.(*Expr)
	expr.Comma = c.(*Keyword)
	return expr
}

//Strip is a parser utility function that removes all versions of the given sub string from the slit string and also removes possible surrounding parentheses.
func Strip(slit string, sub string) []byte {
	slit = strings.Replace(slit, sub, "", -1)
	if slit[0] != '(' {
		return []byte(slit)
	}
	return []byte(slit[1 : len(slit)-1])
}

//NewVariableTerminal is a parser utility function that returns a Terminal given a type.
func NewVariableTerminal(typ types.Type) (*Terminal, error) {
	return &Terminal{Variable: &Variable{Type: typ}}, nil
}

//NewBoolTerminal is a parser utility function that returns a Terminal of type bool given a bool.
func NewBoolTerminal(v interface{}) *Terminal {
	b := v.(bool)
	if b {
		return &Terminal{BoolValue: proto.Bool(b), Literal: "true"}
	}
	return &Terminal{BoolValue: proto.Bool(b), Literal: "false"}
}

//NewStringTerminal is a parser utility function that returns a Terminal of type string given a string literal.
//The input string is also unquoted.
func NewStringTerminal(slit string) (*Terminal, error) {
	return &Terminal{StringValue: proto.String(ToString(slit)), Literal: slit}, nil
}

//ToString unquotes a quoted string or returns the original string.
func ToString(s1 string) string {
	s, err := strconv.Unquote(s1)
	if err != nil {
		return s1
	}
	return s
}

//NewIntTerminal is a parser utility function that parses the int value out of the input string and returns a Terminal of type int.
func NewIntTerminal(slit string) (*Terminal, error) {
	return &Terminal{IntValue: ToInt64(Strip(slit, "int")), Literal: slit}, nil
}

//ToInt64 is a parser utility function that returns a pointer to a parsed an int64 or panics.
func ToInt64(tok []byte) *int64 {
	i, err := strconv.ParseInt(string(tok), 10, 64)
	if err != nil {
		panic(err)
	}
	return &i
}

//NewUintTerminal is a parser utility function that parses the uint value out of the input string and returns a Terminal of type uint.
func NewUintTerminal(slit string) (*Terminal, error) {
	return &Terminal{UintValue: ToUint64(Strip(slit, "uint")), Literal: slit}, nil
}

//ToUint64 is a parser utility function that returns a pointer to a parsed an uint64 or panics.
func ToUint64(tok []byte) *uint64 {
	i, err := strconv.ParseUint(string(tok), 10, 64)
	if err != nil {
		panic(err)
	}
	return &i
}

//NewDoubleTerminal is a parser utility function that parses the double value out of the input string and returns a Terminal of type double.
func NewDoubleTerminal(slit string) (*Terminal, error) {
	return &Terminal{DoubleValue: ToFloat64(Strip(slit, "double")), Literal: slit}, nil
}

//ToFloat64 is a parser utility function that returns a pointer to a parsed an float64 or panics.
func ToFloat64(tok []byte) *float64 {
	f, err := strconv.ParseFloat(string(tok), 64)
	if err != nil {
		panic(err)
	}
	return &f
}

//NewBytesTerminal is a parser utility function that parses the []byte value out of the input string and returns a Terminal of type []byte.
func NewBytesTerminal(stringLit string) (*Terminal, error) {
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
		r, _ := utf8.DecodeRune([]byte(s)[1:])
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

//NewSDTName is a parser utility function that returns a NameExpr given a white space and a terminal value expression.
func NewSDTName(space *Space, term *Terminal) *NameExpr {
	name := &NameExpr{
		Name: &Name{
			Before: space,
		},
	}
	if term.DoubleValue != nil {
		name.Name.DoubleValue = term.DoubleValue
	} else if term.IntValue != nil {
		name.Name.IntValue = term.IntValue
	} else if term.UintValue != nil {
		name.Name.UintValue = term.UintValue
	} else if term.BoolValue != nil {
		name.Name.BoolValue = term.BoolValue
	} else if term.StringValue != nil {
		name.Name.StringValue = term.StringValue
	} else if term.BytesValue != nil {
		name.Name.BytesValue = term.BytesValue
	} else {
		panic(fmt.Sprintf("unreachable name type %#v", term))
	}
	return name
}
