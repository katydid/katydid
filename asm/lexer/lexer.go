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
	NumStates  = 489
	NumSymbols = 250
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
15: '('
16: ')'
17: '$'
18: '('
19: ')'
20: '$'
21: '('
22: ')'
23: '$'
24: '('
25: ')'
26: '$'
27: '('
28: ')'
29: '$'
30: '('
31: ')'
32: '$'
33: '('
34: ')'
35: '$'
36: '('
37: ')'
38: '$'
39: '('
40: ')'
41: '{'
42: ','
43: '}'
44: 'r'
45: 'o'
46: 'o'
47: 't'
48: '.'
49: 'i'
50: 'f'
51: '{'
52: '['
53: ']'
54: 'b'
55: 'o'
56: 'o'
57: 'l'
58: '['
59: ']'
60: 'i'
61: 'n'
62: 't'
63: '6'
64: '4'
65: '['
66: ']'
67: 'i'
68: 'n'
69: 't'
70: '3'
71: '2'
72: '['
73: ']'
74: 'u'
75: 'i'
76: 'n'
77: 't'
78: '6'
79: '4'
80: '['
81: ']'
82: 'u'
83: 'i'
84: 'n'
85: 't'
86: '3'
87: '2'
88: '['
89: ']'
90: 'd'
91: 'o'
92: 'u'
93: 'b'
94: 'l'
95: 'e'
96: '['
97: ']'
98: 'f'
99: 'l'
100: 'o'
101: 'a'
102: 't'
103: '['
104: ']'
105: 's'
106: 't'
107: 'r'
108: 'i'
109: 'n'
110: 'g'
111: '['
112: ']'
113: '['
114: ']'
115: 'b'
116: 'y'
117: 't'
118: 'e'
119: 't'
120: 'r'
121: 'u'
122: 'e'
123: 'f'
124: 'a'
125: 'l'
126: 's'
127: 'e'
128: '='
129: 't'
130: 'h'
131: 'e'
132: 'n'
133: 'e'
134: 'l'
135: 's'
136: 'e'
137: '('
138: ')'
139: '}'
140: ','
141: '/'
142: '/'
143: '\n'
144: '/'
145: '*'
146: '*'
147: '*'
148: '/'
149: ' '
150: '\t'
151: '\n'
152: '\r'
153: '0'
154: '0'
155: 'x'
156: 'X'
157: '-'
158: 'e'
159: 'E'
160: '+'
161: '-'
162: '.'
163: '.'
164: '.'
165: '_'
166: '_'
167: 'd'
168: 'o'
169: 'u'
170: 'b'
171: 'l'
172: 'e'
173: 'f'
174: 'l'
175: 'o'
176: 'a'
177: 't'
178: 'i'
179: 'n'
180: 't'
181: '6'
182: '4'
183: 'u'
184: 'i'
185: 'n'
186: 't'
187: '6'
188: '4'
189: 'i'
190: 'n'
191: 't'
192: '3'
193: '2'
194: 'u'
195: 'i'
196: 'n'
197: 't'
198: '3'
199: '2'
200: '['
201: ']'
202: 'b'
203: 'y'
204: 't'
205: 'e'
206: 's'
207: 't'
208: 'r'
209: 'i'
210: 'n'
211: 'g'
212: 'b'
213: 'o'
214: 'o'
215: 'l'
216: '.'
217: '\'
218: 'U'
219: '\'
220: 'u'
221: '\'
222: 'x'
223: '\'
224: '`'
225: '`'
226: '\'
227: 'a'
228: 'b'
229: 'f'
230: 'n'
231: 'r'
232: 't'
233: 'v'
234: '\'
235: '''
236: '"'
237: '"'
238: '"'
239: '''
240: '''
241: '0'-'9'
242: '0'-'7'
243: '0'-'9'
244: 'A'-'F'
245: 'a'-'f'
246: '1'-'9'
247: 'A'-'Z'
248: 'a'-'z'
249: .

*/
