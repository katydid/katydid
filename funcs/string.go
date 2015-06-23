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

type toLower struct {
	S String
}

func (this *toLower) Eval() string {
	return strings.ToLower(this.S.Eval())
}

func init() {
	Register("toLower", new(toLower))
}

type toUpper struct {
	S String
}

func (this *toUpper) Eval() string {
	return strings.ToUpper(this.S.Eval())
}

func init() {
	Register("toUpper", new(toUpper))
}

type contains struct {
	S      String
	Substr String
}

func (this *contains) Eval() bool {
	return strings.Contains(this.S.Eval(), this.Substr.Eval())
}

func init() {
	Register("contains", new(contains))
}

type equalFold struct {
	V1 String
	V2 String
}

func (this *equalFold) Eval() bool {
	v1 := this.V1.Eval()
	v2 := this.V2.Eval()
	return strings.EqualFold(v1, v2)
}

func init() {
	Register("eqFold", new(equalFold))
}

type hasPrefix struct {
	V1 String
	V2 String
}

func (this *hasPrefix) Eval() bool {
	v1 := this.V1.Eval()
	v2 := this.V2.Eval()
	return strings.HasPrefix(v1, v2)
}

func init() {
	Register("hasPrefix", new(hasPrefix))
}

type hasSuffix struct {
	V1 String
	V2 String
}

func (this *hasSuffix) Eval() bool {
	v1 := this.V1.Eval()
	v2 := this.V2.Eval()
	return strings.HasSuffix(v1, v2)
}

func init() {
	Register("hasSuffix", new(hasSuffix))
}
