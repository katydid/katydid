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
		panic(err)
	}
	cmdErr := &bytes.Buffer{}
	cmd.Stderr = cmdErr
	_, err = stdin.Write([]byte(dot))
	if err != nil {
		panic(err)
	}
	if err := stdin.Close(); err != nil {
		panic(err)
	}
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	err = nil
	errString := strings.TrimSpace(string(cmdErr.Bytes()))
	if len(errString) > 0 {
		err = errors.New(errString)
	}
	return string(out), err
}
