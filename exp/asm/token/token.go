
package token

import(
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const(
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line int
	Column int
}

func (this Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", this.Offset, this.Line, this.Column)
}

type TokenMap struct {
	typeMap  []string
	idMap map[string]Type
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
		"int64",
		"int_lit",
		"uint64",
		"string_lit",
		"true",
		"True",
		"false",
		"False",
	},

	idMap: map[string]Type {
		"INVALID": 0,
		"$": 1,
		"root": 2,
		"=": 3,
		"id": 4,
		".": 5,
		"if": 6,
		"then": 7,
		"else": 8,
		"{": 9,
		"}": 10,
		"(": 11,
		")": 12,
		"==": 13,
		"<": 14,
		"<=": 15,
		">": 16,
		">=": 17,
		"&&": 18,
		"||": 19,
		"or": 20,
		"and": 21,
		",": 22,
		"int64": 23,
		"int_lit": 24,
		"uint64": 25,
		"string_lit": 26,
		"true": 27,
		"True": 28,
		"false": 29,
		"False": 30,
	},
}

