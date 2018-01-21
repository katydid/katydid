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
	"fmt"
)

//ErrRangeCheck specifies a range check error that resulted while trying to access an element in a list.
type ErrRangeCheck struct {
	Index int
	Len   int
}

func (this ErrRangeCheck) Error() string {
	return fmt.Sprintf("relapse/funcs: range check error: trying to access index %d in list of length %d", this.Index, this.Len)
}

//NewRangeCheckErr returns a range check error
func NewRangeCheckErr(index, l int) ErrRangeCheck {
	return ErrRangeCheck{index, l}
}

//ErrRange specifies a range error that tried to slice a list where the first specified index is bigger than the last specified index.
type ErrRange struct {
	First int
	Last  int
}

func (this ErrRange) Error() string {
	return fmt.Sprintf("relapse/funcs: range error: first %d > last %d", this.First, this.Last)
}

//NewRangeErr returns a range error
func NewRangeErr(first, last int) ErrRange {
	return ErrRange{first, last}
}

type ErrContainsListNotConst struct{}

func (ErrContainsListNotConst) Error() string {
	return fmt.Sprintf("relapse/funcs: list passed to contains is not a constant")
}
