package lexer

import (

	// "fmt"
	// "github.com/katydid/katydid/expr/util"

	"github.com/katydid/katydid/expr/token"
	"io/ioutil"
	"unicode/utf8"
)

const (
	NoState    = -1
	NumStates  = 323
	NumSymbols = 217
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
37: '6'
38: '4'
39: '['
40: ']'
41: 'i'
42: 'n'
43: 't'
44: '3'
45: '2'
46: '['
47: ']'
48: 'u'
49: 'i'
50: 'n'
51: 't'
52: '6'
53: '4'
54: '['
55: ']'
56: 'u'
57: 'i'
58: 'n'
59: 't'
60: '3'
61: '2'
62: '['
63: ']'
64: 'd'
65: 'o'
66: 'u'
67: 'b'
68: 'l'
69: 'e'
70: '['
71: ']'
72: 'f'
73: 'l'
74: 'o'
75: 'a'
76: 't'
77: '['
78: ']'
79: 's'
80: 't'
81: 'r'
82: 'i'
83: 'n'
84: 'g'
85: '['
86: ']'
87: '['
88: ']'
89: 'b'
90: 'y'
91: 't'
92: 'e'
93: 't'
94: 'r'
95: 'u'
96: 'e'
97: 'f'
98: 'a'
99: 'l'
100: 's'
101: 'e'
102: '='
103: '('
104: ')'
105: '{'
106: '}'
107: ','
108: '/'
109: '/'
110: '\n'
111: '/'
112: '*'
113: '*'
114: '*'
115: '/'
116: ' '
117: '\t'
118: '\n'
119: '\r'
120: '0'
121: '0'
122: 'x'
123: 'X'
124: '-'
125: 'e'
126: 'E'
127: '+'
128: '-'
129: '.'
130: '.'
131: '.'
132: '_'
133: '_'
134: 'd'
135: 'o'
136: 'u'
137: 'b'
138: 'l'
139: 'e'
140: 'f'
141: 'l'
142: 'o'
143: 'a'
144: 't'
145: 'i'
146: 'n'
147: 't'
148: '6'
149: '4'
150: 'u'
151: 'i'
152: 'n'
153: 't'
154: '6'
155: '4'
156: 'i'
157: 'n'
158: 't'
159: '3'
160: '2'
161: 'u'
162: 'i'
163: 'n'
164: 't'
165: '3'
166: '2'
167: '['
168: ']'
169: 'b'
170: 'y'
171: 't'
172: 'e'
173: 's'
174: 't'
175: 'r'
176: 'i'
177: 'n'
178: 'g'
179: 'b'
180: 'o'
181: 'o'
182: 'l'
183: '.'
184: '\'
185: 'U'
186: '\'
187: 'u'
188: '\'
189: 'x'
190: '\'
191: '`'
192: '`'
193: '\'
194: 'a'
195: 'b'
196: 'f'
197: 'n'
198: 'r'
199: 't'
200: 'v'
201: '\'
202: '''
203: '"'
204: '"'
205: '"'
206: '''
207: '''
208: '0'-'9'
209: '0'-'7'
210: '0'-'9'
211: 'A'-'F'
212: 'a'-'f'
213: '1'-'9'
214: 'A'-'Z'
215: 'a'-'z'
216: .

*/
