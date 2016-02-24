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
	NumStates  = 294
	NumSymbols = 207
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
2: '0'
3: '-'
4: '('
5: ')'
6: '('
7: '-'
8: ')'
9: '-'
10: '$'
11: '$'
12: '$'
13: '$'
14: '$'
15: '$'
16: '{'
17: ','
18: '}'
19: '<'
20: 'e'
21: 'm'
22: 'p'
23: 't'
24: 'y'
25: '>'
26: '['
27: ']'
28: 'b'
29: 'o'
30: 'o'
31: 'l'
32: '['
33: ']'
34: 'i'
35: 'n'
36: 't'
37: '['
38: ']'
39: 'u'
40: 'i'
41: 'n'
42: 't'
43: '['
44: ']'
45: 'd'
46: 'o'
47: 'u'
48: 'b'
49: 'l'
50: 'e'
51: '['
52: ']'
53: 's'
54: 't'
55: 'r'
56: 'i'
57: 'n'
58: 'g'
59: '['
60: ']'
61: '['
62: ']'
63: 'b'
64: 'y'
65: 't'
66: 'e'
67: 't'
68: 'r'
69: 'u'
70: 'e'
71: 'f'
72: 'a'
73: 'l'
74: 's'
75: 'e'
76: '='
77: '('
78: ')'
79: '{'
80: '}'
81: ','
82: ';'
83: '#'
84: '&'
85: '|'
86: '['
87: ']'
88: ':'
89: '!'
90: '*'
91: '_'
92: '~'
93: '.'
94: '@'
95: '-'
96: '>'
97: '='
98: '='
99: '!'
100: '='
101: '<'
102: '>'
103: '<'
104: '='
105: '>'
106: '='
107: '~'
108: '='
109: '*'
110: '='
111: '^'
112: '='
113: '$'
114: '='
115: ':'
116: ':'
117: '?'
118: '/'
119: '/'
120: '\n'
121: '/'
122: '*'
123: '*'
124: '*'
125: '/'
126: ' '
127: '\t'
128: '\n'
129: '\r'
130: '0'
131: '0'
132: 'x'
133: 'X'
134: '-'
135: 'e'
136: 'E'
137: '+'
138: '-'
139: '.'
140: '.'
141: '.'
142: '_'
143: '_'
144: 'd'
145: 'o'
146: 'u'
147: 'b'
148: 'l'
149: 'e'
150: 'i'
151: 'n'
152: 't'
153: 'u'
154: 'i'
155: 'n'
156: 't'
157: '['
158: ']'
159: 'b'
160: 'y'
161: 't'
162: 'e'
163: 's'
164: 't'
165: 'r'
166: 'i'
167: 'n'
168: 'g'
169: 'b'
170: 'o'
171: 'o'
172: 'l'
173: '.'
174: '\'
175: 'U'
176: '\'
177: 'u'
178: '\'
179: 'x'
180: '\'
181: '`'
182: '`'
183: '\'
184: 'a'
185: 'b'
186: 'f'
187: 'n'
188: 'r'
189: 't'
190: 'v'
191: '\'
192: '''
193: '"'
194: '"'
195: '"'
196: '''
197: '''
198: '0'-'9'
199: '0'-'7'
200: '0'-'9'
201: 'A'-'F'
202: 'a'-'f'
203: '1'-'9'
204: 'A'-'Z'
205: 'a'-'z'
206: .

*/
