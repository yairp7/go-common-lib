package os

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func ExecCmd(cmdStr string) (out []byte, exitCode int, err error) {
	parts := strings.Split(cmdStr, " ")
	cmd := exec.Command(parts[0], parts[1:]...)
	out, err = cmd.Output()
	exitCode = cmd.ProcessState.ExitCode()
	return
}

func IsFileExists(f *os.File) bool {
	_, err := f.Stat()
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
