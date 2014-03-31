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

type Float64 interface {
	Eval([]byte) float64
}

type Float32 interface {
	Eval([]byte) float32
}

type Int64 interface {
	Eval([]byte) int64
}

type Uint64 interface {
	Eval([]byte) uint64
}

type Int32 interface {
	Eval([]byte) int32
}

type Uint32 interface {
	Eval([]byte) uint32
}

type Bool interface {
	Eval([]byte) bool
}

type String interface {
	Eval([]byte) string
}

type Bytes interface {
	Eval([]byte) []byte
}
