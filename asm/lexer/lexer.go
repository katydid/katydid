package lexer

import (

	// "fmt"
	// "github.com/awalterschulze/katydid/asm/util"

	"github.com/awalterschulze/katydid/asm/token"
	"io/ioutil"
	"unicode/utf8"
)

const (
	NoState    = -1
	NumStates  = 307
	NumSymbols = 233
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
0: 'i'
1: 'n'
2: 't'
3: '6'
4: '4'
5: '('
6: ')'
7: 'i'
8: 'n'
9: 't'
10: '3'
11: '2'
12: '('
13: ')'
14: 'u'
15: 'i'
16: 'n'
17: 't'
18: '6'
19: '4'
20: '('
21: ')'
22: 'u'
23: 'i'
24: 'n'
25: 't'
26: '3'
27: '2'
28: '('
29: ')'
30: 'd'
31: 'o'
32: 'u'
33: 'b'
34: 'l'
35: 'e'
36: '('
37: '-'
38: ')'
39: 'f'
40: 'l'
41: 'o'
42: 'a'
43: 't'
44: '('
45: '-'
46: ')'
47: '_'
48: '['
49: ']'
50: 'b'
51: 'y'
52: 't'
53: 'e'
54: '{'
55: ','
56: '}'
57: 'r'
58: 'o'
59: 'o'
60: 't'
61: '='
62: '.'
63: 'i'
64: 'f'
65: 't'
66: 'h'
67: 'e'
68: 'n'
69: 'e'
70: 'l'
71: 's'
72: 'e'
73: '{'
74: '}'
75: '('
76: ')'
77: '='
78: '='
79: '<'
80: '<'
81: '='
82: '>'
83: '>'
84: '='
85: '&'
86: '&'
87: '|'
88: '|'
89: 'o'
90: 'r'
91: 'a'
92: 'n'
93: 'd'
94: ','
95: '['
96: ']'
97: 'b'
98: 'o'
99: 'o'
100: 'l'
101: '['
102: ']'
103: 'i'
104: 'n'
105: 't'
106: '6'
107: '4'
108: '['
109: ']'
110: 'i'
111: 'n'
112: 't'
113: '3'
114: '2'
115: '['
116: ']'
117: 'u'
118: 'i'
119: 'n'
120: 't'
121: '6'
122: '4'
123: '['
124: ']'
125: 'u'
126: 'i'
127: 'n'
128: 't'
129: '3'
130: '2'
131: '['
132: ']'
133: 'd'
134: 'o'
135: 'u'
136: 'b'
137: 'l'
138: 'e'
139: '['
140: ']'
141: 'f'
142: 'l'
143: 'o'
144: 'a'
145: 't'
146: '['
147: ']'
148: 's'
149: 't'
150: 'r'
151: 'i'
152: 'n'
153: 'g'
154: '['
155: ']'
156: '['
157: ']'
158: 'b'
159: 'y'
160: 't'
161: 'e'
162: 't'
163: 'r'
164: 'u'
165: 'e'
166: 'f'
167: 'a'
168: 'l'
169: 's'
170: 'e'
171: '/'
172: '/'
173: '\n'
174: '/'
175: '*'
176: '*'
177: '*'
178: '/'
179: ' '
180: '\t'
181: '\n'
182: '\r'
183: '0'
184: '0'
185: 'x'
186: 'X'
187: '-'
188: 'e'
189: 'E'
190: '+'
191: '-'
192: '.'
193: '.'
194: '.'
195: '_'
196: '\'
197: 'U'
198: '\'
199: 'u'
200: '\'
201: 'x'
202: '\'
203: 'a'
204: 'b'
205: 'f'
206: 'n'
207: 'r'
208: 't'
209: 'v'
210: '\'
211: '''
212: '"'
213: '\'
214: '`'
215: '`'
216: '"'
217: '"'
218: '''
219: '''
220: ' '
221: '\t'
222: '\n'
223: '\r'
224: '0'-'9'
225: '0'-'7'
226: '0'-'9'
227: 'A'-'F'
228: 'a'-'f'
229: '1'-'9'
230: 'A'-'Z'
231: 'a'-'z'
232: .

*/
