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

import (
	"bytes"
	"strings"
)

//Const is an interface that when implemented implies that the function is actually a constant value.
type Const interface {
	IsConst()
}

//NewConst returns the appropriate constant function given a native go type or list of native go types.
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

func (this *constDouble) Compare(that Comparable) int {
	if other, ok := that.(*constDouble); ok {
		if this.v != other.v {
			if this.v < other.v {
				return -1
			}
			return 1
		}
		return 0
	}
	return strings.Compare("constDouble", nameOfStruct(that))
}

func (this *constInt) Compare(that Comparable) int {
	if other, ok := that.(*constInt); ok {
		if this.v != other.v {
			if this.v < other.v {
				return -1
			}
			return 1
		}
		return 0
	}
	return strings.Compare("constInt", nameOfStruct(that))
}

func (this *constUint) Compare(that Comparable) int {
	if other, ok := that.(*constUint); ok {
		if this.v != other.v {
			if this.v < other.v {
				return -1
			}
			return 1
		}
		return 0
	}
	return strings.Compare("constUint", nameOfStruct(that))
}

func (this *constBool) Compare(that Comparable) int {
	if other, ok := that.(*constBool); ok {
		if this.v != other.v {
			if !this.v {
				return -1
			}
			return 1
		}
		return 0
	}
	return strings.Compare("constBool", nameOfStruct(that))
}

func (this *constString) Compare(that Comparable) int {
	if other, ok := that.(*constString); ok {
		return strings.Compare(this.v, other.v)
	}
	return strings.Compare("constString", nameOfStruct(that))
}

func (this *constBytes) Compare(that Comparable) int {
	if other, ok := that.(*constBytes); ok {
		return bytes.Compare(this.v, other.v)
	}
	return strings.Compare("constBytes", nameOfStruct(that))
}

func (this *constDoubles) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*constDoubles); ok {
		if len(this.v) != len(other.v) {
			if len(this.v) < len(other.v) {
				return -1
			}
			return 1
		}
		for i := range this.v {
			if this.v[i] != other.v[i] {
				if this.v[i] < other.v[i] {
					return -1
				}
				return 1
			}
		}
		return 0
	}
	return strings.Compare("constDoubles", nameOfStruct(that))
}

func (this *constInts) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*constInts); ok {
		if len(this.v) != len(other.v) {
			if len(this.v) < len(other.v) {
				return -1
			}
			return 1
		}
		for i := range this.v {
			if this.v[i] != other.v[i] {
				if this.v[i] < other.v[i] {
					return -1
				}
				return 1
			}
		}
		return 0
	}
	return strings.Compare("constInts", nameOfStruct(that))
}

func (this *constUints) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*constUints); ok {
		if len(this.v) != len(other.v) {
			if len(this.v) < len(other.v) {
				return -1
			}
			return 1
		}
		for i := range this.v {
			if this.v[i] != other.v[i] {
				if this.v[i] < other.v[i] {
					return -1
				}
				return 1
			}
		}
		return 0
	}
	return strings.Compare("constUints", nameOfStruct(that))
}

func (this *constBools) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*constBools); ok {
		if len(this.v) != len(other.v) {
			if len(this.v) < len(other.v) {
				return -1
			}
			return 1
		}
		for i := range this.v {
			if this.v[i] != other.v[i] {
				if !this.v[i] {
					return -1
				}
				return 1
			}
		}
		return 0
	}
	return strings.Compare("constBools", nameOfStruct(that))
}

func (this *constStrings) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*constStrings); ok {
		if len(this.v) != len(other.v) {
			if len(this.v) < len(other.v) {
				return -1
			}
			return 1
		}
		for i := range this.v {
			if c := strings.Compare(this.v[i], other.v[i]); c != 0 {
				return c
			}
		}
		return 0
	}
	return strings.Compare("constStrings", nameOfStruct(that))
}

func (this *constListOfBytes) Compare(that Comparable) int {
	if this.Hash() != that.Hash() {
		if this.Hash() < that.Hash() {
			return -1
		}
		return 1
	}
	if other, ok := that.(*constListOfBytes); ok {
		if len(this.v) != len(other.v) {
			if len(this.v) < len(other.v) {
				return -1
			}
			return 1
		}
		for i := range this.v {
			if c := bytes.Compare(this.v[i], other.v[i]); c != 0 {
				return c
			}
		}
		return 0
	}
	return strings.Compare("constListOfBytes", nameOfStruct(that))
}
