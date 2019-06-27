package core

import (
	"bytes"
	"io/ioutil"
	"os/exec"
)

// vmig_exec 同步指定给定的命令
func vmig_exec(cmd string, args []string) (stdout string, stderr string, err error) {
	c := exec.Command(cmd, args...)
	c.Stdout = new(bytes.Buffer)
	c.Stderr = new(bytes.Buffer)

	var b []byte

	if err = c.Run(); err != nil {
		if b, err = ioutil.ReadAll(c.Stderr.(*bytes.Buffer)); err != nil {
			return
		}
		stderr = string(b)
		return
	}

	if b, err = ioutil.ReadAll(c.Stdout.(*bytes.Buffer)); err != nil {
		return
	}
	stdout = string(b)

	return
}
