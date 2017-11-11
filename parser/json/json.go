//  Copyright 2013 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

//Package json contains the implementation of a JSON parser.
package json

import (
	"bytes"
	"fmt"
	"io"
	"strconv"

	"github.com/katydid/katydid/parser"
)

//ErrUnquote returns an error that resulted from trying to unquote a string.
var ErrUnquote = fmt.Errorf("json: error unquoting string")

func errInString(buf []byte) error {
	return fmt.Errorf("katydid/json error in json string: %s", string(buf))
}

func isString(buf []byte) bool {
	return buf[0] == '"'
}

func scanString(buf []byte) (int, error) {
	escaped := false
	udigits := -1
	if buf[0] != '"' {
		return 0, errInString(buf)
	}
	for i, c := range buf[1:] {
		if escaped {
			switch c {
			case 'b', 'f', 'n', 'r', 't', '\\', '/', '"':
				escaped = false
				continue
			case 'u':
				udigits = 0
				continue
			}
			return 0, errInString(buf)
		}
		if udigits >= 0 {
			if '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F' {
				udigits++
			} else {
				return 0, errInString(buf)
			}
			if udigits == 4 {
				udigits = -1
			}
			continue
		}
		if c == '"' {
			return i + 2, nil
		}
		if c == '\\' {
			escaped = true
			continue
		}
		if c < 0x20 {
			return 0, errInString(buf)
		}
	}
	panic("unreachable")
}

func isNumber(c byte) bool {
	return (c == '-') || ((c >= '0') && (c <= '9'))
}

func isSpace(c byte) bool {
	return (c == ' ') || (c == '\n') || (c == '\r') || (c == '\t')
}

func skipSpace(buf []byte) int {
	for i, c := range buf {
		if !isSpace(c) {
			return i
		}
	}
	return len(buf)
}

func (s *jsonParser) expected(expected string) error {
	return fmt.Errorf("katydid/json/parser: expected '%s' at offset %d, but got '%c'", expected, s.offset, s.buf[s.offset])
}

func (s *jsonParser) scanOpenObject() error {
	if s.buf[s.offset] == '{' {
		s.offset++
	} else {
		return s.expected("{")
	}
	return s.skipSpace()
}

func (s *jsonParser) scanOpenArray() error {
	if s.buf[s.offset] == '[' {
		s.offset++
	} else {
		return s.expected("[")
	}
	return s.skipSpace()
}

func (s *jsonParser) scanString() error {
	s.startValueOffset = s.offset
	n, err := scanString(s.buf[s.offset:])
	if err != nil {
		return err
	}
	s.offset += n
	if s.offset >= len(s.buf) {
		s.offset = len(s.buf)
	}
	s.endValueOffset = s.offset
	return s.skipSpace()
}

func (s *jsonParser) scanName() error {
	startOffset := s.offset
	n, err := scanString(s.buf[s.offset:])
	if err != nil {
		return err
	}
	s.offset += n
	var ok bool
	s.name, ok = unquote(s.buf[startOffset:s.offset])
	if !ok {
		return ErrUnquote
	}
	return s.skipSpace()
}

func (s *jsonParser) skipSpace() error {
	if s.offset > len(s.buf) {
		return io.ErrShortBuffer
	}
	s.offset += skipSpace(s.buf[s.offset:])
	if s.offset > len(s.buf) {
		return io.ErrShortBuffer
	}
	return nil
}

func (s *jsonParser) scanTrue() error {
	if s.offset+4 > len(s.buf) {
		return io.ErrShortBuffer
	}
	if !bytes.Equal(s.buf[s.offset:s.offset+4], []byte("true")) {
		return s.expected("true")
	}
	s.startValueOffset = s.offset
	s.endValueOffset = s.offset + 4
	s.offset += 4
	return s.skipSpace()
}

func (s *jsonParser) scanFalse() error {
	if s.offset+5 > len(s.buf) {
		return io.ErrShortBuffer
	}
	if !bytes.Equal(s.buf[s.offset:s.offset+5], []byte("false")) {
		return s.expected("false")
	}
	s.startValueOffset = s.offset
	s.endValueOffset = s.offset + 5
	s.offset += 5
	return s.skipSpace()
}

func (s *jsonParser) scanNull() error {
	if s.offset+4 > len(s.buf) {
		return io.ErrShortBuffer
	}
	if !bytes.Equal(s.buf[s.offset:s.offset+4], []byte("null")) {
		return s.expected("null")
	}
	s.startValueOffset = s.offset
	s.endValueOffset = s.offset + 4
	s.offset += 4
	return s.skipSpace()
}

func (s *jsonParser) scanArray() error {
	count := 0
	index := 0
	for i, c := range s.buf[s.offset:] {
		if c == '[' {
			count++
		}
		if c == ']' {
			count--
		}
		if count == 0 {
			index = i
			break
		}
	}
	if count != 0 {
		return s.expected("]")
	}
	s.startValueOffset = s.offset
	s.endValueOffset = s.offset + index + 1
	s.offset += index + 1
	s.isValueArray = true
	return s.skipSpace()
}

