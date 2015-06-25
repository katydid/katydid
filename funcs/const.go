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
		return NewConstFloat64(v)
	case int64:
		return NewConstInt64(v)
	case uint64:
		return NewConstUint64(v)
	case bool:
		return NewConstBool(v)
	case string:
		return NewConstString(v)
	case []byte:
		return NewConstBytes(v)
	case []float64:
		return NewConstFloat64s(v)
	case []int64:
		return NewConstInt64s(v)
	case []uint64:
		return NewConstUint64s(v)
	case []bool:
		return NewConstBools(v)
	case []string:
		return NewConstStrings(v)
	case [][]byte:
		return NewConstListOfBytes(v)
	}
	panic("unreachable")
}
