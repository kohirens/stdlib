package test

import (
	"os"
	"testing"
)

const (
	// FixtureDir Fixture location
	FixtureDir = "testdata"

	// TmpDir Temporarily stores test run output
	TmpDir = "tmp"

	FileMode = 0774
)

func MainSetup(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	if e := os.RemoveAll(TmpDir); e != nil {
		panic(e.Error())
	}
	// Set up a temporary dir for generate files
	if e := os.Mkdir(TmpDir, FileMode); e != nil { // set up a temporary dir for generate files
		panic(e.Error())
	}
	// Run all tests
	exitCode := m.Run()
	// Clean up
	os.Exit(exitCode)
}
