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

type aConst interface {
	Stringer
	isConst()
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

func (this *constBool) isConst()   {}
func (this *constBytes) isConst()  {}
func (this *constDouble) isConst() {}
func (this *constInt) isConst()    {}
func (this *constString) isConst() {}
func (this *constUint) isConst()   {}

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
	return strings.Compare(this.String(), that.String())
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
	return strings.Compare(this.String(), that.String())
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
	return strings.Compare(this.String(), that.String())
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
	return strings.Compare(this.String(), that.String())
}

func (this *constString) Compare(that Comparable) int {
	if other, ok := that.(*constString); ok {
		return strings.Compare(this.v, other.v)
	}
	return strings.Compare(this.String(), that.String())
}

func (this *constBytes) Compare(that Comparable) int {
	if other, ok := that.(*constBytes); ok {
		return bytes.Compare(this.v, other.v)
	}
	return strings.Compare(this.String(), that.String())
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
	return strings.Compare(this.String(), that.String())
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
	return strings.Compare(this.String(), that.String())
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
	return strings.Compare(this.String(), that.String())
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
	return strings.Compare(this.String(), that.String())
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
	return strings.Compare(this.String(), that.String())
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
	return strings.Compare(this.String(), that.String())
}
