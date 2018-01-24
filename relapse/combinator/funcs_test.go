//  Copyright 2018 Walter Schulze
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

package combinator

import "testing"

func TestConst(t *testing.T) {
	v := Value(BoolConst(true))
	if !v.GetLeafNode().GetExpr().GetTerminal().GetBoolValue() {
		t.Fatalf("expected true")
	}

	vs := Value(BoolsConst([]bool{true}))
	if !vs.GetLeafNode().GetExpr().GetList().GetElems()[0].GetTerminal().GetBoolValue() {
		t.Fatalf("expected true")
	}
}
