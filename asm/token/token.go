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
		"root",
		"id",
		".",
		"init",
		"final",
		"func",
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
		"root":       2,
		"id":         3,
		".":          4,
		"init":       5,
		"final":      6,
		"func":       7,
		"[]bool":     8,
		"[]int64":    9,
		"[]int32":    10,
		"[]uint64":   11,
		"[]uint32":   12,
		"[]double":   13,
		"[]float":    14,
		"[]string":   15,
		"[][]byte":   16,
		"int64_lit":  17,
		"int32_lit":  18,
		"uint64_lit": 19,
		"uint32_lit": 20,
		"double_lit": 21,
		"float_lit":  22,
		"string_lit": 23,
		"bytes_lit":  24,
		"bool_var":   25,
		"int64_var":  26,
		"int32_var":  27,
		"uint64_var": 28,
		"uint32_var": 29,
		"double_var": 30,
		"float_var":  31,
		"string_var": 32,
		"bytes_var":  33,
		"true":       34,
		"false":      35,
		"=":          36,
		"(":          37,
		")":          38,
		"{":          39,
		"}":          40,
		",":          41,
		"space":      42,
	},
}
