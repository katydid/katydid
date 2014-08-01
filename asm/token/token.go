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
		"if",
		"{",
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
		"=",
		"then",
		"else",
		"(",
		")",
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
		"if":         5,
		"{":          6,
		"[]bool":     7,
		"[]int64":    8,
		"[]int32":    9,
		"[]uint64":   10,
		"[]uint32":   11,
		"[]double":   12,
		"[]float":    13,
		"[]string":   14,
		"[][]byte":   15,
		"int64_lit":  16,
		"int32_lit":  17,
		"uint64_lit": 18,
		"uint32_lit": 19,
		"double_lit": 20,
		"float_lit":  21,
		"string_lit": 22,
		"bytes_lit":  23,
		"true":       24,
		"false":      25,
		"=":          26,
		"then":       27,
		"else":       28,
		"(":          29,
		")":          30,
		"}":          31,
		",":          32,
		"space":      33,
	},
}
