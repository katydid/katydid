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
	NumStates  = 276
	NumSymbols = 209
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
16: '<'
17: 'e'
18: 'm'
19: 'p'
20: 't'
21: 'y'
22: '>'
23: '<'
24: 'e'
25: 'm'
26: 'p'
27: 't'
28: 'y'
29: 's'
30: 'e'
31: 't'
32: '>'
33: '['
34: ']'
35: 'b'
36: 'o'
37: 'o'
38: 'l'
39: '['
40: ']'
41: 'i'
42: 'n'
43: 't'
44: '['
45: ']'
46: 'u'
47: 'i'
48: 'n'
49: 't'
50: '['
51: ']'
52: 'd'
53: 'o'
54: 'u'
55: 'b'
56: 'l'
57: 'e'
58: '['
59: ']'
60: 's'
61: 't'
62: 'r'
63: 'i'
64: 'n'
65: 'g'
66: '['
67: ']'
68: '['
69: ']'
70: 'b'
71: 'y'
72: 't'
73: 'e'
74: 't'
75: 'r'
76: 'u'
77: 'e'
78: 'f'
79: 'a'
80: 'l'
81: 's'
82: 'e'
83: '='
84: '('
85: ')'
86: '{'
87: '}'
88: ','
89: ';'
90: '#'
91: '&'
92: '|'
93: '['
94: ']'
95: ':'
96: '!'
97: '*'
98: '_'
99: '~'
100: '.'
101: '@'
102: '-'
103: '>'
104: '='
105: '='
106: '<'
107: '>'
108: '<'
109: '='
110: '>'
111: '='
112: '~'
113: '='
114: '*'
115: '='
116: '^'
117: '='
118: '$'
119: '='
120: '/'
121: '/'
122: '\n'
123: '/'
124: '*'
125: '*'
126: '*'
127: '/'
128: ' '
129: '\t'
130: '\n'
131: '\r'
132: '0'
133: '0'
134: 'x'
135: 'X'
136: '-'
137: 'e'
138: 'E'
139: '+'
140: '-'
141: '.'
142: '.'
143: '.'
144: '_'
145: '_'
146: 'd'
147: 'o'
148: 'u'
149: 'b'
150: 'l'
151: 'e'
152: 'i'
153: 'n'
154: 't'
155: 'u'
156: 'i'
157: 'n'
158: 't'
159: '['
160: ']'
161: 'b'
162: 'y'
163: 't'
164: 'e'
165: 's'
166: 't'
167: 'r'
168: 'i'
169: 'n'
170: 'g'
171: 'b'
172: 'o'
173: 'o'
174: 'l'
175: '.'
176: '\'
177: 'U'
178: '\'
179: 'u'
180: '\'
181: 'x'
182: '\'
183: '`'
184: '`'
185: '\'
186: 'a'
187: 'b'
188: 'f'
189: 'n'
190: 'r'
191: 't'
192: 'v'
193: '\'
194: '''
195: '"'
196: '"'
197: '"'
198: '''
199: '''
200: '0'-'9'
201: '0'-'7'
202: '0'-'9'
203: 'A'-'F'
204: 'a'-'f'
205: '1'-'9'
206: 'A'-'Z'
207: 'a'-'z'
208: .

*/
