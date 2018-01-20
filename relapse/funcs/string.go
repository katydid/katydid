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
	h := uint64(17)
	h = 31*h + 47
	h = 31*h + s.Hash()
	return TrimString(&toLower{s, h})
}

type toLower struct {
	S    String
	hash uint64
}

func (this *toLower) Eval() (string, error) {
	s, err := this.S.Eval()
	if err != nil {
		return "", err
	}
	return strings.ToLower(s), nil
}

func (this *toLower) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*toLower); ok {
		if c := this.S.Compare(other.S); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *toLower) HasVariable() bool {
	return this.S.HasVariable()
}

func (this *toLower) String() string {
	return "toLower(" + this.S.String() + ")"
}

func (this *toLower) Hash() uint64 {
	return this.hash
}

func init() {
	Register("toLower", ToLower)
}

//ToUpper returns a toUpper function with the input function as its parameter.
func ToUpper(s String) String {
	h := uint64(17)
	h = 31*h + 53
	h = 31*h + s.Hash()
	return TrimString(&toUpper{s, h})
}

type toUpper struct {
	S    String
	hash uint64
}

func (this *toUpper) Eval() (string, error) {
	s, err := this.S.Eval()
	if err != nil {
		return "", err
	}
	return strings.ToUpper(s), nil
}

func (this *toUpper) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*toUpper); ok {
		if c := this.S.Compare(other.S); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *toUpper) HasVariable() bool {
	return this.S.HasVariable()
}

func (this *toUpper) String() string {
	return "toUpper(" + this.S.String() + ")"
}

func (this *toUpper) Hash() uint64 {
	return this.hash
}

func init() {
	Register("toUpper", ToUpper)
}

//Contains returns a contains function with the two input function as its parameter.
func Contains(s, sub String) Bool {
	h := uint64(17)
	h = 31*h + 59
	h = 31*h + s.Hash()
	h = 31*h + sub.Hash()
	return TrimBool(&contains{s, sub, h})
}

type contains struct {
	S      String
	Substr String
	hash   uint64
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

func (this *contains) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*contains); ok {
		if c := this.S.Compare(other.S); c != 0 {
			return c
		}
		if c := this.Substr.Compare(other.Substr); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *contains) HasVariable() bool {
	return this.S.HasVariable() || this.Substr.HasVariable()
}

func (this *contains) String() string {
	return "contains(" + sjoin(this.S, this.Substr) + ")"
}

func (this *contains) Hash() uint64 {
	return this.hash
}

func init() {
	Register("contains", Contains)
}

//EqualFold returns a eqFold function with the two input functions as its parameters.
func EqualFold(s, t String) Bool {
	h := uint64(17)
	h = 31*h + 71
	h = 31*h + s.Hash()
	h = 31*h + t.Hash()
	return TrimBool(&equalFold{s, t, h})
}

type equalFold struct {
	S    String
	T    String
	hash uint64
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

func (this *equalFold) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*equalFold); ok {
		if c := this.S.Compare(other.S); c != 0 {
			return c
		}
		if c := this.T.Compare(other.T); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *equalFold) HasVariable() bool {
	return this.S.HasVariable() || this.T.HasVariable()
}

func (this *equalFold) String() string {
	return "eqFold(" + sjoin(this.S, this.T) + ")"
}

func (this *equalFold) Hash() uint64 {
	return this.hash
}

func init() {
	Register("eqFold", EqualFold)
}

//HasPrefix returns a hasPrefix function with the two input functions as its parameters.
func HasPrefix(a, b String) Bool {
	h := uint64(17)
	h = 31*h + 73
	h = 31*h + a.Hash()
	h = 31*h + b.Hash()
	return TrimBool(&hasPrefix{a, b, h})
}

type hasPrefix struct {
	V1   String
	V2   String
	hash uint64
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

func (this *hasPrefix) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*hasPrefix); ok {
		if c := this.V1.Compare(other.V1); c != 0 {
			return c
		}
		if c := this.V2.Compare(other.V2); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *hasPrefix) HasVariable() bool {
	return this.V1.HasVariable() || this.V2.HasVariable()
}

func (this *hasPrefix) String() string {
	return "hasPrefix(" + sjoin(this.V1, this.V2) + ")"
}

func (this *hasPrefix) Hash() uint64 {
	return this.hash
}

func init() {
	Register("hasPrefix", HasPrefix)
}

//HasSuffix returns a hasSuffix function with the two input functions as its parameters.
func HasSuffix(a, b String) Bool {
	h := uint64(17)
	h = 31*h + 79
	h = 31*h + a.Hash()
	h = 31*h + b.Hash()
	return TrimBool(&hasSuffix{a, b, h})
}

type hasSuffix struct {
	V1   String
	V2   String
	hash uint64
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

func (this *hasSuffix) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*hasSuffix); ok {
		if c := this.V1.Compare(other.V1); c != 0 {
			return c
		}
		if c := this.V2.Compare(other.V2); c != 0 {
			return c
		}
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *hasSuffix) HasVariable() bool {
	return this.V1.HasVariable() || this.V2.HasVariable()
}

func (this *hasSuffix) String() string {
	return "hasSuffix(" + sjoin(this.V1, this.V2) + ")"
}

func (this *hasSuffix) Hash() uint64 {
	return this.hash
}

func init() {
	Register("hasSuffix", HasSuffix)
}
