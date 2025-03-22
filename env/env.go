package env

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/kohirens/stdlib/cli"
	"io"
	"os"
	"regexp"
	"strings"
)

// Env Get output of `emv` as map of strings
func Env() (map[string]string, error) {
	stdout, stderr, exitCode, _ := cli.RunCommand(".", "env", []string{})

	if exitCode != 0 || stderr != nil {
		return nil, fmt.Errorf("exit code %v: %v", exitCode, stderr.Error())
	}

	envMap, e1 := parseEnv(bytes.NewReader(stdout))
	if e1 != nil {
		return nil, fmt.Errorf("could not parsing env output: %s", e1.Error())
	}

	return envMap, nil
}

// Get An environment variable falling back to a default value when not set.
func Get(key, def string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return v
}

// PrintEnv we'll need to use the environment shell to get the environment variable since Terraform forbids getting
// variables other than TF_VAR_* with os.LookupEnv.
func PrintEnv(name string) (string, bool) {
	ok := true

	value, err, exitCode, _ := cli.RunCommand(".", "printenv", []string{name})

	if exitCode != 0 || err != nil {
		ok = false
	}

	return string(value), ok
}

// parseEnv Parse `env` command output into a map.
func parseEnv(r io.Reader) (map[string]string, error) {
	envMap := make(map[string]string)
	scanner := bufio.NewScanner(r)
	var previousEnvName string

	for scanner.Scan() {
		lineB := scanner.Bytes()
		line := strings.TrimSpace(string(lineB))

		// Let's use RegExp to extract a variables name, vs split by equal sign.
		// Details for how this regular expression was crafted can be read at:
		// https://pubs.opengroup.org/onlinepubs/9699919799.2018edition/
		reForEnvNames := regexp.MustCompile("^([a-zA-Z_][a-z-A-Z0-9_]+)=(.*)$")

		if reForEnvNames.Match(lineB) { // begin variable
			res := reForEnvNames.FindStringSubmatch(line)
			envMap[res[1]] = res[2]
			// Find the end of a multi-line variable value by looking for
			// the beginning of the next variable or EOF
			// There ar 2 possibilities:
			// 1. End on the current line, we are done.
			// 2. Is multi-line and ends somewhere later.
			previousEnvName = res[1]
		} else { // append value to a previous variable
			// assume multiline because of the nature of /etc/environment files,
			// not to be confused with .env files. Should only contain key=value
			// pairs on every line until EOF; and no comments or empty lines.
			envMap[previousEnvName] += "\n" + line
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return envMap, nil
}
