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

package readable

import (
	"io"
)

func endOfVarint(buf []byte) (int, error) {
	l := len(buf)
	i := 0
	for {
		if i >= l {
			return 0, io.ErrUnexpectedEOF
		}
		if buf[i] < 0x80 {
			break
		}
		i++
	}
	return i + 1, nil
}

func decodeVarint(buf []byte) (x uint64, n int, err error) {
	i := 0
	l := len(buf)

	for shift := uint(0); ; shift += 7 {
		if i >= l {
			err = io.ErrUnexpectedEOF
			return
		}
		b := buf[i]
		i++
		x |= (uint64(b) & 0x7F) << shift
		if b < 0x80 {
			break
		}
	}
	n = i
	return
}

func next(buf []byte) (int, error) {
	key, n, err := decodeVarint(buf)
	if err != nil {
		return n, err
	}
	//currentField := int32(u >> 3)
	wireType := int(key & 0x7)
	switch wireType {
	case 0: //WireVarint:
		nn, err := endOfVarint(buf[n:])
		return n + nn, err
	case 1: //WireFixed64:
		return n + 8, nil
	case 2: //WireLengthDelimited:
		l, nn, err := decodeVarint(buf[n:])
		n += nn + int(l)
		return n, err
	case 5: //WireFixed32:
		return n + 4, nil
	}
	panic("unreachable")
}
