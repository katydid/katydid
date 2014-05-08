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

package table

import (
	"fmt"
	"github.com/awalterschulze/katydid/asm/ast"
	"testing"
)

func transTest(t *testing.T, table Table, src string, input string, dst string) {
	srcState, err := table.NameToState(src)
	if err != nil {
		panic(err)
	}
	inputState, err := table.NameToState(input)
	if err != nil {
		panic(err)
	}
	dstState, err := table.NameToState(dst)
	if err != nil {
		panic(err)
	}
	if d, err := table.Trans(srcState, inputState); err != nil {
		panic(err)
	} else if d != dstState {
		t.Fatalf("%d %d should go to %d, but got %d", srcState, inputState, dstState, d)
	}
}

/*
start numberAndStreet = accept
start _ = start
accept _ accept
address number = number
address street = street
address _ = address
number street = numberAndStreet
number _ = number
street number = numberAndStreet
street _ = street
numberAndStreet _ = numberAndStreet
*/
func TestTable(t *testing.T) {
	transitions := []*ast.Transition{
		{
			Src:   "start",
			Input: "numberAndStreet",
			Dst:   "accept",
		},
		{
			Src:   "start",
			Input: "_",
			Dst:   "start",
		},
		{
			Src:   "accept",
			Input: "_",
			Dst:   "accept",
		},
		{
			Src:   "address",
			Input: "number",
			Dst:   "number",
		},
		{
			Src:   "address",
			Input: "street",
			Dst:   "street",
		},
		{
			Src:   "address",
			Input: "_",
			Dst:   "address",
		},
		{
			Src:   "number",
			Input: "street",
			Dst:   "numberAndStreet",
		},
		{
			Src:   "number",
			Input: "_",
			Dst:   "number",
		},
		{
			Src:   "street",
			Input: "number",
			Dst:   "numberAndStreet",
		},
		{
			Src:   "street",
			Input: "_",
			Dst:   "street",
		},
		{
			Src:   "numberAndStreet",
			Input: "_",
			Dst:   "numberAndStreet",
		},
	}
	table := New(transitions, nil)
	fmt.Printf("%v\n", table.Dot())
	if s, err := table.NameToState("_"); err != nil {
		panic(err)
	} else if s != 0 {
		t.Fatalf("'_' should be state 0")
	}
	transitions = append(transitions, &ast.Transition{
		Src:   "number",
		Input: "accept",
		Dst:   "number",
	})
	for _, transition := range transitions {
		transTest(t, table, transition.GetSrc(), transition.GetInput(), transition.GetDst())
	}
	acc, err := table.NameToState("accept")
	if err != nil {
		panic(err)
	}
	if !table.NoEscapeFrom(acc) {
		t.Fatalf("excepted no escape from accept state")
	}
	numberAndStreet, err := table.NameToState("numberAndStreet")
	if err != nil {
		panic(err)
	}
	if !table.NoEscapeFrom(numberAndStreet) {
		t.Fatalf("excepted no escape from numberAndStreet state")
	}
	street, err := table.NameToState("street")
	if err != nil {
		panic(err)
	}
	if table.NoEscapeFrom(street) {
		t.Fatalf("expected escape from street to be possible")
	}
}
