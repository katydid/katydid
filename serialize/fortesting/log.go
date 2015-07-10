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

package fortesting

import (
	"github.com/katydid/katydid/serialize"
	"log"
	"os"
	"time"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

func NewStdLogger() Logger {
	return log.New(os.Stderr, "", log.Llongfile)
}

func NewDelayLogger(delay time.Duration) Logger {
	return &d{
		delay: delay,
		log:   NewStdLogger(),
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
	s serialize.Scanner
	l Logger
}

func NewLogger(s serialize.Scanner, logger Logger) serialize.Scanner {
	return &l{s, logger}
}

func (l *l) Double() (float64, error) {
	v, err := l.s.Double()
	l.l.Printf("scanner.Double() (%v, %v)", v, err)
	return v, err
}

func (l *l) Int() (int64, error) {
	v, err := l.s.Int()
	l.l.Printf("scanner.Int() (%v, %v)", v, err)
	return v, err
}

func (l *l) Uint() (uint64, error) {
	v, err := l.s.Uint()
	l.l.Printf("scanner.Uint() (%v, %v)", v, err)
	return v, err
}

func (l *l) Bool() (bool, error) {
	v, err := l.s.Bool()
	l.l.Printf("scanner.Bool() (%v, %v)", v, err)
	return v, err
}

func (l *l) String() (string, error) {
	v, err := l.s.String()
	l.l.Printf("scanner.String() (%v, %v)", v, err)
	return v, err
}

func (l *l) Bytes() ([]byte, error) {
	v, err := l.s.Bytes()
	l.l.Printf("scanner.Bytes() (%v, %v)", v, err)
	return v, err
}

func (l *l) Copy() serialize.Scanner {
	s := l.s.Copy()
	l.l.Printf("scanner.Copy()")
	return s
}

func (l *l) Next() error {
	err := l.s.Next()
	l.l.Printf("scanner.Next() (%v)", err)
	return err
}

func (l *l) IsLeaf() bool {
	v := l.s.IsLeaf()
	l.l.Printf("scanner.IsLeaf() (%v)", v)
	return v
}

func (l *l) Name() string {
	v := l.s.Name()
	l.l.Printf("scanner.Name() (%v)", v)
	return v
}

func (l *l) Up() {
	l.s.Up()
	l.l.Printf("scanner.Up()")
	return
}

func (l *l) Down() {
	l.s.Down()
	l.l.Printf("scanner.Down()")
	return
}
