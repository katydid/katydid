//  Copyright 2016 Walter Schulze
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

func (this *Grammar) Equal(that *Grammar) bool {
	return deriveEqualGrammar(this, that)
}

func (this *Pattern) Compare(that *Pattern) int {
	return deriveComparePattern(this, that)
}

func (this *Pattern) Equal(that *Pattern) bool {
	return deriveEqualPattern(this, that)
}

func (p *Pattern) Hash() uint64 {
	return deriveHash(p)
}

func (this *NameExpr) Equal(that *NameExpr) bool {
	return deriveEqualNameExpr(this, that)
}

func (this *NameExpr) Compare(that *NameExpr) int {
	return deriveCompareNameExpr(this, that)
}
