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
	_ = os.RemoveAll(TestTmp)
	// Set up a temporary dir for generate files
	_ = os.Mkdir(TestTmp, DirMode) // set up a temporary dir for generate files
	// Run all tests
	exitCode := m.Run()
	// Clean up
	os.Exit(exitCode)
}
