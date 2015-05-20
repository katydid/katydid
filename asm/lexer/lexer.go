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
	NumStates  = 337
	NumSymbols = 235
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
26: 'r'
27: 'o'
28: 'o'
29: 't'
30: '.'
31: 'i'
32: 'n'
33: 'i'
34: 't'
35: 'f'
36: 'i'
37: 'n'
38: 'a'
39: 'l'
40: 'f'
41: 'u'
42: 'n'
43: 'c'
44: '['
45: ']'
46: 'b'
47: 'o'
48: 'o'
49: 'l'
50: '['
51: ']'
52: 'i'
53: 'n'
54: 't'
55: '6'
56: '4'
57: '['
58: ']'
59: 'i'
60: 'n'
61: 't'
62: '3'
63: '2'
64: '['
65: ']'
66: 'u'
67: 'i'
68: 'n'
69: 't'
70: '6'
71: '4'
72: '['
73: ']'
74: 'u'
75: 'i'
76: 'n'
77: 't'
78: '3'
79: '2'
80: '['
81: ']'
82: 'd'
83: 'o'
84: 'u'
85: 'b'
86: 'l'
87: 'e'
88: '['
89: ']'
90: 'f'
91: 'l'
92: 'o'
93: 'a'
94: 't'
95: '['
96: ']'
97: 's'
98: 't'
99: 'r'
100: 'i'
101: 'n'
102: 'g'
103: '['
104: ']'
105: '['
106: ']'
107: 'b'
108: 'y'
109: 't'
110: 'e'
111: 't'
112: 'r'
113: 'u'
114: 'e'
115: 'f'
116: 'a'
117: 'l'
118: 's'
119: 'e'
120: '='
121: '('
122: ')'
123: '{'
124: '}'
125: ','
126: '/'
127: '/'
128: '\n'
129: '/'
130: '*'
131: '*'
132: '*'
133: '/'
134: ' '
135: '\t'
136: '\n'
137: '\r'
138: '0'
139: '0'
140: 'x'
141: 'X'
142: '-'
143: 'e'
144: 'E'
145: '+'
146: '-'
147: '.'
148: '.'
149: '.'
150: '_'
151: '_'
152: 'd'
153: 'o'
154: 'u'
155: 'b'
156: 'l'
157: 'e'
158: 'f'
159: 'l'
160: 'o'
161: 'a'
162: 't'
163: 'i'
164: 'n'
165: 't'
166: '6'
167: '4'
168: 'u'
169: 'i'
170: 'n'
171: 't'
172: '6'
173: '4'
174: 'i'
175: 'n'
176: 't'
177: '3'
178: '2'
179: 'u'
180: 'i'
181: 'n'
182: 't'
183: '3'
184: '2'
185: '['
186: ']'
187: 'b'
188: 'y'
189: 't'
190: 'e'
191: 's'
192: 't'
193: 'r'
194: 'i'
195: 'n'
196: 'g'
197: 'b'
198: 'o'
199: 'o'
200: 'l'
201: '.'
202: '\'
203: 'U'
204: '\'
205: 'u'
206: '\'
207: 'x'
208: '\'
209: '`'
210: '`'
211: '\'
212: 'a'
213: 'b'
214: 'f'
215: 'n'
216: 'r'
217: 't'
218: 'v'
219: '\'
220: '''
221: '"'
222: '"'
223: '"'
224: '''
225: '''
226: '0'-'9'
227: '0'-'7'
228: '0'-'9'
229: 'A'-'F'
230: 'a'-'f'
231: '1'-'9'
232: 'A'-'Z'
233: 'a'-'z'
234: .

*/
