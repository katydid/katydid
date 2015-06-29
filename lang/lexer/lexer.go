package lexer

import (

	// "fmt"
	// "github.com/katydid/katydid/lang/util"

	"github.com/katydid/katydid/lang/token"
	"io/ioutil"
	"unicode/utf8"
)

const (
	NoState    = -1
	NumStates  = 320
	NumSymbols = 257
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
16: 'A'
17: 'n'
18: 'y'
19: 'N'
20: 'a'
21: 'm'
22: 'e'
23: 'N'
24: 'a'
25: 'm'
26: 'e'
27: 'A'
28: 'n'
29: 'y'
30: 'N'
31: 'a'
32: 'm'
33: 'e'
34: 'E'
35: 'x'
36: 'c'
37: 'e'
38: 'p'
39: 't'
40: 'N'
41: 'a'
42: 'm'
43: 'e'
44: 'C'
45: 'h'
46: 'o'
47: 'i'
48: 'c'
49: 'e'
50: 'E'
51: 'm'
52: 'p'
53: 't'
54: 'y'
55: 'E'
56: 'm'
57: 'p'
58: 't'
59: 'y'
60: 'S'
61: 'e'
62: 't'
63: 'T'
64: 'r'
65: 'e'
66: 'e'
67: 'N'
68: 'o'
69: 'd'
70: 'e'
71: 'L'
72: 'e'
73: 'a'
74: 'f'
75: 'N'
76: 'o'
77: 'd'
78: 'e'
79: 'C'
80: 'o'
81: 'n'
82: 'c'
83: 'a'
84: 't'
85: 'O'
86: 'r'
87: 'A'
88: 'n'
89: 'd'
90: 'Z'
91: 'e'
92: 'r'
93: 'o'
94: 'O'
95: 'r'
96: 'M'
97: 'o'
98: 'r'
99: 'e'
100: 'R'
101: 'e'
102: 'f'
103: 'e'
104: 'r'
105: 'e'
106: 'n'
107: 'c'
108: 'e'
109: 'N'
110: 'o'
111: 't'
112: '['
113: ']'
114: 'b'
115: 'o'
116: 'o'
117: 'l'
118: '['
119: ']'
120: 'i'
121: 'n'
122: 't'
123: '['
124: ']'
125: 'u'
126: 'i'
127: 'n'
128: 't'
129: '['
130: ']'
131: 'd'
132: 'o'
133: 'u'
134: 'b'
135: 'l'
136: 'e'
137: '['
138: ']'
139: 's'
140: 't'
141: 'r'
142: 'i'
143: 'n'
144: 'g'
145: '['
146: ']'
147: '['
148: ']'
149: 'b'
150: 'y'
151: 't'
152: 'e'
153: 't'
154: 'r'
155: 'u'
156: 'e'
157: 'f'
158: 'a'
159: 'l'
160: 's'
161: 'e'
162: '='
163: '('
164: ')'
165: '{'
166: '}'
167: ','
168: '/'
169: '/'
170: '\n'
171: '/'
172: '*'
173: '*'
174: '*'
175: '/'
176: ' '
177: '\t'
178: '\n'
179: '\r'
180: '0'
181: '0'
182: 'x'
183: 'X'
184: '-'
185: 'e'
186: 'E'
187: '+'
188: '-'
189: '.'
190: '.'
191: '.'
192: '_'
193: '_'
194: 'd'
195: 'o'
196: 'u'
197: 'b'
198: 'l'
199: 'e'
200: 'i'
201: 'n'
202: 't'
203: 'u'
204: 'i'
205: 'n'
206: 't'
207: '['
208: ']'
209: 'b'
210: 'y'
211: 't'
212: 'e'
213: 's'
214: 't'
215: 'r'
216: 'i'
217: 'n'
218: 'g'
219: 'b'
220: 'o'
221: 'o'
222: 'l'
223: '.'
224: '\'
225: 'U'
226: '\'
227: 'u'
228: '\'
229: 'x'
230: '\'
231: '`'
232: '`'
233: '\'
234: 'a'
235: 'b'
236: 'f'
237: 'n'
238: 'r'
239: 't'
240: 'v'
241: '\'
242: '''
243: '"'
244: '"'
245: '"'
246: '''
247: '''
248: '0'-'9'
249: '0'-'7'
250: '0'-'9'
251: 'A'-'F'
252: 'a'-'f'
253: '1'-'9'
254: 'A'-'Z'
255: 'a'-'z'
256: .

*/
