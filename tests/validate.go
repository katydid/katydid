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

package tests

import (
	"github.com/katydid/katydid/relapse/ast"
	"github.com/katydid/katydid/relapse/combinator"
	"sort"
)

type Validator struct {
	Name        string
	CodecName   string
	Grammar     *relapse.Grammar
	Scanner     NewScanner
	Expected    bool
	Description string
}

var Validators = make(map[string]map[string]Validator)

type validatorList []interface{}

func (this validatorList) Less(i, j int) bool {
	return this[i].(Validator).Name < this[j].(Validator).Name
}

func (this validatorList) Len() int {
	return len(this)
}

func (this validatorList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func ValidatorList() []interface{} {
	vs := make(validatorList, 0, len(Validators)*3)
	for name, cs := range Validators {
		for cname, _ := range cs {
			vs = append(vs, Validators[name][cname])
		}
	}
	sort.Sort(vs)
	return vs
}

func Validate(name string, grammar combinator.G, codecs Codecs, expected bool) {
	if _, ok := Validators[name]; !ok {
		Validators[name] = make(map[string]Validator)
	}
	for codecName, s := range codecs.Scanners {
		Validators[name][codecName] = Validator{
			Name:        name,
			CodecName:   codecName,
			Grammar:     grammar.Grammar(),
			Scanner:     s,
			Expected:    expected,
			Description: codecs.Description,
		}
	}
}
