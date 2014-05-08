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
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
)

func ProtoToEncodedType(fieldType descriptor.FieldDescriptorProto_Type) Type {
	return Type(fieldType)
}

func Decoded(Encoded Type) Type {
	switch Encoded {
	case ENCODED_DOUBLE:
		return SINGLE_DOUBLE
	case ENCODED_FLOAT:
		return SINGLE_FLOAT
	case ENCODED_INT64:
		return SINGLE_INT64
	case ENCODED_UINT64:
		return SINGLE_UINT64
	case ENCODED_INT32:
		return SINGLE_INT32
	case ENCODED_FIXED64:
		return SINGLE_UINT64
	case ENCODED_FIXED32:
		return SINGLE_UINT32
	case ENCODED_BOOL:
		return SINGLE_BOOL
	case ENCODED_STRING:
		return SINGLE_STRING
	case ENCODED_GROUP:
		return 0
	case ENCODED_MESSAGE:
		return 0
	case ENCODED_BYTES:
		return SINGLE_BYTES
	case ENCODED_UINT32:
		return SINGLE_UINT32
	case ENCODED_ENUM:
		return SINGLE_INT32
	case ENCODED_SFIXED32:
		return SINGLE_INT32
	case ENCODED_SFIXED64:
		return SINGLE_INT64
	case ENCODED_SINT32:
		return SINGLE_INT32
	case ENCODED_SINT64:
		return SINGLE_INT64
	}
	panic("unreachable")
}

func ProtoToType(field *descriptor.FieldDescriptorProto) Type {
	decType := Decoded(ProtoToEncodedType(field.GetType()))
	if !field.IsRepeated() {
		return decType
	}
	return Type(int32(decType) + 100)
}
