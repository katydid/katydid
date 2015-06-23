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
		"[]int64":    14,
		"[]int32":    15,
		"[]uint64":   16,
		"[]uint32":   17,
		"[]double":   18,
		"[]float":    19,
		"[]string":   20,
		"[][]byte":   21,
		"int64_lit":  22,
		"int32_lit":  23,
		"uint64_lit": 24,
		"uint32_lit": 25,
		"double_lit": 26,
		"float_lit":  27,
		"string_lit": 28,
		"bytes_lit":  29,
		"bool_var":   30,
		"int64_var":  31,
		"int32_var":  32,
		"uint64_var": 33,
		"uint32_var": 34,
		"double_var": 35,
		"float_var":  36,
		"string_var": 37,
		"bytes_var":  38,
		"true":       39,
		"false":      40,
		"=":          41,
		"(":          42,
		")":          43,
		"{":          44,
		"}":          45,
		",":          46,
		"space":      47,
	},
}
