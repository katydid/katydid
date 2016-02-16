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

package debug

import (
	"github.com/katydid/katydid/serialize"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

func NewLineLogger() Logger {
	return &line{log.New(os.Stderr, "", 0)}
}

type line struct {
	l Logger
}

func (l *line) Printf(format string, v ...interface{}) {
	_, thisfile, _, ok := runtime.Caller(0)
	if !ok {
		l.l.Printf("<weirdlyunknown>:0: "+format, v...)
		return
	}
	i := 0
	for {
		i++
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			l.l.Printf("<unknown>:"+strconv.Itoa(i)+": "+format, v...)
			return
		}
		if file == thisfile {
			continue
		}
		_, name := filepath.Split(file)
		l.l.Printf(name+":"+strconv.Itoa(line)+": "+format, v...)
		return
	}
	panic("unreachable")
}

func NewDelayLogger(delay time.Duration) Logger {
	return &d{
		delay: delay,
		log:   NewLineLogger(),
	}
}

type d struct {
	log   Logger
	delay time.Duration
}

func (d *d) Printf(format string, v ...interface{}) {
	d.log.Printf(format, v...)
	time.Sleep(d.delay)
}

type l struct {
	name   string
	s      serialize.Parser
	l      Logger
	copies int
}

func NewLogger(s serialize.Parser, logger Logger) serialize.Parser {
	return &l{"parser", s, logger, 0}
}

func (l *l) Double() (float64, error) {
	v, err := l.s.Double()
	l.l.Printf(l.name+".Double() (%v, %v)", v, err)
	return v, err
}

func (l *l) Int() (int64, error) {
	v, err := l.s.Int()
	l.l.Printf(l.name+".Int() (%v, %v)", v, err)
	return v, err
}

func (l *l) Uint() (uint64, error) {
	v, err := l.s.Uint()
	l.l.Printf(l.name+".Uint() (%v, %v)", v, err)
	return v, err
}

func (l *l) Bool() (bool, error) {
	v, err := l.s.Bool()
	l.l.Printf(l.name+".Bool() (%v, %v)", v, err)
	return v, err
}

func (l *l) String() (string, error) {
	v, err := l.s.String()
	l.l.Printf(l.name+".String() (%v, %v)", v, err)
	return v, err
}

func (l *l) Bytes() ([]byte, error) {
	v, err := l.s.Bytes()
	l.l.Printf(l.name+".Bytes() (%v, %v)", v, err)
	return v, err
}

func (l *l) Next() error {
	err := l.s.Next()
	l.l.Printf(l.name+".Next() (%v)", err)
	return err
}

func (l *l) IsLeaf() bool {
	v := l.s.IsLeaf()
	l.l.Printf(l.name+".IsLeaf() (%v)", v)
	return v
}

func (l *l) Up() {
	l.s.Up()
	l.l.Printf(l.name + ".Up()")
	return
}

func (l *l) Down() {
	l.s.Down()
	l.l.Printf(l.name + ".Down()")
	return
}
