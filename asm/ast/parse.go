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
	"errors"
	"fmt"
	"github.com/awalterschulze/katydid/asm/token"
	"github.com/awalterschulze/katydid/asm/util"
	"strconv"
	"strings"
)

func IdToString(v interface{}) string {
	t := v.(*token.Token)
	return string(t.Lit)
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

func Strip(v interface{}, sub string) []byte {
	slit := string(v.(*token.Token).Lit)
	slit = strings.Replace(slit, sub, "", -1)
	return []byte(slit[1 : len(slit)-1])
}

func ToInt64(tok []byte) *int64 {
	i, err := util.IntValue(tok)
	if err != nil {
		panic(err)
	}
	return &i
}

func ToInt32(tok []byte) *int32 {
	i, err := util.IntValue(tok)
	if err != nil {
		panic(err)
	}
	ii := int32(i)
	return &ii
}

func ToUint64(tok []byte) *uint64 {
	i, err := util.UintValue(tok)
	if err != nil {
		panic(err)
	}
	return &i
}

func ToUint32(tok []byte) *uint32 {
	i, err := util.UintValue(tok)
	if err != nil {
		panic(err)
	}
	ii := uint32(i)
	return &ii
}

func ToFloat64(tok []byte) *float64 {
	f, err := strconv.ParseFloat(string(tok), 64)
	if err != nil {
		panic(err)
	}
	return &f
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
	data, err := parseBytes(string(lit))
	if err != nil {
		return nil, err
	}
	return &Terminal{BytesValue: data}, nil
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
	} else {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		if i >= 0 && i <= 255 {
			return byte(i), nil
		}
		return 0, fmt.Errorf("int too large %d", i)
	}
	panic("unreachable")
}

func NewRule(r interface{}) (*Rules, error) {
	switch rule := r.(type) {
	case *Init:
		if rule.GetState() == "root" {
			return &Rules{Root: rule}, nil
		}
		return &Rules{Init: []*Init{rule}}, nil
	case *Transition:
		return &Rules{Transition: []*Transition{rule}}, nil
	case *IfExpr:
		return &Rules{IfExpr: []*IfExpr{rule}}, nil
	}
	panic("unreachable")
}

func AppendRule(rs, r interface{}) (*Rules, error) {
	rules := rs.(*Rules)
	switch rule := r.(type) {
	case *Init:
		if rule.GetState() == "root" {
			if rules.Root != nil {
				return nil, errors.New("only one root is allowed")
			}
			rules.Root = rule
			for _, i := range rules.GetInit() {
				if i.GetPackage() == rules.GetRoot().GetPackage() &&
					i.GetMessage() == rules.GetRoot().GetMessage() {
					rules.Root.State = i.State
				}
			}
			return rules, nil
		} else {
			if rules.Root != nil {
				if rule.GetPackage() == rules.GetRoot().GetPackage() &&
					rule.GetMessage() == rules.GetRoot().GetMessage() {
					if rules.GetRoot().GetState() != "root" {
						return nil, errors.New("root can only start in a single state")
					}
					rules.Root.State = rule.State
				}
			}
		}
		rules.Init = append(rules.Init, rule)
		return rules, nil
	case *Transition:
		rules.Transition = append(rules.Transition, rule)
		return rules, nil
	case *IfExpr:
		rules.IfExpr = append(rules.IfExpr, rule)
		return rules, nil
	}
	panic("unreachable")
}