func (s *jsonParser) scanObject() error {
	count := 0
	index := 0
	for i, c := range s.buf[s.offset:] {
		if c == '{' {
			count++
		}
		if c == '}' {
			count--
		}
		if count == 0 {
			index = i
			break
		}
	}
	if count != 0 {
		return s.expected("}")
	}
	s.startValueOffset = s.offset
	s.endValueOffset = s.offset + index + 1
	s.offset += index + 1
	s.isValueObject = true
	return s.skipSpace()
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isDigit19(c byte) bool {
	return c >= '1' && c <= '9'
}

func (s *jsonParser) scanNumber() error {
	s.startValueOffset = s.offset
	if s.buf[s.offset] == '-' {
		s.offset++
		if s.offset >= len(s.buf) {
			return io.ErrShortBuffer
		}
	}
	if s.buf[s.offset] == '0' {
		s.offset++
		if s.offset > len(s.buf) {
			return io.ErrShortBuffer
		}
	} else if isDigit19(s.buf[s.offset]) {
		s.offset++
		if s.offset > len(s.buf) {
			return io.ErrShortBuffer
		}
		for s.offset < len(s.buf) && isDigit(s.buf[s.offset]) {
			s.offset++
			if s.offset > len(s.buf) {
				return io.ErrShortBuffer
			}
		}
	}
	if s.offset < len(s.buf) && s.buf[s.offset] == '.' {
		s.offset++
		if s.offset > len(s.buf) {
			return io.ErrShortBuffer
		}
		for s.offset < len(s.buf) && isDigit(s.buf[s.offset]) {
			s.offset++
			if s.offset > len(s.buf) {
				return io.ErrShortBuffer
			}
		}
	}
	if s.offset < len(s.buf) &&
		(s.buf[s.offset] == 'e' || s.buf[s.offset] == 'E') {
		s.offset++
		if s.offset > len(s.buf) {
			return io.ErrShortBuffer
		}
		if s.offset < len(s.buf) {
			if s.buf[s.offset] == '+' || s.buf[s.offset] == '-' {
				s.offset++
				if s.offset > len(s.buf) {
					return io.ErrShortBuffer
				}
			}
		}
		for s.offset < len(s.buf) && isDigit(s.buf[s.offset]) {
			s.offset++
			if s.offset > len(s.buf) {
				return io.ErrShortBuffer
			}
		}
	}
	s.endValueOffset = s.offset
	return nil
}

func (s *jsonParser) scanValue() error {
	c := s.buf[s.offset]
	if isNumber(c) {
		return s.scanNumber()
	}
	switch c {
	case '"':
		return s.scanString()
	case '{':
		return s.scanObject()
	case '[':
		return s.scanArray()
	case 't':
		return s.scanTrue()
	case 'f':
		return s.scanFalse()
	case 'n':
		return s.scanNull()
	}
	return s.expected("value")
}

func (s *jsonParser) scanColon() error {
	if s.buf[s.offset] != ':' {
		return s.expected(":")
	}
	s.offset++
	return s.skipSpace()
}

func (s *jsonParser) scanCloseObject() error {
	if s.buf[s.offset] == '}' {
		return io.EOF
	}
	return s.expected("}")
}

func (s *jsonParser) scanCloseArray() error {
	if s.buf[s.offset] == ']' {
		return io.EOF
	}
	return s.expected("]")
}

func (s *jsonParser) scanComma() error {
	if s.buf[s.offset] != ',' {
		return s.expected(",")
	}
	s.offset++
	return s.skipSpace()
}

func (s *jsonParser) nextValueInArray() error {
	if s.firstArrayValue {
		if err := s.scanOpenArray(); err != nil {
			return err
		}
		s.firstArrayValue = false
	} else {
		if s.buf[s.offset] == ',' {
			if err := s.scanComma(); err != nil {
				return err
			}
		} else {
			return s.scanCloseArray()
		}
	}
	if s.buf[s.offset] == ']' {
		return s.scanCloseArray()
	}
	return s.scanValue()
}

func (s *jsonParser) nextValueInObject() error {
	if s.firstObjectValue {
		if err := s.scanOpenObject(); err != nil {
			return err
		}
		s.firstObjectValue = false
	} else {
		if s.buf[s.offset] == ',' {
			if err := s.scanComma(); err != nil {
				return err
			}
		} else {
			if err := s.scanCloseObject(); err != nil {
				return err
			}
		}
	}
	if isString(s.buf[s.offset:]) {
		if err := s.scanName(); err != nil {
			return err
		}
		if err := s.scanColon(); err != nil {
			return err
		}
		if err := s.scanValue(); err != nil {
			return err
		}
		return nil
	}
	return s.scanCloseObject()
}

func (s *jsonParser) Next() error {
	if s.isLeaf {
		if s.firstObjectValue {
			s.firstObjectValue = false
			return nil
		}
		return io.EOF
	}
	s.isValueObject = false
	s.isValueArray = false
	if err := s.skipSpace(); err != nil {
		return err
	}
	if s.inArray {
		if !s.firstArrayValue {
			s.arrayIndex++
		}
		return s.nextValueInArray()
	}
	return s.nextValueInObject()
}

func (s *jsonParser) IsLeaf() bool {
	return s.isLeaf
}

func (s *jsonParser) Value() []byte {
	return s.buf[s.startValueOffset:s.endValueOffset]
}

func (s *jsonParser) Double() (float64, error) {
	if s.isLeaf {
		v := string(s.Value())
		i, err := strconv.ParseFloat(v, 64)
		return i, err
	}
	return 0, parser.ErrNotDouble
}

func (s *jsonParser) Int() (int64, error) {
	if s.isLeaf {
		v := string(s.Value())
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			f, ferr := strconv.ParseFloat(v, 64)
			if ferr != nil {
				return i, err
			}
			if float64(int64(f)) == f {
				return int64(f), nil
			}
		}
		return i, err
	}
	if s.inArray {
		return int64(s.arrayIndex), nil
	}
	return 0, parser.ErrNotInt
}

