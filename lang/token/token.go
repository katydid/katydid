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
		"Empty",
		"EmptySet",
		"TreeNode",
		"LeafNode",
		"Concat",
		"Or",
		"And",
		"ZeroOrMore",
		"Reference",
		"Not",
		"[]bool",
		"[]int",
		"[]uint",
		"[]double",
		"[]string",
		"[][]byte",
		"int_lit",
		"uint_lit",
		"double_lit",
		"string_lit",
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
		"space",
	},

	idMap: map[string]Type{
		"INVALID":    0,
		"$":          1,
		"id":         2,
		"Empty":      3,
		"EmptySet":   4,
		"TreeNode":   5,
		"LeafNode":   6,
		"Concat":     7,
		"Or":         8,
		"And":        9,
		"ZeroOrMore": 10,
		"Reference":  11,
		"Not":        12,
		"[]bool":     13,
		"[]int":      14,
		"[]uint":     15,
		"[]double":   16,
		"[]string":   17,
		"[][]byte":   18,
		"int_lit":    19,
		"uint_lit":   20,
		"double_lit": 21,
		"string_lit": 22,
		"bytes_lit":  23,
		"bool_var":   24,
		"int_var":    25,
		"uint_var":   26,
		"double_var": 27,
		"string_var": 28,
		"bytes_var":  29,
		"true":       30,
		"false":      31,
		"=":          32,
		"(":          33,
		")":          34,
		"{":          35,
		"}":          36,
		",":          37,
		"space":      38,
	},
}
