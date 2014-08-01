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

package exec

import (
	"fmt"
)

type errVarint struct {
	buf []byte
}

func newErrVarint(buf []byte) error {
	buffer := make([]byte, 10)
	nn := copy(buffer, buf)
	return &errVarint{buffer[:nn]}
}

func (this *errVarint) Error() string {
	return fmt.Sprintf("error decoding varint from %#v", this.buf)
}

var errBufferOverlow = fmt.Errorf("buffer overflow")

var errUnknownWireType = fmt.Errorf("unknown wire type")

type ProtoMap interface {
	Trans(src int, key uint64) (int, bool)
	IsLeave(int) bool
}

type Table interface {
	Trans(src int, input int) (int, error)
	NoEscapeFrom(src int) bool
	StateToName(state int) string
}

type Link interface {
	ProtoToStart(protoState int) (startState int, exists bool)
	IfEval(protoState int, buf []byte) (state int, exists bool)
}

type Catcher interface {
	Catch(error)
	GetError() error
	Clear()
}

func NewExec(protomap ProtoMap, table Table, link Link, catcher Catcher, rootState int, startState int, acceptState int) *Exec {
	return &Exec{
		protomap:    protomap,
		table:       table,
		Link:        link,
		catcher:     catcher,
		rootState:   rootState,
		startState:  startState,
		acceptState: acceptState,
	}
}

type Exec struct {
	protomap    ProtoMap
	table       Table
	Link        Link
	catcher     Catcher
	rootState   int
	startState  int
	acceptState int
}

func (this *Exec) Eval(buf []byte) (bool, error) {
	this.catcher.Clear()
	i, err := this.eval(this.rootState, this.startState, buf)
	if err != nil {
		return false, err
	}
	return i == this.acceptState, this.catcher.GetError()
}

func uvarint(buf []byte) (uint64, int, error) {
	var uv uint64
	var shift uint
	for i, b := range buf {
		if b < 0x80 {
			if i > 9 || i == 9 && b > 1 {
				return 0, 0, newErrVarint(buf)
			}
			return uv | uint64(b)<<shift, i + 1, nil
		}
		uv |= (uint64(b) & 0x7F) << shift
		shift += 7
	}
	return 0, 0, newErrVarint(buf)
}

func length(wireType int, buf []byte) (prefix int, l int, err error) {
	switch wireType {
	case 0:
		_, n2, err := uvarint(buf)
		return 0, n2, err
	case 1:
		return 0, 8, nil
	case 2:
		l, n2, err := uvarint(buf)
		return n2, int(l), err
	case 5:
		return 0, 4, nil
	}
	return 0, 0, errUnknownWireType
}

func (this *Exec) eval(mapState int, automataState int, buf []byte) (int, error) {
	if this.table.NoEscapeFrom(automataState) {
		return automataState, nil
	}
	i := 0
	for i < len(buf) {
		v, n, err := uvarint(buf[i:])
		if err != nil {
			return 0, err
		}
		i += n
		wireType := int(v & 0x7)
		p, n, err := length(wireType, buf[i:])
		if err != nil {
			return 0, err
		}
		i += p
		newMapState, ok := this.protomap.Trans(mapState, v)
		if !ok {
			i += n
			continue
		}
		var input int
		if this.protomap.IsLeave(newMapState) {
			var ok bool
			input, ok = this.Link.IfEval(newMapState, buf[i:i+n])
			i += n
			if !ok {
				continue
			}
		} else {
			newStartState, ok := this.Link.ProtoToStart(newMapState)
			if !ok {
				i += n
				continue
			} else {
				input, err = this.eval(newMapState, newStartState, buf[i:i+n])
				if err != nil {
					return 0, err
				}
				i += n
			}
		}
		//debugging old := automataState
		automataState, err = this.table.Trans(automataState, input)
		//debugging fmt.Printf("transition: <%s> <%s> = <%s>\n", this.table.StateToName(old), this.table.StateToName(input), this.table.StateToName(automataState))
		if err != nil {
			return 0, err
		}
		if this.table.NoEscapeFrom(automataState) {
			return automataState, nil
		}
	}
	if i > len(buf) {
		return 0, errBufferOverlow
	}
	return automataState, nil
}
