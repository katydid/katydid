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

package funcs

import (
	"regexp"
)

func Regex(e ConstString, s String) Bool {
	return &regex{Expr: e, S: s}
}

type regex struct {
	r    *regexp.Regexp
	Expr ConstString
	S    String
}

func (this *regex) Init() error {
	e, err := this.Expr.Eval()
	if err != nil {
		return err
	}
	r, err := regexp.Compile(e)
	if err != nil {
		return err
	}
	this.r = r
	return nil
}

func (this *regex) Eval() (bool, error) {
	s, err := this.S.Eval()
	if err != nil {
		return false, err
	}
	return this.r.MatchString(s), nil
}

func init() {
	Register("regex", new(regex))
}
