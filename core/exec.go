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

	err = c.Run()

	if b, err2 := ioutil.ReadAll(c.Stderr.(*bytes.Buffer)); err2 != nil {
		err = err2
	} else {
		stderr = string(b)
	}

	if b, err2 := ioutil.ReadAll(c.Stdout.(*bytes.Buffer)); err2 != nil {
		err = err2
	} else {
		stdout = string(b)
	}

	return
}
