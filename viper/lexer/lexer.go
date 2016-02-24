package lexer

import (

	// "fmt"
	// "github.com/katydid/katydid/viper/util"

	"github.com/katydid/katydid/viper/token"
	"io/ioutil"
	"unicode/utf8"
)

const (
	NoState    = -1
	NumStates  = 312
	NumSymbols = 228
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
19: 's'
20: 't'
21: 'a'
22: 'r'
23: 't'
24: 'f'
25: 'i'
26: 'n'
27: 'a'
28: 'l'
29: 'i'
30: 'n'
31: 't'
32: 'e'
33: 'r'
34: 'n'
35: 'a'
36: 'l'
37: 'c'
38: 'a'
39: 'l'
40: 'l'
41: 'r'
42: 'e'
43: 't'
44: 'u'
45: 'r'
46: 'n'
47: '['
48: ']'
49: 'b'
50: 'o'
51: 'o'
52: 'l'
53: '['
54: ']'
55: 'i'
56: 'n'
57: 't'
58: '['
59: ']'
60: 'u'
61: 'i'
62: 'n'
63: 't'
64: '['
65: ']'
66: 'd'
67: 'o'
68: 'u'
69: 'b'
70: 'l'
71: 'e'
72: '['
73: ']'
74: 's'
75: 't'
76: 'r'
77: 'i'
78: 'n'
79: 'g'
80: '['
81: ']'
82: '['
83: ']'
84: 'b'
85: 'y'
86: 't'
87: 'e'
88: 't'
89: 'r'
90: 'u'
91: 'e'
92: 'f'
93: 'a'
94: 'l'
95: 's'
96: 'e'
97: '='
98: '('
99: ')'
100: '{'
101: '}'
102: ','
103: ';'
104: '#'
105: '&'
106: '|'
107: '['
108: ']'
109: ':'
110: '!'
111: '*'
112: '_'
113: '~'
114: '.'
115: '@'
116: '-'
117: '>'
118: '='
119: '='
120: '!'
121: '='
122: '<'
123: '>'
124: '<'
125: '='
126: '>'
127: '='
128: '~'
129: '='
130: '*'
131: '='
132: '^'
133: '='
134: '$'
135: '='
136: ':'
137: ':'
138: '?'
139: '/'
140: '/'
141: '\n'
142: '/'
143: '*'
144: '*'
145: '*'
146: '/'
147: ' '
148: '\t'
149: '\n'
150: '\r'
151: '0'
152: '0'
153: 'x'
154: 'X'
155: '-'
156: 'e'
157: 'E'
158: '+'
159: '-'
160: '.'
161: '.'
162: '.'
163: '_'
164: '_'
165: 'd'
166: 'o'
167: 'u'
168: 'b'
169: 'l'
170: 'e'
171: 'i'
172: 'n'
173: 't'
174: 'u'
175: 'i'
176: 'n'
177: 't'
178: '['
179: ']'
180: 'b'
181: 'y'
182: 't'
183: 'e'
184: 's'
185: 't'
186: 'r'
187: 'i'
188: 'n'
189: 'g'
190: 'b'
191: 'o'
192: 'o'
193: 'l'
194: '.'
195: '\'
196: 'U'
197: '\'
198: 'u'
199: '\'
200: 'x'
201: '\'
202: '`'
203: '`'
204: '\'
205: 'a'
206: 'b'
207: 'f'
208: 'n'
209: 'r'
210: 't'
211: 'v'
212: '\'
213: '''
214: '"'
215: '"'
216: '"'
217: '''
218: '''
219: '0'-'9'
220: '0'-'7'
221: '0'-'9'
222: 'A'-'F'
223: 'a'-'f'
224: '1'-'9'
225: 'A'-'Z'
226: 'a'-'z'
227: .

*/
