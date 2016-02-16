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
	NumStates  = 313
	NumSymbols = 227
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
9: '$'
10: '$'
11: '$'
12: '$'
13: '$'
14: '$'
15: '{'
16: ','
17: '}'
18: 's'
19: 't'
20: 'a'
21: 'r'
22: 't'
23: 'f'
24: 'i'
25: 'n'
26: 'a'
27: 'l'
28: 'i'
29: 'n'
30: 't'
31: 'e'
32: 'r'
33: 'n'
34: 'a'
35: 'l'
36: 'c'
37: 'a'
38: 'l'
39: 'l'
40: 'r'
41: 'e'
42: 't'
43: 'u'
44: 'r'
45: 'n'
46: '['
47: ']'
48: 'b'
49: 'o'
50: 'o'
51: 'l'
52: '['
53: ']'
54: 'i'
55: 'n'
56: 't'
57: '['
58: ']'
59: 'u'
60: 'i'
61: 'n'
62: 't'
63: '['
64: ']'
65: 'd'
66: 'o'
67: 'u'
68: 'b'
69: 'l'
70: 'e'
71: '['
72: ']'
73: 's'
74: 't'
75: 'r'
76: 'i'
77: 'n'
78: 'g'
79: '['
80: ']'
81: '['
82: ']'
83: 'b'
84: 'y'
85: 't'
86: 'e'
87: 't'
88: 'r'
89: 'u'
90: 'e'
91: 'f'
92: 'a'
93: 'l'
94: 's'
95: 'e'
96: '='
97: '('
98: ')'
99: '{'
100: '}'
101: ','
102: ';'
103: '#'
104: '&'
105: '|'
106: '['
107: ']'
108: ':'
109: '!'
110: '*'
111: '_'
112: '~'
113: '.'
114: '@'
115: '-'
116: '>'
117: '='
118: '='
119: '!'
120: '='
121: '<'
122: '>'
123: '<'
124: '='
125: '>'
126: '='
127: '~'
128: '='
129: '*'
130: '='
131: '^'
132: '='
133: '$'
134: '='
135: ':'
136: ':'
137: '?'
138: '/'
139: '/'
140: '\n'
141: '/'
142: '*'
143: '*'
144: '*'
145: '/'
146: ' '
147: '\t'
148: '\n'
149: '\r'
150: '0'
151: '0'
152: 'x'
153: 'X'
154: '-'
155: 'e'
156: 'E'
157: '+'
158: '-'
159: '.'
160: '.'
161: '.'
162: '_'
163: '_'
164: 'd'
165: 'o'
166: 'u'
167: 'b'
168: 'l'
169: 'e'
170: 'i'
171: 'n'
172: 't'
173: 'u'
174: 'i'
175: 'n'
176: 't'
177: '['
178: ']'
179: 'b'
180: 'y'
181: 't'
182: 'e'
183: 's'
184: 't'
185: 'r'
186: 'i'
187: 'n'
188: 'g'
189: 'b'
190: 'o'
191: 'o'
192: 'l'
193: '.'
194: '\'
195: 'U'
196: '\'
197: 'u'
198: '\'
199: 'x'
200: '\'
201: '`'
202: '`'
203: '\'
204: 'a'
205: 'b'
206: 'f'
207: 'n'
208: 'r'
209: 't'
210: 'v'
211: '\'
212: '''
213: '"'
214: '"'
215: '"'
216: '''
217: '''
218: '0'-'9'
219: '0'-'7'
220: '0'-'9'
221: 'A'-'F'
222: 'a'-'f'
223: '1'-'9'
224: 'A'-'Z'
225: 'a'-'z'
226: .

*/
