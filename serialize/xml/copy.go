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
	"io"
)

func copyMap(m map[string]string) map[string]string {
	mm := make(map[string]string, len(m))
	for key := range m {
		mm[key] = m[key]
	}
	return mm
}

func copyStack(s *stack) *stack {
	if s == nil {
		return nil
	}
	return &stack{
		next: copyStack(s.next),
		kind: s.kind,
		name: Name{s.name.Space, s.name.Local},
		ok:   s.ok,
	}
}

func copyReader(r io.ByteReader) io.ByteReader {
	b := r.(*Buffer)
	return copyBufferPnt(b)
}

func copyBuffer(b Buffer) Buffer {
	bbuf := make([]byte, len(b.buf))
	copy(bbuf, b.buf)
	return Buffer{
		buf:       bbuf,
		off:       b.off,
		runeBytes: b.runeBytes,
		lastRead:  b.lastRead,
	}
}

func copyBufferPnt(b *Buffer) *Buffer {
	if b == nil {
		return nil
	}
	bb := copyBuffer(*b)
	return &bb
}

func (this *Decoder) Copy() *Decoder {
	return &Decoder{
		Strict:        this.Strict,
		AutoClose:     this.AutoClose,
		Entity:        this.Entity,
		CharsetReader: this.CharsetReader,
		DefaultSpace:  this.DefaultSpace,

		r:         copyReader(this.r),
		buf:       copyBuffer(this.buf),
		saved:     copyBufferPnt(this.saved),
		stk:       copyStack(this.stk),
		free:      copyStack(this.free),
		needClose: this.needClose,
		toClose:   Name{this.toClose.Space, this.toClose.Local},

		nextToken:      CopyToken(this.nextToken),
		nextByte:       this.nextByte,
		ns:             copyMap(this.ns),
		err:            this.err,
		line:           this.line,
		offset:         this.offset,
		unmarshalDepth: this.unmarshalDepth,
	}
}

func copyAttrs(as []Attr) []Attr {
	aas := make([]Attr, len(as))
	for i := range as {
		aas[i] = Attr{as[i].Name, as[i].Value}
	}
	return aas
}
