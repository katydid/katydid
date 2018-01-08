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

type Comparable interface {
	Hash() uint64
	Compare(Comparable) int
}

//Double is an interface that represents a function that returns a double or an error.
type Double interface {
	Comparable
	Eval() (float64, error)
}

//Int is an interface that represents a function that returns an int or an error.
type Int interface {
	Comparable
	Eval() (int64, error)
}

//Uint is an interface that represents a function that returns a uint or an error.
type Uint interface {
	Comparable
	Eval() (uint64, error)
}

//Bool is an interface that represents a function that returns a bool or an error.
type Bool interface {
	Comparable
	Eval() (bool, error)
}

//String is an interface that represents a function that returns a string or an error.
type String interface {
	Comparable
	Eval() (string, error)
}

//Bytes is an interface that represents a function that returns []byte or an error.
type Bytes interface {
	Comparable
	Eval() ([]byte, error)
}

//Doubles is an interface that represents a function that returns a list of doubles or an error.
type Doubles interface {
	Comparable
	Eval() ([]float64, error)
}

//Ints is an interface that represents a function that returns a list of ints or an error.
type Ints interface {
	Comparable
	Eval() ([]int64, error)
}

//Uints is an interface that represents a function that returns a list of uints or an error.
type Uints interface {
	Comparable
	Eval() ([]uint64, error)
}

//Bools is an interface that represents a function that returns a list of bools or an error.
type Bools interface {
	Comparable
	Eval() ([]bool, error)
}

//Strings is an interface that represents a function that returns a list of strings or an error.
type Strings interface {
	Comparable
	Eval() ([]string, error)
}

//ListOfBytes is an interface that represents a function that returns a list of []byte or an error.
type ListOfBytes interface {
	Comparable
	Eval() ([][]byte, error)
}
