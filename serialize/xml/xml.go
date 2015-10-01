//  Copyright 2015 Walter Schulze
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

package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/katydid/katydid/serialize"
	"io"
	"strconv"
	"strings"
)

type parser struct {
	buf       []byte
	dec       *xml.Decoder
	tok       xml.Token
	isLeaf    bool
	attrs     []xml.Attr
	attrIndex int
	attrValue bool
	attrFirst bool
}

func NewXMLParser() *parser {
	return &parser{}
}

func (p *parser) Init(buf []byte) error {
	p.buf = buf
	p.dec = xml.NewDecoder(bytes.NewBuffer(buf))
	p.dec.Strict = false
	return nil
}

func (p *parser) Copy() serialize.Parser {
	offset := p.dec.InputOffset()
	newp := NewXMLParser()
	err := newp.Init(p.buf[offset:])
	if err != nil {
		panic(err)
	}
	newp.tok, err = newp.dec.Token()
	if err != nil {
		panic(err)
	}
	return newp
}

func hasContent(c xml.CharData) bool {
	return len(strings.TrimSpace(string(c))) > 0
}

func (p *parser) Next() (err error) {
	if p.attrValue {
		if p.attrFirst {
			p.attrFirst = false
			return nil
		} else {
			return io.EOF
		}
	}
	if p.tok == nil {
		if p.attrs != nil {
			p.attrIndex++
			if p.attrIndex < len(p.attrs) {
				return nil
			}
		}
	}
	if p.tok != nil {
		for {
			if s, ok := p.tok.(xml.StartElement); ok {
				fmt.Printf("Skipping %s\n", s.Name)
				if err := p.dec.Skip(); err != nil {
					return err
				}
				break
			} else if _, ok := p.tok.(xml.EndElement); ok {
				return io.EOF
			} else if c, ok := p.tok.(xml.CharData); ok {
				if hasContent(c) {
					break
				}
			} else if _, ok := p.tok.(xml.Comment); ok {
				p.tok, err = p.dec.Token()
				fmt.Printf("Comment Next Token %#v\n", p.tok)
				if err != nil {
					return err
				}
			} else {
				panic(fmt.Sprintf("unknown token %T", p.tok))
			}
		}
	}
	p.tok, err = p.dec.Token()
	fmt.Printf("Next Token %#v\n", p.tok)
	for err == nil {
		if _, ok := p.tok.(xml.StartElement); ok {
			break
		} else if c, ok := p.tok.(xml.CharData); ok {
			if hasContent(c) {
				break
			}
		} else if _, ok := p.tok.(xml.EndElement); ok {
			return io.EOF
		}
		p.tok, err = p.dec.Token()
		fmt.Printf("Next Next Token %#v\n", p.tok)
	}
	return err
}

func (p *parser) IsLeaf() bool {
	if p.tok == nil {
		if p.attrValue {
			return true
		}
		return false
	}
	_, ok := p.tok.(xml.CharData)
	return ok
}

func (p *parser) getValue() string {
	if p.tok == nil && p.attrValue {
		return p.attrs[p.attrIndex].Value
	}
	if c, ok := p.tok.(xml.CharData); ok {
		return string(c)
	}
	return ""
}

func (p *parser) Double() (float64, error) {
	return strconv.ParseFloat(p.getValue(), 64)
}

func (p *parser) Int() (int64, error) {
	i, err := strconv.ParseInt(p.getValue(), 10, 64)
	return int64(i), err
}

func (p *parser) Uint() (uint64, error) {
	i, err := strconv.ParseUint(p.getValue(), 10, 64)
	return uint64(i), err
}

func (p *parser) Bool() (bool, error) {
	return strconv.ParseBool(strings.TrimSpace(p.getValue()))
}

func (p *parser) String() (string, error) {
	if p.tok == nil && p.attrIndex < len(p.attrs) {
		if p.attrValue {
			return p.attrs[p.attrIndex].Value, nil
		} else {
			return p.attrs[p.attrIndex].Name.Local, nil
		}
	}
	if s, ok := p.tok.(xml.StartElement); ok {
		return s.Name.Local, nil
	}
	if c, ok := p.tok.(xml.CharData); ok {
		return string(c), nil
	}
	return "", serialize.ErrNotString
}

func (p *parser) Bytes() ([]byte, error) {
	if c, ok := p.tok.(xml.CharData); ok {
		return []byte(c), nil
	}
	return nil, serialize.ErrNotBytes
}

func (p *parser) Up() {
	if p.tok == nil {
		if p.attrValue {
			p.attrValue = false
			p.attrFirst = false
			return
		}
	}
	if _, ok := p.tok.(xml.EndElement); ok {
		p.tok = nil
		p.attrs = nil
		p.attrIndex = 0
		return
	}
	if err := p.dec.Skip(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}
}

func (p *parser) Down() {
	if p.tok == nil {
		if p.attrIndex < len(p.attrs) {
			p.attrValue = true
			p.attrFirst = true
			return
		}
	}
	if s, ok := p.tok.(xml.StartElement); ok {
		p.tok = nil
		p.attrs = s.Attr
		p.attrIndex = -1
		return
	}
	panic(fmt.Sprintf("not a start element %T", p.tok))
}
