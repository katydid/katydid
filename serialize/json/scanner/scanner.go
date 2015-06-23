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

package scanner

import (
	"bytes"
	"fmt"
	"github.com/katydid/katydid/serialize"
	"io"
	"strconv"
)

type JsonTokens interface {
	GetTokenId(tokenString string) (int, error)
	Lookup(parentToken int, name []byte) int
}

func errInString(buf []byte) error {
	return fmt.Errorf("katydid/json/scanner error in json string: %s", string(buf))
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

func (s *jsonScanner) expected(expected string) error {
	panic(fmt.Errorf("katydid/json/scanner: expected '%s' at offset %d, but got '%c'", expected, s.offset, s.buf[s.offset]))
}

func (s *jsonScanner) scanOpenObject() error {
	if s.buf[s.offset] == '{' {
		s.offset++
	} else {
		return s.expected("{")
	}
	return s.skipSpace()
}

func (s *jsonScanner) scanOpenArray() error {
	if s.buf[s.offset] == '[' {
		s.offset++
	} else {
		return s.expected("[")
	}
	return s.skipSpace()
}

func (s *jsonScanner) scanString() error {
	s.startValueOffset = s.offset
	n, err := scanString(s.buf[s.offset:])
	if err != nil {
		return err
	}
	s.endValueOffset = s.offset + n
	s.offset += n
	return s.skipSpace()
}

func (s *jsonScanner) scanName() error {
	startOffset := s.offset
	n, err := scanString(s.buf[s.offset:])
	if err != nil {
		return err
	}
	s.offset += n
	s.name = s.buf[startOffset+1 : s.offset-1]
	return s.skipSpace()
}

func (s *jsonScanner) skipSpace() error {
	if s.offset >= len(s.buf) {
		return io.ErrShortBuffer
	}
	s.offset += skipSpace(s.buf[s.offset:])
	if s.offset >= len(s.buf) {
		return io.ErrShortBuffer
	}
	return nil
}

func (s *jsonScanner) scanTrue() error {
	if s.offset+4 >= len(s.buf) {
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

func (s *jsonScanner) scanFalse() error {
	if s.offset+5 >= len(s.buf) {
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

func (s *jsonScanner) scanNull() error {
	if s.offset+4 >= len(s.buf) {
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

func (s *jsonScanner) scanObject() error {
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
		s.expected("}")
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

func (s *jsonScanner) scanNumber() error {
	s.startValueOffset = s.offset
	if s.buf[s.offset] == '-' {
		s.offset++
		if s.offset >= len(s.buf) {
			return io.ErrShortBuffer
		}
	}
	if s.buf[s.offset] == '0' {
		s.offset++
		if s.offset >= len(s.buf) {
			return io.ErrShortBuffer
		}
	} else if isDigit19(s.buf[s.offset]) {
		s.offset++
		if s.offset >= len(s.buf) {
			return io.ErrShortBuffer
		}
		for isDigit(s.buf[s.offset]) {
			s.offset++
			if s.offset >= len(s.buf) {
				return io.ErrShortBuffer
			}
		}
	}
	if s.buf[s.offset] == '.' {
		s.offset++
		if s.offset >= len(s.buf) {
			return io.ErrShortBuffer
		}
		for isDigit(s.buf[s.offset]) {
			s.offset++
			if s.offset >= len(s.buf) {
				return io.ErrShortBuffer
			}
		}
	}
	if s.buf[s.offset] == 'e' || s.buf[s.offset] == 'E' {
		s.offset++
		if s.offset >= len(s.buf) {
			return io.ErrShortBuffer
		}
		if s.buf[s.offset] == '+' || s.buf[s.offset] == '-' {
			s.offset++
			if s.offset >= len(s.buf) {
				return io.ErrShortBuffer
			}
		}
		for isDigit(s.buf[s.offset]) {
			s.offset++
			if s.offset >= len(s.buf) {
				return io.ErrShortBuffer
			}
		}
	}
	s.endValueOffset = s.offset
	return nil
}

func (s *jsonScanner) scanValue() error {
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
		s.firstArrayValue = true
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

func (s *jsonScanner) scanColon() error {
	if s.buf[s.offset] != ':' {
		return s.expected(":")
	}
	s.offset++
	return s.skipSpace()
}

func (s *jsonScanner) scanCloseObject() error {
	if s.buf[s.offset] == '}' {
		return io.EOF
	}
	return s.expected("}")
}

func (s *jsonScanner) scanCloseArray() error {
	if s.buf[s.offset] == ']' {
		return io.EOF
	}
	return s.expected("]")
}

func (s *jsonScanner) scanComma() error {
	if s.buf[s.offset] != ',' {
		return s.expected(",")
	}
	s.offset++
	return s.skipSpace()
}

func (s *jsonScanner) scanArray() error {
	s.inArray = true
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
			s.inArray = false
			return s.scanCloseArray()
		}
	}
	if s.buf[s.offset] == ']' {
		s.inArray = false
		return s.scanCloseArray()
	}
	return s.scanValue()
}

func (s *jsonScanner) nextValueInObject() error {
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

func (s *jsonScanner) Next() error {
	s.isValueObject = false
	if err := s.skipSpace(); err != nil {
		return err
	}
	if !s.inArray {
		return s.nextValueInObject()
	}
	err := s.scanArray()
	if err != nil {
		if err == io.EOF {
			s.offset++
			return s.nextValueInObject()
		}
		return err
	}
	return nil
}

func (s *jsonScanner) Id() int {
	return s.tokens.Lookup(s.parentToken, s.name)
}

func (s *jsonScanner) IsLeaf() bool {
	return !s.isValueObject
}

func (s *jsonScanner) Value() []byte {
	return s.buf[s.startValueOffset:s.endValueOffset]
}

func (s *jsonScanner) Float64() (float64, error) {
	v := string(s.Value())
	i, err := strconv.ParseFloat(v, 64)
	return i, err
}

func (s *jsonScanner) Float32() (float32, error) {
	v := string(s.Value())
	i, err := strconv.ParseFloat(v, 32)
	return float32(i), err
}

func (s *jsonScanner) Int64() (int64, error) {
	v := string(s.Value())
	i, err := strconv.ParseInt(v, 10, 64)
	return int64(i), err
}

func (s *jsonScanner) Uint64() (uint64, error) {
	v := string(s.Value())
	i, err := strconv.ParseUint(v, 10, 64)
	return uint64(i), err
}

func (s *jsonScanner) Int32() (int32, error) {
	v := string(s.Value())
	i, err := strconv.ParseInt(v, 10, 32)
	return int32(i), err
}

func (s *jsonScanner) Bool() (bool, error) {
	v := string(s.Value())
	if v == "true" {
		return true, nil
	}
	return false, nil
}

func (s *jsonScanner) String() (string, error) {
	v := s.Value()
	if v[0] != '"' {
		return "", serialize.ErrNotString
	}
	return string(v[1 : len(v)-1]), nil
}

func (s *jsonScanner) Bytes() ([]byte, error) {
	return nil, serialize.ErrNotBytes
}

func (s *jsonScanner) Uint32() (uint32, error) {
	v := string(s.Value())
	i, err := strconv.ParseUint(v, 10, 32)
	return uint32(i), err
}

func NewScanner(tokens JsonTokens, rootToken int) *jsonScanner {
	return &jsonScanner{
		tokens: tokens,
		state: state{
			parentToken:      rootToken,
			firstObjectValue: true,
		},
		stack: make([]state, 0, 10),
	}
}

func (s *jsonScanner) init(buf []byte) {
	s.state = state{
		parentToken:      s.parentToken,
		firstObjectValue: true,
		buf:              buf,
	}
	s.stack = s.stack[:0]
}

func (s *jsonScanner) Init(buf []byte) error {
	s.init(buf)
	return nil
}

type jsonScanner struct {
	tokens JsonTokens
	state
	stack []state
}

type state struct {
	parentToken      int
	buf              []byte
	offset           int
	name             []byte
	startValueOffset int
	endValueOffset   int
	inArray          bool
	firstObjectValue bool
	firstArrayValue  bool
	isValueObject    bool
}

func (s *jsonScanner) Name() string {
	return string(s.name)
}

func (s *jsonScanner) Up() {
	top := len(s.stack) - 1
	s.state = s.stack[top]
	s.stack = s.stack[:top]
}

func (s *jsonScanner) Down() {
	s.stack = append(s.stack, s.state)
	s.state = state{
		parentToken:      s.tokens.Lookup(s.parentToken, s.name),
		buf:              s.buf[s.startValueOffset:s.endValueOffset],
		firstObjectValue: true,
	}
}
