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

package prototests

import (
	"github.com/gogo/protobuf/proto"
	. "github.com/katydid/katydid/parser/debug"
)

var AContainer = &Container{
	Field1: proto.Int64(123),
}

func init() {
	f := float64(0.123)
	if err := proto.SetExtension(AContainer, E_FieldA, &f); err != nil {
		panic(err)
	}
	if err := proto.SetExtension(AContainer, E_FieldB, &Small{SmallField: proto.Int64(456)}); err != nil {
		panic(err)
	}
	if err := proto.SetExtension(AContainer, E_FieldC, &Big{BigField: proto.Int64(789)}); err != nil {
		panic(err)
	}
}

var AContainerOutput = Nodes{
	Field(`FieldA`, `0.123`),
	Nested(`FieldB`,
		Field(`SmallField`, `456`),
	),
	Nested(`FieldC`,
		Field(`BigField`, `789`),
	),
	Field(`Field1`, `123`),
}

var ABigContainer = &BigContainer{
	Field13: proto.Int64(987),
	M:       AContainer,
}

var ABigContainerOutput = Nodes{
	Nested(`M`, AContainerOutput...),
	Field(`Field13`, `987`),
}
