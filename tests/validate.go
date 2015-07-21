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
)

type Validator struct {
	Name        string
	Grammar     *relapse.Grammar
	Expected    bool
	Description string
	Scanners    map[string]NewScanner
}

var Validators = make(map[string]Validator)

func Validate(name string, grammar combinator.G, codecs Codecs, expected bool) {
	Validators[name] = Validator{
		Name:        name,
		Grammar:     grammar.Grammar(),
		Expected:    expected,
		Scanners:    codecs.Scanners,
		Description: codecs.Description,
	}
}
