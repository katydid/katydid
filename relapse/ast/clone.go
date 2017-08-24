//  Copyright 2015 Walter Schulze
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

//Clone returns a copy of the Grammar struct
func (this *Grammar) Clone() *Grammar {
	if this == nil {
		return nil
	}
	that := &Grammar{}
	deriveDeepCopyGrammar(that, this)
	return that
}

//Clone returns a copy of the RefLookup map
func (this RefLookup) Clone() RefLookup {
	that := make(RefLookup, len(this))
	for name := range this {
		that[name] = this[name].Clone()
	}
	return that
}

//Clone returns a copy of the Pattern struct
func (this *Pattern) Clone() *Pattern {
	if this == nil {
		return nil
	}
	that := &Pattern{}
	deriveDeepCopyPattern(that, this)
	return that
}

//Clone returns a copy of the Expr struct
func (this *Expr) Clone() *Expr {
	if this == nil {
		return nil
	}
	that := &Expr{}
	deriveDeepCopyExpr(that, this)
	return that
}
