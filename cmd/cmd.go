package cmd

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
)

// 运行cmd
func Cmd(command string, param ...string) error {
	cmd := exec.Command(command, param...)
	cmd.Stdout = os.Stdout
	err := &bytes.Buffer{}
	cmd.Stderr = err
	_ = cmd.Run()
	if err.Len() > 0 {
		return errors.New(err.String())
	} else {
		return nil
	}
}
