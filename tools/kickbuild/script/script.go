package script

import (
	"os/exec"
)

// RunDump runs a command capturing it's output.
func RunDump(env []string, dir, command string, args ...string) (dump []byte, err error) {
	cmd := exec.Command(command, args...)
	if len(env) > 0 {
		cmd.Env = env
	}
	cmd.Dir = dir
	dump, err = cmd.CombinedOutput()
	return
}

// Run runs a command.
func Run(command string, args ...string) (err error) {
	cmd := exec.Command(command, args...)
	err = cmd.Start()
	return
}
