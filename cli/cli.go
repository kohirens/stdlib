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

// RunCommandWithInput run an external program.
func RunCommandWithInput(workDir string, input []byte, program string, args ...string) (cmdOut []byte, cmdErr error, exitCode int) {
	cmd := exec.Command(program, args...)
	cmd.Env = os.Environ()
	cmd.Dir = workDir

	if input != nil {
		cmdIn, err1 := cmd.StdinPipe()
		if err1 != nil {
			return []byte{}, err1, 1
		}

		defer func() {
			_ = cmdIn.Close()
		}()

		// write the input
		_, err2 := cmdIn.Write(input)
		if err2 != nil {
			return []byte{}, err1, 1
		}
	}

	cmdOut, cmdErr = cmd.CombinedOutput()
	exitCode = cmd.ProcessState.ExitCode()

	log.Infof("sub command run: %q", cmd.String())

	return
}