func (s *jsonParser) Uint() (uint64, error) {
	if s.isLeaf {
		v := string(s.Value())
		i, err := strconv.ParseUint(v, 10, 64)
		return uint64(i), err
	}
	return 0, parser.ErrNotUint
}

func (s *jsonParser) Bool() (bool, error) {
	if s.isLeaf {
		v := string(s.Value())
		if v == "true" {
			return true, nil
		}
		if v == "false" {
			return false, nil
		}
	}
	return false, parser.ErrNotBool
}

func (s *jsonParser) String() (string, error) {
	if s.isLeaf {
		v := s.Value()
		if v[0] != '"' {
			return "", parser.ErrNotString
		}
		res, ok := unquote(v)
		if !ok {
			return "", ErrUnquote
		}
		return res, nil
	}
	if s.inArray {
		return "", parser.ErrNotString
	}
	return s.name, nil
}

func (s *jsonParser) Bytes() ([]byte, error) {
	return nil, parser.ErrNotBytes
}

//JsonParser is a parser for JSON
type JsonParser interface {
	parser.Interface
	//Init initialises the parser with a byte buffer containing JSON.
	Init(buf []byte) error
	Reset() error
}

//NewJsonParser returns a new JSON parser.
func NewJsonParser() JsonParser {
	return &jsonParser{
		state: state{
			firstObjectValue: true,
		},
		stack: make([]state, 0, 10),
	}
}

func (s *jsonParser) Init(buf []byte) error {
	s.state = state{
		firstObjectValue: true,
		buf:              buf,
	}
	s.stack = s.stack[:0]
	if err := s.skipSpace(); err != nil {
		return err
	}
	if s.buf[s.offset] == '{' {
		//do nothing
	} else if s.buf[s.offset] == '[' {
		if err := s.scanValue(); err != nil {
			return err
		}
		s.inArray = true
		s.firstArrayValue = true
		s.buf = s.buf[s.startValueOffset:s.endValueOffset]
		s.offset = 0
	} else {
		if err := s.scanValue(); err != nil {
			return err
		}
		s.state.isLeaf = true
		s.state.firstObjectValue = true
	}
	return nil
}

func (s *jsonParser) Reset() error {
	if len(s.stack) > 0 {
		return s.Init(s.stack[0].buf)
	}
	return s.Init(s.buf)
}

type jsonParser struct {
	state
	stack []state
}

type state struct {
	buf              []byte
	offset           int
	name             string
	startValueOffset int
	endValueOffset   int
	inArray          bool
	firstObjectValue bool
	firstArrayValue  bool
	isValueObject    bool
	isValueArray     bool
	isLeaf           bool
	arrayIndex       int
}

func (s *jsonParser) Up() {
	top := len(s.stack) - 1
	s.state = s.stack[top]
	s.stack = s.stack[:top]
}

func (s *jsonParser) Down() {
	if s.isValueObject {
		s.stack = append(s.stack, s.state)
		s.state = state{
			buf:              s.buf[s.startValueOffset:s.endValueOffset],
			firstObjectValue: true,
		}
	} else if s.isValueArray {
		s.stack = append(s.stack, s.state)
		s.state = state{
			buf:             s.buf[s.startValueOffset:s.endValueOffset],
			firstArrayValue: true,
			inArray:         true,
		}
	} else {
		s.stack = append(s.stack, s.state)
		s.state.isLeaf = true
		s.state.firstObjectValue = true
	}
}
