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
	"code.google.com/p/gogoprotobuf/proto"
	"fmt"
	"github.com/awalterschulze/katydid/exp/asm/ast"
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
address number = number
address street = street
address _ = address
number street = numberAndStreet
number _ = number
street number = numberAndStreet
street _ = street
*/
func TestTable(t *testing.T) {
	transitions := []*ast.Transition{
		&ast.Transition{
			Src:   proto.String("start"),
			Input: proto.String("numberAndStreet"),
			Dst:   proto.String("accept"),
		},
		&ast.Transition{
			Src:   proto.String("start"),
			Input: proto.String("_"),
			Dst:   proto.String("start"),
		},
		&ast.Transition{
			Src:   proto.String("address"),
			Input: proto.String("number"),
			Dst:   proto.String("number"),
		},
		&ast.Transition{
			Src:   proto.String("address"),
			Input: proto.String("street"),
			Dst:   proto.String("street"),
		},
		&ast.Transition{
			Src:   proto.String("address"),
			Input: proto.String("_"),
			Dst:   proto.String("address"),
		},
		&ast.Transition{
			Src:   proto.String("number"),
			Input: proto.String("street"),
			Dst:   proto.String("numberAndStreet"),
		},
		&ast.Transition{
			Src:   proto.String("number"),
			Input: proto.String("_"),
			Dst:   proto.String("number"),
		},
		&ast.Transition{
			Src:   proto.String("street"),
			Input: proto.String("number"),
			Dst:   proto.String("numberAndStreet"),
		},
		&ast.Transition{
			Src:   proto.String("street"),
			Input: proto.String("_"),
			Dst:   proto.String("street"),
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
		Src:   proto.String("number"),
		Input: proto.String("accept"),
		Dst:   proto.String("number"),
	})
	for _, transition := range transitions {
		transTest(t, table, transition.GetSrc(), transition.GetInput(), transition.GetDst())
	}
}
