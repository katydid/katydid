package funcs

import (
	descriptor "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	"testing"
)

type which struct {
	exp string
}

func (this which) test(t *testing.T, name string, params ...descriptor.FieldDescriptorProto_Type) {
	uniq, err := funcsMap.which(name, params...)
	if err != nil {
		panic(err)
	}
	if uniq != this.exp {
		t.Fatalf("expected %v got %v", this.exp, uniq)
	}
}

func TestWhichStringEq(t *testing.T) {
	which{"stringEq"}.test(t, "eq", descriptor.FieldDescriptorProto_TYPE_STRING, descriptor.FieldDescriptorProto_TYPE_STRING)
}

func TestWhichInt64Eq(t *testing.T) {
	which{"int64Eq"}.test(t, "eq", descriptor.FieldDescriptorProto_TYPE_INT64, descriptor.FieldDescriptorProto_TYPE_INT64)
}

func TestWhichInt64Ge(t *testing.T) {
	which{"int64Ge"}.test(t, "ge", descriptor.FieldDescriptorProto_TYPE_INT64, descriptor.FieldDescriptorProto_TYPE_INT64)
}
