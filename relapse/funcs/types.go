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

//Double is an interface that represents a function that returns a double or an error.
type Double interface {
	Eval() (float64, error)
	Hash() uint64
}

//Int is an interface that represents a function that returns an int or an error.
type Int interface {
	Eval() (int64, error)
	Hash() uint64
}

//Uint is an interface that represents a function that returns a uint or an error.
type Uint interface {
	Eval() (uint64, error)
	Hash() uint64
}

//Bool is an interface that represents a function that returns a bool or an error.
type Bool interface {
	Eval() (bool, error)
	Hash() uint64
}

//String is an interface that represents a function that returns a string or an error.
type String interface {
	Eval() (string, error)
	Hash() uint64
}

//Bytes is an interface that represents a function that returns []byte or an error.
type Bytes interface {
	Eval() ([]byte, error)
	Hash() uint64
}

//Doubles is an interface that represents a function that returns a list of doubles or an error.
type Doubles interface {
	Eval() ([]float64, error)
	Hash() uint64
}

//Ints is an interface that represents a function that returns a list of ints or an error.
type Ints interface {
	Eval() ([]int64, error)
	Hash() uint64
}

//Uints is an interface that represents a function that returns a list of uints or an error.
type Uints interface {
	Eval() ([]uint64, error)
	Hash() uint64
}

//Bools is an interface that represents a function that returns a list of bools or an error.
type Bools interface {
	Eval() ([]bool, error)
	Hash() uint64
}

//Strings is an interface that represents a function that returns a list of strings or an error.
type Strings interface {
	Eval() ([]string, error)
	Hash() uint64
}

//ListOfBytes is an interface that represents a function that returns a list of []byte or an error.
type ListOfBytes interface {
	Eval() ([][]byte, error)
	Hash() uint64
}
