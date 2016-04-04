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

package ast

import (
	"fmt"
)

//HasVar returns whether there exists a variable somewhere in the expression tree.
//This function is executed recursively.
func (this *Expr) HasVar() bool {
	if this.Terminal != nil {
		return this.GetTerminal().Variable != nil
	}
	if this.List != nil {
		return this.GetList().HasVar()
	}
	if this.Function != nil {
		return this.GetFunction().HasVar()
	}
	if this.BuiltIn != nil {
		return true
	}
	panic(fmt.Sprintf("unknown expr %#v", this))
}

//HasVar returns whether there exists a variable somewhere in the typed list.
//This function is executed recursively.
func (this *List) HasVar() bool {
	for _, e := range this.GetElems() {
		if e.HasVar() {
			return true
		}
	}
	return false
}

//HasVar returns whether there exists a varaible somewhere in the function parameters.
//This function is executed recursively.
func (this *Function) HasVar() bool {
	for _, p := range this.GetParams() {
		if p.HasVar() {
			return true
		}
	}
	return false
}
