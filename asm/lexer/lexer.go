package lexer

import (

	// "fmt"
	// "github.com/katydid/katydid/asm/util"

	"github.com/katydid/katydid/asm/token"
	"io/ioutil"
	"unicode/utf8"
)

const (
	NoState    = -1
	NumStates  = 257
	NumSymbols = 179
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
0: '('
1: ')'
2: '('
3: ')'
4: '('
5: '-'
6: ')'
7: '$'
8: '$'
9: '$'
10: '$'
11: '$'
12: '$'
13: '{'
14: ','
15: '}'
16: 'r'
17: 'o'
18: 'o'
19: 't'
20: '.'
21: 'i'
22: 'n'
23: 'i'
24: 't'
25: 'f'
26: 'i'
27: 'n'
28: 'a'
29: 'l'
30: 'f'
31: 'u'
32: 'n'
33: 'c'
34: '['
35: ']'
36: 'b'
37: 'o'
38: 'o'
39: 'l'
40: '['
41: ']'
42: 'i'
43: 'n'
44: 't'
45: '['
46: ']'
47: 'u'
48: 'i'
49: 'n'
50: 't'
51: '['
52: ']'
53: 'd'
54: 'o'
55: 'u'
56: 'b'
57: 'l'
58: 'e'
59: '['
60: ']'
61: 's'
62: 't'
63: 'r'
64: 'i'
65: 'n'
66: 'g'
67: '['
68: ']'
69: '['
70: ']'
71: 'b'
72: 'y'
73: 't'
74: 'e'
75: 't'
76: 'r'
77: 'u'
78: 'e'
79: 'f'
80: 'a'
81: 'l'
82: 's'
83: 'e'
84: '='
85: '('
86: ')'
87: '{'
88: '}'
89: ','
90: '/'
91: '/'
92: '\n'
93: '/'
94: '*'
95: '*'
96: '*'
97: '/'
98: ' '
99: '\t'
100: '\n'
101: '\r'
102: '0'
103: '0'
104: 'x'
105: 'X'
106: '-'
107: 'e'
108: 'E'
109: '+'
110: '-'
111: '.'
112: '.'
113: '.'
114: '_'
115: '_'
116: 'd'
117: 'o'
118: 'u'
119: 'b'
120: 'l'
121: 'e'
122: 'i'
123: 'n'
124: 't'
125: 'u'
126: 'i'
127: 'n'
128: 't'
129: '['
130: ']'
131: 'b'
132: 'y'
133: 't'
134: 'e'
135: 's'
136: 't'
137: 'r'
138: 'i'
139: 'n'
140: 'g'
141: 'b'
142: 'o'
143: 'o'
144: 'l'
145: '.'
146: '\'
147: 'U'
148: '\'
149: 'u'
150: '\'
151: 'x'
152: '\'
153: '`'
154: '`'
155: '\'
156: 'a'
157: 'b'
158: 'f'
159: 'n'
160: 'r'
161: 't'
162: 'v'
163: '\'
164: '''
165: '"'
166: '"'
167: '"'
168: '''
169: '''
170: '0'-'9'
171: '0'-'7'
172: '0'-'9'
173: 'A'-'F'
174: 'a'-'f'
175: '1'-'9'
176: 'A'-'Z'
177: 'a'-'z'
178: .

*/
