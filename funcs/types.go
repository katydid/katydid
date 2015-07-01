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

type Double interface {
	Eval() float64
}

type Int interface {
	Eval() int64
}

type Uint interface {
	Eval() uint64
}

type Bool interface {
	Eval() bool
}

type String interface {
	Eval() string
}

type Bytes interface {
	Eval() []byte
}

type Doubles interface {
	Eval() []float64
}

type Ints interface {
	Eval() []int64
}

type Uints interface {
	Eval() []uint64
}

type Bools interface {
	Eval() []bool
}

type Strings interface {
	Eval() []string
}

type ListOfBytes interface {
	Eval() [][]byte
}
