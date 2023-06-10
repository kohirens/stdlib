package cli

import (
	"github.com/kohirens/stdlib/log"
	"os"
	"os/exec"
	"strings"
)

type CommandRunner func(program, workDir string, args ...string) (cmdOut []byte, cmdErr error, exitCode int)

// RunCommand run an external program in a sub process.
func RunCommand(
	workDir, program string,
	args []string,
	env map[string]string,
	input []byte,
) (cmdOut []byte, cmdErr error, exitCode int) {
	cmd := exec.Command(program, args...)
	cmd.Dir = workDir
	ce := os.Environ()

	if env != nil {
		for i, kv := range ce {
			p := strings.Split(kv, "=")
			v, ok := env[p[0]]
			if ok {
				ce[i] = p[0] + "=" + v
			}
		}
	}

	cmd.Env = ce

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
