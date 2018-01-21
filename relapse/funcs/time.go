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
	"time"
)

//Now returns a new now function.
func Now() Int {
	return &now{}
}

type now struct{}

func (this *now) Eval() (int64, error) {
	return time.Now().UnixNano(), nil
}

func (this *now) Hash() uint64 {
	return Hash("now")
}

func (this *now) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if _, ok := that.(*now); ok {
		return 0
	}
	return strings.Compare(this.String(), that.String())
}

func (this *now) String() string {
	return "now()"
}

func (this *now) HasVariable() bool { return true }

func init() {
	Register("now", Now)
}
