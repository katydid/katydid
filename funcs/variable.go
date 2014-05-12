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

type Variable interface {
	SetVariable(v []byte)
}

type variable struct {
	Value []byte
}

func (this *variable) Eval() []byte {
	return this.Value
}

func (this *variable) SetVariable(v []byte) {
	this.Value = v
}

func (this *variable) String() string {
	return "var"
}

func NewVariable() *variable {
	return &variable{}
}
