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

package debug

import (
	"github.com/gogo/protobuf/proto"
)

//Input is a sample instance of the Debug struct.
var Input = &Debug{
	A: int64(1),
	B: []string{"b2", "b3"},
	C: &Debug{
		A: int64(2),
		D: proto.Int32(3),
		E: []*Debug{
			&Debug{
				B: []string{"b4"},
			},
			&Debug{
				B: []string{"b5"},
			},
		},
	},
	D: proto.Int32(4),
	F: []uint32{5},
}

//Output is a sample instance of Nodes that repesents the Input variable after it has been parsed by Walk.
var Output = Nodes{
	Field(`A`, `1`),
	Nested(`B`,
		Field(`0`, `b2`),
		Field(`1`, `b3`),
	),
	Nested(`C`,
		Field(`A`, `2`),
		Field(`D`, `3`),
		Nested(`E`,
			Nested(`0`,
				Field(`A`, `0`),
				Nested(`B`,
					Field(`0`, `b4`),
				),
			),
			Nested(`1`,
				Field(`A`, `0`),
				Nested(`B`,
					Field(`0`, `b5`),
				),
			),
		),
	),
	Field(`D`, `4`),
	Nested(`F`,
		Field(`0`, `5`),
	),
}

//Field is a helper function for creating a Node with a label and one child label.
//This is how a field with a value is typically represented.
func Field(name string, value string) Node {
	return Node{
		Label: name,
		Children: Nodes{
			Node{
				Label: value,
			},
		},
	}
}

//Nested is a helper function for creating a Node.
func Nested(name string, fs ...Node) Node {
	return Node{
		Label:    name,
		Children: Nodes(fs),
	}
}
