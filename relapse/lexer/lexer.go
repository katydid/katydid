package lexer

import (

	// "fmt"
	// "github.com/katydid/katydid/relapse/util"

	"github.com/katydid/katydid/relapse/token"
	"io/ioutil"
	"unicode/utf8"
)

const (
	NoState    = -1
	NumStates  = 259
	NumSymbols = 183
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
16: 'E'
17: 'm'
18: 'p'
19: 't'
20: 'y'
21: 'E'
22: 'm'
23: 'p'
24: 't'
25: 'y'
26: 'S'
27: 'e'
28: 't'
29: '['
30: ']'
31: 'b'
32: 'o'
33: 'o'
34: 'l'
35: '['
36: ']'
37: 'i'
38: 'n'
39: 't'
40: '['
41: ']'
42: 'u'
43: 'i'
44: 'n'
45: 't'
46: '['
47: ']'
48: 'd'
49: 'o'
50: 'u'
51: 'b'
52: 'l'
53: 'e'
54: '['
55: ']'
56: 's'
57: 't'
58: 'r'
59: 'i'
60: 'n'
61: 'g'
62: '['
63: ']'
64: '['
65: ']'
66: 'b'
67: 'y'
68: 't'
69: 'e'
70: 't'
71: 'r'
72: 'u'
73: 'e'
74: 'f'
75: 'a'
76: 'l'
77: 's'
78: 'e'
79: '='
80: '('
81: ')'
82: '{'
83: '}'
84: ','
85: ';'
86: '#'
87: '&'
88: '|'
89: '['
90: ']'
91: ':'
92: '!'
93: '*'
94: '/'
95: '/'
96: '\n'
97: '/'
98: '*'
99: '*'
100: '*'
101: '/'
102: ' '
103: '\t'
104: '\n'
105: '\r'
106: '0'
107: '0'
108: 'x'
109: 'X'
110: '-'
111: 'e'
112: 'E'
113: '+'
114: '-'
115: '.'
116: '.'
117: '.'
118: '_'
119: '_'
120: 'd'
121: 'o'
122: 'u'
123: 'b'
124: 'l'
125: 'e'
126: 'i'
127: 'n'
128: 't'
129: 'u'
130: 'i'
131: 'n'
132: 't'
133: '['
134: ']'
135: 'b'
136: 'y'
137: 't'
138: 'e'
139: 's'
140: 't'
141: 'r'
142: 'i'
143: 'n'
144: 'g'
145: 'b'
146: 'o'
147: 'o'
148: 'l'
149: '.'
150: '\'
151: 'U'
152: '\'
153: 'u'
154: '\'
155: 'x'
156: '\'
157: '`'
158: '`'
159: '\'
160: 'a'
161: 'b'
162: 'f'
163: 'n'
164: 'r'
165: 't'
166: 'v'
167: '\'
168: '''
169: '"'
170: '"'
171: '"'
172: '''
173: '''
174: '0'-'9'
175: '0'-'7'
176: '0'-'9'
177: 'A'-'F'
178: 'a'-'f'
179: '1'-'9'
180: 'A'-'Z'
181: 'a'-'z'
182: .

*/
