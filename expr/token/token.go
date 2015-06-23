package token

import (
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (this Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", this.Offset, this.Line, this.Column)
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (this TokenMap) Id(tok Type) string {
	if int(tok) < len(this.typeMap) {
		return this.typeMap[tok]
	}
	return "unknown"
}

func (this TokenMap) Type(tok string) Type {
	if typ, exist := this.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (this TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", this.Id(tok.Type), tok.Type, tok.Lit)
}

func (this TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", this.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"id",
		"[]bool",
		"[]int64",
		"[]int32",
		"[]uint64",
		"[]uint32",
		"[]double",
		"[]float",
		"[]string",
		"[][]byte",
		"int64_lit",
		"int32_lit",
		"uint64_lit",
		"uint32_lit",
		"double_lit",
		"float_lit",
		"string_lit",
		"bytes_lit",
		"bool_var",
		"int64_var",
		"int32_var",
		"uint64_var",
		"uint32_var",
		"double_var",
		"float_var",
		"string_var",
		"bytes_var",
		"true",
		"false",
		"=",
		"(",
		")",
		"{",
		"}",
		",",
		"space",
	},

	idMap: map[string]Type{
		"INVALID":    0,
		"$":          1,
		"id":         2,
		"[]bool":     3,
		"[]int64":    4,
		"[]int32":    5,
		"[]uint64":   6,
		"[]uint32":   7,
		"[]double":   8,
		"[]float":    9,
		"[]string":   10,
		"[][]byte":   11,
		"int64_lit":  12,
		"int32_lit":  13,
		"uint64_lit": 14,
		"uint32_lit": 15,
		"double_lit": 16,
		"float_lit":  17,
		"string_lit": 18,
		"bytes_lit":  19,
		"bool_var":   20,
		"int64_var":  21,
		"int32_var":  22,
		"uint64_var": 23,
		"uint32_var": 24,
		"double_var": 25,
		"float_var":  26,
		"string_var": 27,
		"bytes_var":  28,
		"true":       29,
		"false":      30,
		"=":          31,
		"(":          32,
		")":          33,
		"{":          34,
		"}":          35,
		",":          36,
		"space":      37,
	},
}
