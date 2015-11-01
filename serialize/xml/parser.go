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
	"fmt"
	"github.com/katydid/katydid/serialize"
	"io"
	"strconv"
	"strings"
)

type parser struct {
	buf       []byte
	dec       *Decoder
	tok       Token
	isLeaf    bool
	attrs     []Attr
	attrIndex int
	attrValue bool
	attrFirst bool
}

func NewXMLParser() *parser {
	return &parser{}
}

func (p *parser) Init(buf []byte) error {
	p.buf = buf
	p.dec = NewDecoder(NewBuffer(buf))
	p.dec.Strict = false
	return nil
}

func (p *parser) Copy() serialize.Parser {
	return &parser{
		buf:       p.buf,
		dec:       p.dec.Copy(),
		tok:       CopyToken(p.tok),
		isLeaf:    p.isLeaf,
		attrs:     copyAttrs(p.attrs),
		attrIndex: p.attrIndex,
		attrValue: p.attrValue,
		attrFirst: p.attrFirst,
	}
}

func hasContent(c CharData) bool {
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
			if s, ok := p.tok.(StartElement); ok {
				fmt.Printf("Skipping %s\n", s.Name)
				if err := p.dec.Skip(); err != nil {
					fmt.Printf("Skip err = %v\n", err)
					return err
				}
				break
			} else if _, ok := p.tok.(EndElement); ok {
				return io.EOF
			} else if c, ok := p.tok.(CharData); ok {
				if hasContent(c) {
					break
				}
			} else if _, ok := p.tok.(Comment); ok {
				p.tok, err = p.dec.Token()
				fmt.Printf("Comment Next Token %#v, err = %v\n", p.tok, err)
				if err != nil {
					return err
				}
			} else {
				panic(fmt.Sprintf("unknown token %T", p.tok))
			}
		}
	}
	p.tok, err = p.dec.Token()
	fmt.Printf("Next Token %#v, err %v\n", p.tok, err)
	for err == nil {
		if _, ok := p.tok.(StartElement); ok {
			break
		} else if c, ok := p.tok.(CharData); ok {
			if hasContent(c) {
				break
			}
		} else if _, ok := p.tok.(EndElement); ok {
			return io.EOF
		}
		p.tok, err = p.dec.Token()
		fmt.Printf("Next Next Token %#v, err = %v\n", p.tok, err)
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
	_, ok := p.tok.(CharData)
	return ok
}

func (p *parser) getValue() string {
	if p.tok == nil && p.attrValue {
		return p.attrs[p.attrIndex].Value
	}
	if c, ok := p.tok.(CharData); ok {
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
	if s, ok := p.tok.(StartElement); ok {
		return s.Name.Local, nil
	}
	if c, ok := p.tok.(CharData); ok {
		return string(c), nil
	}
	return "", serialize.ErrNotString
}

func (p *parser) Bytes() ([]byte, error) {
	if c, ok := p.tok.(CharData); ok {
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
	if _, ok := p.tok.(EndElement); ok {
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
	if s, ok := p.tok.(StartElement); ok {
		p.tok = nil
		p.attrs = s.Attr
		p.attrIndex = -1
		return
	}
	panic(fmt.Sprintf("not a start element %T", p.tok))
}
