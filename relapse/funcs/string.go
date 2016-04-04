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
	"strings"
)

//ToLower returns a toLower function with the input function as its parameter.
func ToLower(s String) String {
	return &toLower{s}
}

type toLower struct {
	S String
}

func (this *toLower) Eval() (string, error) {
	s, err := this.S.Eval()
	if err != nil {
		return "", err
	}
	return strings.ToLower(s), nil
}

func init() {
	Register("toLower", new(toLower))
}

//ToUpper returns a toUpper function with the input function as its parameter.
func ToUpper(s String) String {
	return &toUpper{s}
}

type toUpper struct {
	S String
}

func (this *toUpper) Eval() (string, error) {
	s, err := this.S.Eval()
	if err != nil {
		return "", err
	}
	return strings.ToUpper(s), nil
}

func init() {
	Register("toUpper", new(toUpper))
}

//Contains returns a contains function with the two input function as its parameter.
func Contains(s, sub String) Bool {
	return &contains{s, sub}
}

type contains struct {
	S      String
	Substr String
}

func (this *contains) Eval() (bool, error) {
	s, err := this.S.Eval()
	if err != nil {
		return false, err
	}
	subStr, err := this.Substr.Eval()
	if err != nil {
		return false, err
	}
	return strings.Contains(s, subStr), nil
}

func init() {
	Register("contains", new(contains))
}

//EqualFold returns a eqFold function with the two input functions as its parameters.
func EqualFold(s, t String) Bool {
	return &equalFold{s, t}
}

type equalFold struct {
	S String
	T String
}

func (this *equalFold) Eval() (bool, error) {
	s, err := this.S.Eval()
	if err != nil {
		return false, err
	}
	t, err := this.T.Eval()
	if err != nil {
		return false, err
	}
	return strings.EqualFold(s, t), nil
}

func init() {
	Register("eqFold", new(equalFold))
}

//HasPrefix returns a hasPrefix function with the two input functions as its parameters.
func HasPrefix(a, b String) Bool {
	return &hasPrefix{a, b}
}

type hasPrefix struct {
	V1 String
	V2 String
}

func (this *hasPrefix) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, err
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, err
	}
	return strings.HasPrefix(v1, v2), nil
}

func init() {
	Register("hasPrefix", new(hasPrefix))
}

//HasSuffix returns a hasSuffix function with the two input functions as its parameters.
func HasSuffix(a, b String) Bool {
	return &hasSuffix{a, b}
}

type hasSuffix struct {
	V1 String
	V2 String
}

func (this *hasSuffix) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, err
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, err
	}
	return strings.HasSuffix(v1, v2), nil
}

func init() {
	Register("hasSuffix", new(hasSuffix))
}
