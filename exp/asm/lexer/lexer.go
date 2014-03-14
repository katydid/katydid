package lexer

import (

	// "fmt"
	// "github.com/awalterschulze/katydid/exp/asm/util"

	"github.com/awalterschulze/katydid/exp/asm/token"
	"io/ioutil"
	"unicode/utf8"
)

const (
	NoState    = -1
	NumStates  = 106
	NumSymbols = 112
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (this *Lexer) Scan() (tok *token.Token) {

	// fmt.Printf("Lexer.Scan() pos=%d\n", this.pos)

	tok = new(token.Token)
	if this.pos >= len(this.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = this.pos, this.line, this.column
		return
	}
	start, end := this.pos, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {

		// fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", this.pos, this.line, this.column, state)

		if this.pos >= len(this.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(this.src[this.pos:])
			this.pos += size
		}
		switch rune1 {
		case '\n':
			this.line++
			this.column = 1
		case '\r':
			this.column = 1
		case '\t':
			this.column += 4
		default:
			this.column++
		}

		// Production start
		if rune1 != -1 {
			state = TransTab[state](rune1)
		} else {
			state = -1
		}
		// Production end

		// Debug start
		// nextState := -1
		// if rune1 != -1 {
		// 	nextState = TransTab[state](rune1)
		// }
		// fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		// fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", this.pos, size, start, end)
		// if nextState != -1 {
		// 	fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		// }
		// state = nextState
		// Debug end

		if state != -1 {
			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				// fmt.Printf("\t Accept(%s), %s(%d)\n", string(act), token.TokMap.Id(tok), tok)
				end = this.pos
			case ActTab[state].Ignore != "":
				// fmt.Printf("\t Ignore(%s)\n", string(act))
				start = this.pos
				state = 0
				if start >= len(this.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = this.pos
			}
		}
	}
	if end > start {
		this.pos = end
		tok.Lit = this.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset = start
	tok.Pos.Column = this.column
	tok.Pos.Line = this.line
	return
}

func (this *Lexer) Reset() {
	this.pos = 0
}

/*
Lexer symbols:
0: '-'
1: '_'
2: 'r'
3: 'o'
4: 'o'
5: 't'
6: '='
7: '.'
8: 'i'
9: 'f'
10: 't'
11: 'h'
12: 'e'
13: 'n'
14: 'e'
15: 'l'
16: 's'
17: 'e'
18: '{'
19: '}'
20: '('
21: ')'
22: '='
23: '='
24: '<'
25: '<'
26: '='
27: '>'
28: '>'
29: '='
30: '&'
31: '&'
32: '|'
33: '|'
34: 'o'
35: 'r'
36: 'a'
37: 'n'
38: 'd'
39: ','
40: 'i'
41: 'n'
42: 't'
43: '6'
44: '4'
45: 'u'
46: 'i'
47: 'n'
48: 't'
49: '6'
50: '4'
51: 't'
52: 'r'
53: 'u'
54: 'e'
55: 'T'
56: 'r'
57: 'u'
58: 'e'
59: 'f'
60: 'a'
61: 'l'
62: 's'
63: 'e'
64: 'F'
65: 'a'
66: 'l'
67: 's'
68: 'e'
69: '/'
70: '/'
71: '\n'
72: '/'
73: '*'
74: '*'
75: '*'
76: '/'
77: '_'
78: '\'
79: 'U'
80: '\'
81: 'u'
82: '\'
83: 'x'
84: '\'
85: 'a'
86: 'b'
87: 'f'
88: 'n'
89: 'r'
90: 't'
91: 'v'
92: '\'
93: '''
94: '"'
95: '\'
96: '`'
97: '`'
98: '"'
99: '"'
100: ' '
101: '\t'
102: '\n'
103: '\r'
104: '0'-'9'
105: 'A'-'Z'
106: 'a'-'z'
107: '0'-'9'
108: 'A'-'F'
109: 'a'-'f'
110: '0'-'7'
111: .

*/
