package test

import (
	"os"
	"testing"
)

// FixtureDir Fixture location
const FixtureDir = "testdata"

// TestTmp Temporarily stores test run output,
const TestTmp = "testtmp"

const DirMode = 0774

func TestMainSetup(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	if e := os.RemoveAll(TestTmp); e != nil {
		panic(e.Error())
	}
	// Set up a temporary dir for generate files
	if e := os.Mkdir(TestTmp, DirMode); e != nil { // set up a temporary dir for generate files
		panic(e.Error())
	}
	// Run all tests
	exitCode := m.Run()
	// Clean up
	os.Exit(exitCode)
}
