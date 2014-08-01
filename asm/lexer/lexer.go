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
	NumStates  = 293
	NumSymbols = 212
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
61: '.'
62: 'i'
63: 'f'
64: '{'
65: '['
66: ']'
67: 'b'
68: 'o'
69: 'o'
70: 'l'
71: '['
72: ']'
73: 'i'
74: 'n'
75: 't'
76: '6'
77: '4'
78: '['
79: ']'
80: 'i'
81: 'n'
82: 't'
83: '3'
84: '2'
85: '['
86: ']'
87: 'u'
88: 'i'
89: 'n'
90: 't'
91: '6'
92: '4'
93: '['
94: ']'
95: 'u'
96: 'i'
97: 'n'
98: 't'
99: '3'
100: '2'
101: '['
102: ']'
103: 'd'
104: 'o'
105: 'u'
106: 'b'
107: 'l'
108: 'e'
109: '['
110: ']'
111: 'f'
112: 'l'
113: 'o'
114: 'a'
115: 't'
116: '['
117: ']'
118: 's'
119: 't'
120: 'r'
121: 'i'
122: 'n'
123: 'g'
124: '['
125: ']'
126: '['
127: ']'
128: 'b'
129: 'y'
130: 't'
131: 'e'
132: 't'
133: 'r'
134: 'u'
135: 'e'
136: 'f'
137: 'a'
138: 'l'
139: 's'
140: 'e'
141: '='
142: 't'
143: 'h'
144: 'e'
145: 'n'
146: 'e'
147: 'l'
148: 's'
149: 'e'
150: '('
151: ')'
152: '}'
153: ','
154: '/'
155: '/'
156: '\n'
157: '/'
158: '*'
159: '*'
160: '*'
161: '/'
162: ' '
163: '\t'
164: '\n'
165: '\r'
166: '0'
167: '0'
168: 'x'
169: 'X'
170: '-'
171: 'e'
172: 'E'
173: '+'
174: '-'
175: '.'
176: '.'
177: '.'
178: '_'
179: '\'
180: 'U'
181: '\'
182: 'u'
183: '\'
184: 'x'
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
196: '\'
197: '`'
198: '`'
199: '"'
200: '"'
201: '''
202: '''
203: '0'-'9'
204: '0'-'7'
205: '0'-'9'
206: 'A'-'F'
207: 'a'-'f'
208: '1'-'9'
209: 'A'-'Z'
210: 'a'-'z'
211: .

*/
