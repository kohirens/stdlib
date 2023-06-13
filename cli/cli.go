package cli

import (
	"os"
	"os/exec"
	"strings"
)

type CommandRunner func(wd, program string, args ...string) (cmdOut []byte, cmdErr error, exitCode int, cmdStr string)
type CommandRunnerWithInput func(wd, program string, args ...string) (cmdOut []byte, cmdErr error, exitCode int, cmdStr string)
type CommandRunnerWithInputAndEnv func(wd, program string, args ...string) (cmdOut []byte, cmdErr error, exitCode int, cmdStr string)

// RunCommand run an external program in a sub process.
func RunCommand(wd, program string, args []string) (cmdOut []byte, cmdErr error, exitCode int, cmdStr string) {
	return RunCommandWithInputAndEnv(wd, program, args, nil, nil)
}

// RunCommandWithInput run an external program in a sub process, allowing with
// input.
func RunCommandWithInput(wd string, program string, args []string, input []byte) (cmdOut []byte, cmdErr error, exitCode int, cmdStr string) {
	return RunCommandWithInputAndEnv(wd, program, args, input, nil)
}

// RunCommandWithInputAndEnv run an external program in a sub process, with
// input and setting environment variables in the sub process. It
// will pass in the os.Environ(), overwriting key=value pairs from env map,
// comparison for the key (variable name) is case-sensitive.
func RunCommandWithInputAndEnv(
	wd,
	program string,
	args []string,
	input []byte,
	env map[string]string,
) (cmdOut []byte, cmdErr error, exitCode int, cmdStr string) {
	cmd := exec.Command(program, args...)
	cmd.Dir = wd
	ce := os.Environ()

	// overwrite or set environment variables
	if env != nil {
		ce = AmendStringAry(ce, env)
	}

	cmd.Env = ce

	if input != nil {
		cmdIn, err1 := cmd.StdinPipe()
		if err1 != nil {
			return []byte{}, err1, 1, ""
		}

		defer func() {
			_ = cmdIn.Close()
		}()

		// write the input
		_, err2 := cmdIn.Write(input)
		if err2 != nil {
			return []byte{}, err1, 1, ""
		}
	}

	cmdOut, cmdErr = cmd.CombinedOutput()
	exitCode = cmd.ProcessState.ExitCode()
	cmdStr = cmd.String()

	return
}

// AmendStringAry where []string (like os.Environ()) is string of key=value
// pairs. Pass in a map who's that will overwrite an existing key pair, or append it to the string array.
func AmendStringAry(ce []string, env map[string]string) []string {
	for key, val := range env {
		found := false
		idx := -1
		for i, pair := range ce {
			p := strings.Split(pair, "=")
			k := p[0]
			if key == k {
				found = true
				idx = i
				break
			}
		}
		newPair := key + "=" + val
		if found { //overwrite
			ce[idx] = newPair
		} else { //append
			ce = append(ce, newPair)
		}
	}
	return ce
}
