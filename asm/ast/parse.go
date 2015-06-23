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

package asm

import (
	"errors"
	"github.com/katydid/katydid/expr/ast"
)

func (this *Rules) SetSpace(s *expr.Space) *Rules {
	this.Final = s
	return this
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
	case *Final:
		return newRules(&Rule{Final: rule}), nil
	case *Transition:
		return newRules(&Rule{Transition: rule}), nil
	case *FunctionDecl:
		return newRules(&Rule{FunctionDecl: rule}), nil
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
		rules.Rules = append(rules.Rules, &Rule{Root: rule})
		return rules, nil
	case *Init:
		for _, r := range rules.Rules {
			if r.Init != nil {
				return nil, errors.New("only one init is allowed")
			}
		}
		rules.Rules = append(rules.Rules, &Rule{Init: rule})
		return rules, nil
	case *Final:
		rules.Rules = append(rules.Rules, &Rule{Final: rule})
		return rules, nil
	case *Transition:
		rules.Rules = append(rules.Rules, &Rule{Transition: rule})
		return rules, nil
	case *FunctionDecl:
		rules.Rules = append(rules.Rules, &Rule{FunctionDecl: rule})
		return rules, nil
	}
	panic("unreachable")
}
