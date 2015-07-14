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
	"github.com/katydid/katydid/types"
	"strconv"
	"strings"
	"unicode/utf8"
)

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

func AppendSpace(ss interface{}, s string) *Space {
	space := ss.(*Space)
	space.Space = append(space.Space, s)
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

func NewVariableTerminal(typ types.Type) (*Terminal, error) {
	return &Terminal{Variable: &Variable{Type: typ}}, nil
}

func NewBoolTerminal(v interface{}) (*Terminal, error) {
	b := v.(bool)
	if b {
		return &Terminal{BoolValue: proto.Bool(b), Literal: "true"}, nil
	}
	return &Terminal{BoolValue: proto.Bool(b), Literal: "false"}, nil
}

func NewStringTerminal(slit string) (*Terminal, error) {
	return &Terminal{StringValue: proto.String(ToString(slit)), Literal: slit}, nil
}

func ToString(s1 string) string {
	s, err := strconv.Unquote(s1)
	if err != nil {
		return s1
	}
	return s
}

func NewIntTerminal(slit string) (*Terminal, error) {
	return &Terminal{IntValue: ToInt64(Strip(slit, "int")), Literal: slit}, nil
}

func ToInt64(tok []byte) *int64 {
	i, err := strconv.ParseInt(string(tok), 10, 64)
	if err != nil {
		panic(err)
	}
	return &i
}

func NewUintTerminal(slit string) (*Terminal, error) {
	return &Terminal{UintValue: ToUint64(Strip(slit, "uint")), Literal: slit}, nil
}

func ToUint64(tok []byte) *uint64 {
	i, err := strconv.ParseUint(string(tok), 10, 64)
	if err != nil {
		panic(err)
	}
	return &i
}

func NewDoubleTerminal(slit string) (*Terminal, error) {
	return &Terminal{DoubleValue: ToFloat64(Strip(slit, "double")), Literal: slit}, nil
}

func ToFloat64(tok []byte) *float64 {
	f, err := strconv.ParseFloat(string(tok), 64)
	if err != nil {
		panic(err)
	}
	return &f
}

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
