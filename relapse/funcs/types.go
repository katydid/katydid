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

type Stringer interface {
	String() string
}

type Hashable interface {
	Hash() uint64
}

type Comparable interface {
	Compare(Comparable) int
	Hashable
	Stringer
}

type Func interface {
	Comparable
	HasVariable() bool
}

//Double is an interface that represents a function that returns a double.
type Double interface {
	Func
	Eval() (float64, error)
}

//Int is an interface that represents a function that returns an int.
type Int interface {
	Func
	Eval() (int64, error)
}

//Uint is an interface that represents a function that returns a uint.
type Uint interface {
	Func
	Eval() (uint64, error)
}

//Bool is an interface that represents a function that returns a bool.
type Bool interface {
	Func
	Eval() (bool, error)
}

//String is an interface that represents a function that returns a string.
type String interface {
	Func
	Eval() (string, error)
}

//Bytes is an interface that represents a function that returns []byte.
type Bytes interface {
	Func
	Eval() ([]byte, error)
}

//Doubles is an interface that represents a function that returns a list of doubles.
type Doubles interface {
	Func
	Eval() ([]float64, error)
}

//Ints is an interface that represents a function that returns a list of ints.
type Ints interface {
	Func
	Eval() ([]int64, error)
}

//Uints is an interface that represents a function that returns a list of uints.
type Uints interface {
	Func
	Eval() ([]uint64, error)
}

//Bools is an interface that represents a function that returns a list of bools.
type Bools interface {
	Func
	Eval() ([]bool, error)
}

//Strings is an interface that represents a function that returns a list of strings.
type Strings interface {
	Func
	Eval() ([]string, error)
}

//ListOfBytes is an interface that represents a function that returns a list of []byte.
type ListOfBytes interface {
	Func
	Eval() ([][]byte, error)
}
