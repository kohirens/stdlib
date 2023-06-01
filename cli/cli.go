package cli

import (
	"github.com/kohirens/stdlib/log"
	"os"
	"os/exec"
)

type CommandRunner func(program, workDir string, args ...string) (cmdOut []byte, cmdErr error, exitCode int)

// RunCommand run an external program.
func RunCommand(program, workDir string, args ...string) (cmdOut []byte, cmdErr error, exitCode int) {
	cmd := exec.Command(program, args...)
	cmd.Env = os.Environ()
	cmd.Dir = workDir
	cmdOut, cmdErr = cmd.CombinedOutput()
	exitCode = cmd.ProcessState.ExitCode()

	log.Infof("sub command run: %q", cmd.String())
	return
}
