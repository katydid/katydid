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
	NumStates  = 300
	NumSymbols = 223
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
29: 'T'
30: 'r'
31: 'e'
32: 'e'
33: 'N'
34: 'o'
35: 'd'
36: 'e'
37: 'L'
38: 'e'
39: 'a'
40: 'f'
41: 'N'
42: 'o'
43: 'd'
44: 'e'
45: 'C'
46: 'o'
47: 'n'
48: 'c'
49: 'a'
50: 't'
51: 'O'
52: 'r'
53: 'A'
54: 'n'
55: 'd'
56: 'Z'
57: 'e'
58: 'r'
59: 'o'
60: 'O'
61: 'r'
62: 'M'
63: 'o'
64: 'r'
65: 'e'
66: 'R'
67: 'e'
68: 'f'
69: 'e'
70: 'r'
71: 'e'
72: 'n'
73: 'c'
74: 'e'
75: 'N'
76: 'o'
77: 't'
78: '['
79: ']'
80: 'b'
81: 'o'
82: 'o'
83: 'l'
84: '['
85: ']'
86: 'i'
87: 'n'
88: 't'
89: '['
90: ']'
91: 'u'
92: 'i'
93: 'n'
94: 't'
95: '['
96: ']'
97: 'd'
98: 'o'
99: 'u'
100: 'b'
101: 'l'
102: 'e'
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
129: '('
130: ')'
131: '{'
132: '}'
133: ','
134: '/'
135: '/'
136: '\n'
137: '/'
138: '*'
139: '*'
140: '*'
141: '/'
142: ' '
143: '\t'
144: '\n'
145: '\r'
146: '0'
147: '0'
148: 'x'
149: 'X'
150: '-'
151: 'e'
152: 'E'
153: '+'
154: '-'
155: '.'
156: '.'
157: '.'
158: '_'
159: '_'
160: 'd'
161: 'o'
162: 'u'
163: 'b'
164: 'l'
165: 'e'
166: 'i'
167: 'n'
168: 't'
169: 'u'
170: 'i'
171: 'n'
172: 't'
173: '['
174: ']'
175: 'b'
176: 'y'
177: 't'
178: 'e'
179: 's'
180: 't'
181: 'r'
182: 'i'
183: 'n'
184: 'g'
185: 'b'
186: 'o'
187: 'o'
188: 'l'
189: '.'
190: '\'
191: 'U'
192: '\'
193: 'u'
194: '\'
195: 'x'
196: '\'
197: '`'
198: '`'
199: '\'
200: 'a'
201: 'b'
202: 'f'
203: 'n'
204: 'r'
205: 't'
206: 'v'
207: '\'
208: '''
209: '"'
210: '"'
211: '"'
212: '''
213: '''
214: '0'-'9'
215: '0'-'7'
216: '0'-'9'
217: 'A'-'F'
218: 'a'-'f'
219: '1'-'9'
220: 'A'-'Z'
221: 'a'-'z'
222: .

*/
