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
	"strings"
)

//Regex returns a new regex function given the first parameter as the expression string that needs to compiled and the second as the regex that should be matched.
func Regex(expr ConstString, input String) Bool {
	h := uint64(17)
	h = 31*h + 41
	h = 31*h + expr.Hash()
	h = 31*h + input.Hash()
	return &regex{Expr: expr, S: input, hash: h}
}

type regex struct {
	r    *regexp.Regexp
	Expr ConstString
	S    String
	hash uint64
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

func (this *regex) String() string {
	return "regex(" + sjoin(this.Expr, this.S) + ")"
}

func (this *regex) Eval() (bool, error) {
	s, err := this.S.Eval()
	if err != nil {
		return false, err
	}
	return this.r.MatchString(s), nil
}

func (this *regex) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*regex); ok {
		if c := this.Expr.Compare(other.Expr); c != 0 {
			return c
		}
		if c := this.S.Compare(other.S); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *regex) Hash() uint64 {
	return this.hash
}

func init() {
	Register("regex", Regex)
}
