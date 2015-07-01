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

type Const interface {
	IsConst()
}

func NewConst(value interface{}) interface{} {
	switch v := value.(type) {
	case float64:
		return DoubleConst(v)
	case int64:
		return IntConst(v)
	case uint64:
		return UintConst(v)
	case bool:
		return BoolConst(v)
	case string:
		return StringConst(v)
	case []byte:
		return BytesConst(v)
	case []float64:
		return DoublesConst(v)
	case []int64:
		return IntsConst(v)
	case []uint64:
		return UintsConst(v)
	case []bool:
		return BoolsConst(v)
	case []string:
		return StringsConst(v)
	case [][]byte:
		return ListOfBytesConst(v)
	}
	panic("unreachable")
}
