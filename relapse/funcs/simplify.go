//  Copyright 2016 Walter Schulze
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
	"reflect"
)

type Hashable interface {
	Hash() uint64
}

//IsFalse returns whether a function is a false constant.
func IsFalse(fn Bool) bool {
	v, ok := fn.(*constBool)
	if !ok {
		return false
	}
	return v.v == false
}

//IsTrue returns whether a function is a true constant.
func IsTrue(fn Bool) bool {
	v, ok := fn.(*constBool)
	if !ok {
		return false
	}
	return v.v == true
}

//Equal returns whether two functions are equal.
func Equal(l, r Hashable) bool {
	hl := l.Hash()
	hr := r.Hash()
	if hl != hr {
		return false
	}
	return eq(l, r)
}

func eq(l, r interface{}) bool {
	le := reflect.ValueOf(l).Elem()
	lUniqName := le.Type().Name()
	re := reflect.ValueOf(r).Elem()
	rUniqName := re.Type().Name()
	if lUniqName != rUniqName {
		return false
	}
	if len(reverse(lUniqName)) == 0 {
		//TODO maybe this could be done better or we could just always convert functions to strings to compare
		return sprint(l) == sprint(r)
	}
	numFields := le.NumField()
	for i := 0; i < numFields; i++ {
		if _, ok := le.Field(i).Type().MethodByName("Eval"); !ok {
			continue
		}
		if !eq(le.Field(i).Interface(), re.Field(i).Interface()) {
			return false
		}
	}
	return true
}

//Compare compares two funtions.
func Compare(l, r Hashable) int {
	hl := l.Hash()
	hr := r.Hash()
	if hl != hr {
		if hl < hr {
			return -1
		}
		return 1
	}
	return cmp(l, r)
}

func cmp(l, r interface{}) int {
	le := reflect.ValueOf(l).Elem()
	lUniqName := le.Type().Name()
	re := reflect.ValueOf(r).Elem()
	rUniqName := re.Type().Name()
	if lUniqName != rUniqName {
		if lUniqName < rUniqName {
			return -1
		}
		return 1
	}
	if len(reverse(lUniqName)) == 0 {
		//TODO maybe this could be done better or we could just always convert functions to strings to compare
		sl := sprint(l)
		sr := sprint(r)
		if sl == sr {
			return 0
		}
		if sl < sr {
			return -1
		}
		return 1
	}
	numFields := le.NumField()
	for i := 0; i < numFields; i++ {
		if _, ok := le.Field(i).Type().MethodByName("Eval"); !ok {
			continue
		}
		c := cmp(le.Field(i).Interface(), re.Field(i).Interface())
		if c != 0 {
			return c
		}
	}
	return 0
}
