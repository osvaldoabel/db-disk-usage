package du

import (
	"os"
	"os/exec"
)

func Run(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).Output()
}

func RunWithStdin(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	return cmd.Output()
}

func RunWithStdout(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	return cmd.Output()
}

func RunWithStderr(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stderr
	return cmd.Output()
}

func RunWithStdinStdout(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	return cmd.Output()
}

// LookPath searches for a specific named file in the directories named by the PATH environment variable.
func LookPath(command string) (string, error) {
	return exec.LookPath(command)
}
