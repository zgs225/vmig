package core

import (
	"testing"
)

func TestVmigExec(t *testing.T) {
	stdout, stderr, err := vmig_exec("ls", []string{"."})

	t.Log("Stdout: ", stdout)
	t.Log("Stderr: ", stderr)

	if err != nil {
		t.Error(err)
	}

	if len(stdout) == 0 {
		t.Error("ls command stdout error. Got ", stdout)
	}
}
