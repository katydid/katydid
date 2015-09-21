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
		"start",
		"final",
		"internal",
		"call",
		"return",
		"id",
		"string_lit",
		"[]bool",
		"[]int",
		"[]uint",
		"[]double",
		"[]string",
		"[][]byte",
		"int_lit",
		"uint_lit",
		"double_lit",
		"bytes_lit",
		"bool_var",
		"int_var",
		"uint_var",
		"double_var",
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
		";",
		"#",
		"&",
		"|",
		"[",
		"]",
		":",
		"!",
		"*",
		"_",
		"~",
		"space",
	},

	idMap: map[string]Type{
		"INVALID":    0,
		"$":          1,
		"start":      2,
		"final":      3,
		"internal":   4,
		"call":       5,
		"return":     6,
		"id":         7,
		"string_lit": 8,
		"[]bool":     9,
		"[]int":      10,
		"[]uint":     11,
		"[]double":   12,
		"[]string":   13,
		"[][]byte":   14,
		"int_lit":    15,
		"uint_lit":   16,
		"double_lit": 17,
		"bytes_lit":  18,
		"bool_var":   19,
		"int_var":    20,
		"uint_var":   21,
		"double_var": 22,
		"string_var": 23,
		"bytes_var":  24,
		"true":       25,
		"false":      26,
		"=":          27,
		"(":          28,
		")":          29,
		"{":          30,
		"}":          31,
		",":          32,
		";":          33,
		"#":          34,
		"&":          35,
		"|":          36,
		"[":          37,
		"]":          38,
		":":          39,
		"!":          40,
		"*":          41,
		"_":          42,
		"~":          43,
		"space":      44,
	},
}
