package test

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const (
	FixturesDir = "testdata"
	SubCmdFlags = "SUB_CMD_FLAGS"
	TmpDir      = "tmp"
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
	vFlag := ""
	if testing.Verbose() {
		vFlag = "-test.v=true"
	}

	// call the generated test binary directly
	// Have it the function runAppMain.
	cmd := exec.Command(os.Args[0], vFlag)

	wd, err1 := os.Getwd()
	if err1 != nil {
		panic("could not get current working directory: " + err1.Error())
	}

	// Run in the context of the source directory.
	cmd.Dir = wd

	// Set an environment variable
	// 1. Only exist for the life of the test that calls this function.
	// 2. Passes arguments/flag to your app
	// 3. Lets RunAppMain, called in your TestMain function, know when to run the main function.
	subCmdArgs := subEnvVarName + "=" + strings.Join(args, " ")
	cmd.Env = append(os.Environ(), subCmdArgs)

	return cmd
}

// ResetDir Reset a directory by emptying it out.
func ResetDir(directory string, mode os.FileMode) {
	// Delete all files in the directory and the directory itself.
	if e := os.RemoveAll(directory); e != nil {
		panic(fmt.Sprintf("could not clean up all files in %v directory", directory))
	}

	// Make the directory
	if e := os.Mkdir(directory, mode); e != nil {
		panic(fmt.Sprintf("could not make %v directory", directory))
	}
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

// VerboseSubCmdOut Serves as a pass-through function display output to stdout
// and stderr respectively, but only if the verbosity flag is turned on.
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

// RunMain Used for running the application's main function from a call to
// a test function, but in a sub process.
//
//	Function "m" must call os.Exit.
func RunMain(subEnvVarName string, m func()) {
	subCmdArgs, ok := os.LookupEnv(subEnvVarName)
	if !ok { // Do nothing.
		return
	}

	// This was adapted from https://golang.org/src/flag/flag_test.go; lines
	// 596-657 at the time. This is called recursively, because we will have
	// this test call itself in a sub-command when an environment variable name
	// by `subEnvVarName` is set. Note that a call to `main()` MUST exit or
	// you'll spin out of control.
	args := strings.Split(subCmdArgs, " ")
	ogArgs := os.Args
	os.Args = append([]string{os.Args[0]}, args...)

	defer func() {
		os.Args = ogArgs
	}()

	js := strings.Join(ogArgs, " ")
	if strings.Contains(js, "-test.v=true") {
		fmt.Printf("\nsub process os.Args = %v\n", os.Args)
	}

	m()
}
