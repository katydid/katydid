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

	"github.com/awalterschulze/katydid/exp/asm/token"
	"github.com/awalterschulze/katydid/exp/asm/util"
	"strconv"
)

func IdToString(v interface{}) *string {
	t := v.(*token.Token)
	s := string(t.Lit)
	return &s
}

func ToString(v interface{}) *string {
	t := v.(*token.Token)
	s1 := string(t.Lit)
	s, err := strconv.Unquote(s1)
	if err != nil {
		return &s1
	}
	return &s
}

func ToInt64(v interface{}) *int64 {
	i, err := util.IntValue(v.(*token.Token).Lit)
	if err != nil {
		panic(err)
	}
	return &i
}

func ToUint64(v interface{}) *uint64 {
	i, err := util.UintValue(v.(*token.Token).Lit)
	if err != nil {
		panic(err)
	}
	return &i
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
