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

func IsList(typ Type) bool {
	switch typ {
	case LIST_DOUBLE:
	case LIST_FLOAT:
	case LIST_INT64:
	case LIST_UINT64:
	case LIST_INT32:
	case LIST_BOOL:
	case LIST_STRING:
	case LIST_BYTES:
	case LIST_UINT32:
	default:
		return false
	}
	return true
}

func ListToSingle(typ Type) Type {
	switch typ {
	case LIST_DOUBLE:
		return SINGLE_DOUBLE
	case LIST_FLOAT:
		return SINGLE_FLOAT
	case LIST_INT64:
		return SINGLE_INT64
	case LIST_UINT64:
		return SINGLE_UINT64
	case LIST_INT32:
		return SINGLE_INT32
	case LIST_BOOL:
		return SINGLE_BOOL
	case LIST_STRING:
		return SINGLE_STRING
	case LIST_BYTES:
		return SINGLE_BYTES
	case LIST_UINT32:
		return SINGLE_UINT32
	}
	panic("unreachable")
}

func SingleToList(typ Type) Type {
	switch typ {
	case SINGLE_DOUBLE:
		return LIST_DOUBLE
	case SINGLE_FLOAT:
		return LIST_FLOAT
	case SINGLE_INT64:
		return LIST_INT64
	case SINGLE_UINT64:
		return LIST_UINT64
	case SINGLE_INT32:
		return LIST_INT32
	case SINGLE_BOOL:
		return LIST_BOOL
	case SINGLE_STRING:
		return LIST_STRING
	case SINGLE_BYTES:
		return LIST_BYTES
	case SINGLE_UINT32:
		return LIST_UINT32
	}
	panic("unreachable")
}

func IsSingle(typ Type) bool {
	switch typ {
	case SINGLE_DOUBLE:
	case SINGLE_FLOAT:
	case SINGLE_INT64:
	case SINGLE_UINT64:
	case SINGLE_INT32:
	case SINGLE_BOOL:
	case SINGLE_STRING:
	case SINGLE_BYTES:
	case SINGLE_UINT32:
	default:
		return false
	}
	return true
}

func IsEncoded(typ Type) bool {
	switch typ {
	case ENCODED_DOUBLE:
	case ENCODED_FLOAT:
	case ENCODED_INT64:
	case ENCODED_UINT64:
	case ENCODED_INT32:
	case ENCODED_FIXED64:
	case ENCODED_FIXED32:
	case ENCODED_BOOL:
	case ENCODED_STRING:
	case ENCODED_GROUP:
	case ENCODED_MESSAGE:
	case ENCODED_BYTES:
	case ENCODED_UINT32:
	case ENCODED_ENUM:
	case ENCODED_SFIXED32:
	case ENCODED_SFIXED64:
	case ENCODED_SINT32:
	case ENCODED_SINT64:
	default:
		return false
	}
	return true
}
