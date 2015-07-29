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
		"AnyName",
		"Name",
		"string_lit",
		"AnyNameExcept",
		"NameChoice",
		"Empty",
		"EmptySet",
		"ZeroOrMore",
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
		"space",
	},

	idMap: map[string]Type{
		"INVALID":       0,
		"$":             1,
		"id":            2,
		"AnyName":       3,
		"Name":          4,
		"string_lit":    5,
		"AnyNameExcept": 6,
		"NameChoice":    7,
		"Empty":         8,
		"EmptySet":      9,
		"ZeroOrMore":    10,
		"[]bool":        11,
		"[]int":         12,
		"[]uint":        13,
		"[]double":      14,
		"[]string":      15,
		"[][]byte":      16,
		"int_lit":       17,
		"uint_lit":      18,
		"double_lit":    19,
		"bytes_lit":     20,
		"bool_var":      21,
		"int_var":       22,
		"uint_var":      23,
		"double_var":    24,
		"string_var":    25,
		"bytes_var":     26,
		"true":          27,
		"false":         28,
		"=":             29,
		"(":             30,
		")":             31,
		"{":             32,
		"}":             33,
		",":             34,
		";":             35,
		"#":             36,
		"&":             37,
		"|":             38,
		"[":             39,
		"]":             40,
		":":             41,
		"!":             42,
		"space":         43,
	},
}
