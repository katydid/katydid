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

package graphviz

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func Installed() bool {
	_, err := exec.LookPath("dot")
	return err == nil
}

func DrawSVG(dot string) (string, error) {
	cmd := exec.Command("dot", "-T", "svg")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	cmdErr := &bytes.Buffer{}
	cmd.Stderr = cmdErr
	_, err = stdin.Write([]byte(dot))
	if err != nil {
		return "", err
	}
	if err := stdin.Close(); err != nil {
		return "", err
	}
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	err = nil
	errString := strings.TrimSpace(string(cmdErr.Bytes()))
	if len(errString) > 0 {
		err = errors.New(errString)
	}
	return string(out), err
}
