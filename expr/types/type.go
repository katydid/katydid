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

//IsList returns whether the type is a list type
func IsList(typ Type) bool {
	switch typ {
	case LIST_DOUBLE:
	case LIST_INT:
	case LIST_UINT:
	case LIST_BOOL:
	case LIST_STRING:
	case LIST_BYTES:
	default:
		return false
	}
	return true
}

//ListToSingle returns the respective type as if the type was not a list type.
func ListToSingle(typ Type) Type {
	switch typ {
	case LIST_DOUBLE:
		return SINGLE_DOUBLE
	case LIST_INT:
		return SINGLE_INT
	case LIST_UINT:
		return SINGLE_UINT
	case LIST_BOOL:
		return SINGLE_BOOL
	case LIST_STRING:
		return SINGLE_STRING
	case LIST_BYTES:
		return SINGLE_BYTES
	}
	panic("unreachable")
}

//SingleToList returns the list type of the respective type.
func SingleToList(typ Type) Type {
	switch typ {
	case SINGLE_DOUBLE:
		return LIST_DOUBLE
	case SINGLE_INT:
		return LIST_INT
	case SINGLE_UINT:
		return LIST_UINT
	case SINGLE_BOOL:
		return LIST_BOOL
	case SINGLE_STRING:
		return LIST_STRING
	case SINGLE_BYTES:
		return LIST_BYTES
	}
	panic("unreachable")
}

//IsSingle returns whether the type is not a list type.
func IsSingle(typ Type) bool {
	switch typ {
	case SINGLE_DOUBLE:
	case SINGLE_INT:
	case SINGLE_UINT:
	case SINGLE_BOOL:
	case SINGLE_STRING:
	case SINGLE_BYTES:
	default:
		return false
	}
	return true
}
