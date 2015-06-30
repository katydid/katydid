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

func NotBool(b Bool) Bool {
	return &not{b}
}

type not struct {
	V1 Bool
}

func (this *not) Eval() bool {
	return !this.V1.Eval()
}

func init() {
	Register("not", new(not))
}

func AndBool(a, b Bool) Bool {
	return &and{a, b}
}

type and struct {
	V1 Bool
	V2 Bool
}

func (this *and) Eval() bool {
	return this.V1.Eval() && this.V2.Eval()
}

func init() {
	Register("and", new(and))
}

func OrBool(a, b Bool) Bool {
	return &or{a, b}
}

type or struct {
	V1 Bool
	V2 Bool
}

func (this *or) Eval() bool {
	return this.V1.Eval() || this.V2.Eval()
}

func init() {
	Register("or", new(or))
}
