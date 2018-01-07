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

import "time"

//Now returns a new now function.
func Now() Int {
	return &now{}
}

type now struct{}

func (this *now) Eval() (int64, error) {
	return time.Now().UnixNano(), nil
}

func (this *now) Hash() uint64 {
	h := uint64(17)
	h = 31*h + 43
	return h
}

func (this *now) IsVariable() {}

func init() {
	Register("now", "now", Now)
}
