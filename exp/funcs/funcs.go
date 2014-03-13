package funcs

import (
	"code.google.com/p/go.text/unicode/norm"
	"encoding/binary"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

type decString struct {
	//TODO: add cache
}

func (this *decString) Eval(buf []byte) string {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	strHeader := reflect.StringHeader{header.Data, header.Len}
	return *(*string)(unsafe.Pointer(&strHeader))
}

func init() {
	Register("decString", new(decString))
}

type decInt64 struct {
	//TODO: add cache
}

func (this *decInt64) Eval(buf []byte) int64 {
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return int64(v)
}

func init() {
	Register("decInt64", new(decInt64))
}

type contains struct {
	V1 String
	V2 String
}

func (this *contains) Eval(buf []byte) bool {
	v1 := this.V1.Eval(buf)
	v2 := this.V2.Eval(buf)
	return strings.Contains(v1, v2)
}

func init() {
	Register("contains", new(contains))
}

type int64Eq struct {
	V1 Int64
	V2 Int64
}

func (this *int64Eq) Eval(buf []byte) bool {
	return this.V1.Eval(buf) == this.V2.Eval(buf)
}

func init() {
	Register("eq", new(int64Eq))
}

type stringEq struct {
	V1 String
	V2 String
}

func (this *stringEq) Eval(buf []byte) bool {
	return this.V1.Eval(buf) == this.V2.Eval(buf)
}

func init() {
	Register("eq", new(stringEq))
}

type int64Ge struct {
	V1 Int64
	V2 Int64
}

func (this *int64Ge) Eval(buf []byte) bool {
	return this.V1.Eval(buf) >= this.V2.Eval(buf)
}

func init() {
	Register("ge", new(int64Ge))
}

type nfkc struct {
	V1 String
}

func (this *nfkc) Eval(buf []byte) string {
	return norm.NFKC.String(this.V1.Eval(buf))
}

func init() {
	Register("nfkc", new(nfkc))
}

type not struct {
	V1 Bool
}

func (this *not) Eval(buf []byte) bool {
	return !this.V1.Eval(buf)
}

func init() {
	Register("not", new(not))
}

type exists struct {
}

func (this *exists) Eval(buf []byte) bool {
	return true
}

func init() {
	Register("exists", new(exists))
}

type lenString struct {
	V1 String
}

func (this *lenString) Eval(buf []byte) int64 {
	return int64(len(this.V1.Eval(buf)))
}

func init() {
	Register("length", new(lenString))
}

type decBool struct {
	//TODO: add cache
}

func (this *decBool) Eval(buf []byte) bool {
	v, n := binary.Uvarint(buf)
	if n <= 0 {
		panic(fmt.Sprintf("decodeVarint n = %d", n))
	}
	return v == 1
}

func init() {
	Register("decBool", new(decBool))
}
