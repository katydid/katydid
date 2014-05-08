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
		"=",
		"id",
		".",
		"if",
		"then",
		"else",
		"{",
		"}",
		"(",
		")",
		"==",
		"<",
		"<=",
		">",
		">=",
		"&&",
		"||",
		"or",
		"and",
		",",
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
		"true",
		"false",
	},

	idMap: map[string]Type{
		"INVALID":    0,
		"$":          1,
		"root":       2,
		"=":          3,
		"id":         4,
		".":          5,
		"if":         6,
		"then":       7,
		"else":       8,
		"{":          9,
		"}":          10,
		"(":          11,
		")":          12,
		"==":         13,
		"<":          14,
		"<=":         15,
		">":          16,
		">=":         17,
		"&&":         18,
		"||":         19,
		"or":         20,
		"and":        21,
		",":          22,
		"[]bool":     23,
		"[]int64":    24,
		"[]int32":    25,
		"[]uint64":   26,
		"[]uint32":   27,
		"[]double":   28,
		"[]float":    29,
		"[]string":   30,
		"[][]byte":   31,
		"int64_lit":  32,
		"int32_lit":  33,
		"uint64_lit": 34,
		"uint32_lit": 35,
		"double_lit": 36,
		"float_lit":  37,
		"string_lit": 38,
		"bytes_lit":  39,
		"true":       40,
		"false":      41,
	},
}
