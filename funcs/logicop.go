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

func Not(b Bool) Bool {
	return &not{b}
}

type not struct {
	V1 Bool
}

func (this *not) Eval() (bool, error) {
	b, err := this.V1.Eval()
	if err != nil || !b {
		return true, nil
	}
	return !b, nil
}

func init() {
	Register("not", new(not))
}

func And(a, b Bool) Bool {
	return &and{a, b}
}

type and struct {
	V1 Bool
	V2 Bool
}

func (this *and) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil || !v1 {
		return false, err
	}
	return this.V2.Eval()
}

func init() {
	Register("and", new(and))
}

func Or(a, b Bool) Bool {
	return &or{a, b}
}

type or struct {
	V1 Bool
	V2 Bool
}

func (this *or) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err == nil && v1 {
		return true, nil
	}
	return this.V2.Eval()
}

func init() {
	Register("or", new(or))
}
