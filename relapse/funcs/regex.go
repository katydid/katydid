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
	"fmt"
	"regexp"
	"strings"
)

//Regex returns a new regex function given the first parameter as the expression string that needs to compiled and the second as the regex that should be matched.
func Regex(expr ConstString, input String) (Bool, error) {
	if expr.HasVariable() {
		return nil, fmt.Errorf("regex requires a constant expression as its first parameter, but it has a variable parameter")
	}
	e, err := expr.Eval()
	if err != nil {
		return nil, err
	}
	r, err := regexp.Compile(e)
	if err != nil {
		return nil, err
	}
	return TrimBool(&regex{
		r:           r,
		S:           input,
		expr:        e,
		hash:        Hash("regex", expr, input),
		hasVariable: input.HasVariable(),
	}), nil
}

type regex struct {
	r           *regexp.Regexp
	expr        string
	S           String
	hash        uint64
	hasVariable bool
}

func (this *regex) HasVariable() bool {
	return this.hasVariable
}

func (this *regex) String() string {
	return "regex(" + this.expr + "," + this.S.String() + ")"
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
		if c := strings.Compare(this.expr, other.expr); c != 0 {
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
