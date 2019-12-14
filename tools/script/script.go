package script

import (
	"os/exec"
	"runtime"
)

// RunDump runs a command capturing it's output.
func RunDump(env []string, dir, command string, args ...string) (dump []byte, err error) {
	cmd := exec.Command(command, args...)
	if len(env) > 0 {
		cmd.Env = env
	}
	if len(dir) > 0 {
		cmd.Dir = dir
	}
	dump, err = cmd.CombinedOutput()
	return
}

// Run runs a command.
func Run(command string, args ...string) (err error) {
	cmd := exec.Command(command, args...)
	err = cmd.Start()
	return
}

// Edit starts the default text editor.
func Edit(path string) bool {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	args2 := append(args[1:], path)
	cmd := exec.Command(args[0], args2...)
	return cmd.Start() == nil
}
