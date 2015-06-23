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
	NumStates  = 380
	NumSymbols = 279
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
5: ')'
6: '('
7: ')'
8: '('
9: '-'
10: ')'
11: '('
12: '-'
13: ')'
14: '$'
15: '$'
16: '$'
17: '$'
18: '$'
19: '$'
20: '$'
21: '$'
22: '$'
23: '{'
24: ','
25: '}'
26: 'E'
27: 'm'
28: 'p'
29: 't'
30: 'y'
31: 'E'
32: 'm'
33: 'p'
34: 't'
35: 'y'
36: 'S'
37: 'e'
38: 't'
39: 'T'
40: 'r'
41: 'e'
42: 'e'
43: 'N'
44: 'o'
45: 'd'
46: 'e'
47: 'L'
48: 'e'
49: 'a'
50: 'f'
51: 'N'
52: 'o'
53: 'd'
54: 'e'
55: 'C'
56: 'o'
57: 'n'
58: 'c'
59: 'a'
60: 't'
61: 'O'
62: 'r'
63: 'A'
64: 'n'
65: 'd'
66: 'Z'
67: 'e'
68: 'r'
69: 'o'
70: 'O'
71: 'r'
72: 'M'
73: 'o'
74: 'r'
75: 'e'
76: 'R'
77: 'e'
78: 'f'
79: 'e'
80: 'r'
81: 'e'
82: 'n'
83: 'c'
84: 'e'
85: 'N'
86: 'o'
87: 't'
88: '['
89: ']'
90: 'b'
91: 'o'
92: 'o'
93: 'l'
94: '['
95: ']'
96: 'i'
97: 'n'
98: 't'
99: '6'
100: '4'
101: '['
102: ']'
103: 'i'
104: 'n'
105: 't'
106: '3'
107: '2'
108: '['
109: ']'
110: 'u'
111: 'i'
112: 'n'
113: 't'
114: '6'
115: '4'
116: '['
117: ']'
118: 'u'
119: 'i'
120: 'n'
121: 't'
122: '3'
123: '2'
124: '['
125: ']'
126: 'd'
127: 'o'
128: 'u'
129: 'b'
130: 'l'
131: 'e'
132: '['
133: ']'
134: 'f'
135: 'l'
136: 'o'
137: 'a'
138: 't'
139: '['
140: ']'
141: 's'
142: 't'
143: 'r'
144: 'i'
145: 'n'
146: 'g'
147: '['
148: ']'
149: '['
150: ']'
151: 'b'
152: 'y'
153: 't'
154: 'e'
155: 't'
156: 'r'
157: 'u'
158: 'e'
159: 'f'
160: 'a'
161: 'l'
162: 's'
163: 'e'
164: '='
165: '('
166: ')'
167: '{'
168: '}'
169: ','
170: '/'
171: '/'
172: '\n'
173: '/'
174: '*'
175: '*'
176: '*'
177: '/'
178: ' '
179: '\t'
180: '\n'
181: '\r'
182: '0'
183: '0'
184: 'x'
185: 'X'
186: '-'
187: 'e'
188: 'E'
189: '+'
190: '-'
191: '.'
192: '.'
193: '.'
194: '_'
195: '_'
196: 'd'
197: 'o'
198: 'u'
199: 'b'
200: 'l'
201: 'e'
202: 'f'
203: 'l'
204: 'o'
205: 'a'
206: 't'
207: 'i'
208: 'n'
209: 't'
210: '6'
211: '4'
212: 'u'
213: 'i'
214: 'n'
215: 't'
216: '6'
217: '4'
218: 'i'
219: 'n'
220: 't'
221: '3'
222: '2'
223: 'u'
224: 'i'
225: 'n'
226: 't'
227: '3'
228: '2'
229: '['
230: ']'
231: 'b'
232: 'y'
233: 't'
234: 'e'
235: 's'
236: 't'
237: 'r'
238: 'i'
239: 'n'
240: 'g'
241: 'b'
242: 'o'
243: 'o'
244: 'l'
245: '.'
246: '\'
247: 'U'
248: '\'
249: 'u'
250: '\'
251: 'x'
252: '\'
253: '`'
254: '`'
255: '\'
256: 'a'
257: 'b'
258: 'f'
259: 'n'
260: 'r'
261: 't'
262: 'v'
263: '\'
264: '''
265: '"'
266: '"'
267: '"'
268: '''
269: '''
270: '0'-'9'
271: '0'-'7'
272: '0'-'9'
273: 'A'-'F'
274: 'a'-'f'
275: '1'-'9'
276: 'A'-'Z'
277: 'a'-'z'
278: .

*/
