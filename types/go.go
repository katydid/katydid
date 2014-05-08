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

package types

import (
	"fmt"
	"reflect"
)

func FromGo(typ reflect.Type) Type {
	kind := typ.Kind()
	switch kind {
	case reflect.Float64:
		return SINGLE_DOUBLE
	case reflect.Float32:
		return SINGLE_FLOAT
	case reflect.Int64:
		return SINGLE_INT64
	case reflect.Uint64:
		return SINGLE_UINT64
	case reflect.Int32:
		return SINGLE_INT32
	case reflect.Uint32:
		return SINGLE_UINT32
	case reflect.Bool:
		return SINGLE_BOOL
	case reflect.String:
		return SINGLE_STRING
	case reflect.Slice:
		elemKind := typ.Elem().Kind()
		switch elemKind {
		case reflect.Uint8:
			return SINGLE_BYTES
		case reflect.Float64:
			return LIST_DOUBLE
		case reflect.Float32:
			return LIST_FLOAT
		case reflect.Int64:
			return LIST_INT64
		case reflect.Uint64:
			return LIST_UINT64
		case reflect.Int32:
			return LIST_INT32
		case reflect.Uint32:
			return LIST_UINT32
		case reflect.Bool:
			return LIST_BOOL
		case reflect.String:
			return LIST_STRING
		case reflect.Slice:
			k := typ.Elem().Elem().Kind()
			if k == reflect.Uint8 {
				return LIST_BYTES
			}
		}
	}
	panic(fmt.Sprintf("go Type %v unsupported", typ))
}
