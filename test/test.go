package test

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// AbsPath  Return the absolute path of the directory or panic if error.
func AbsPath(dir string) string {
	tmp, err1 := filepath.Abs(dir)
	if err1 != nil {
		panic(fmt.Sprintf("could not get absolute path for %s: %v", dir, err1.Error()))
	}
	return tmp
}

// GetTestBinCmd return a command to run the test binary in a sub-process, passing it flags as fixtures to produce expected output; `TestMain`, will be run automatically.
func GetTestBinCmd(subEnvVarName string, args []string) *exec.Cmd {
	// call the generated test binary directly
	// Have it the function runAppMain.
	cmd := exec.Command(os.Args[0])
	// Run in the context of the source directory.
	_, filename, _, _ := runtime.Caller(0)
	cmd.Dir = path.Dir(filename)
	// Set an environment variable
	// 1. Only exist for the life of the test that calls this function.
	// 2. Passes arguments/flag to your app
	// 3. Lets TestMain know when to run the main function.
	subCmdFlags := subEnvVarName + "=" + strings.Join(args, " ")
	cmd.Env = append(os.Environ(), subCmdFlags)

	return cmd
}

// Silencer return a function that prevents output during a test run.
func Silencer() func() {
	// Abort in verbose mode.
	if testing.Verbose() {
		return func() {}
	}
	null, _ := os.Open(os.DevNull)
	sOut := os.Stdout
	sErr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer null.Close()
		os.Stdout = sOut
		os.Stderr = sErr
		log.SetOutput(os.Stderr)
	}
}

func VerboseSubCmdOut(stdo []byte, stde error) ([]byte, error) {
	// Debug
	if testing.Verbose() {
		fmt.Print("\nBEGIN sub-command\n")
		fmt.Printf("stdout:\n%s\n", stdo)

		if stde != nil {
			fmt.Printf("stderr:\n%v\n", stde.Error())
		}

		fmt.Print("\nEND sub-command\n\n")
	}

	return stdo, stde
}
