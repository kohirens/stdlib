package test

import (
	"io"
	"os"
	"strings"
	"testing"
)

const fixtureDir = "testdata"

func TestGetHttpResponseFromFile(tRunner *testing.T) {
	gotResponse, err1 := GetHttpResponseFromFile(fixtureDir + string(os.PathSeparator) + "response-body-01.txt")
	if err1 != nil {
		tRunner.Errorf("got an unexpected error: %v", err1.Error())
	}

	gotBytes, err2 := io.ReadAll(gotResponse.Body)
	if err2 != nil {
		tRunner.Errorf("got an unexpected error: %v", err2.Error())
	}

	want := "Salam!"
	if string(gotBytes) != want {
		tRunner.Errorf("got %s, but want %v", string(gotBytes), want)
	}
}

func TestGetHttpResponseFromFileErr(tRunner *testing.T) {
	gotResponse, gotErr := GetHttpResponseFromFile(fixtureDir + string(os.PathSeparator) + "does-not-exist.txt")
	if gotResponse != nil {
		tRunner.Error("a response was returned; expected nil")
	}

	want := "The system cannot find the file specified"
	want2 := "no such file or directory"
	if !strings.Contains(gotErr.Error(), want) && !strings.Contains(gotErr.Error(), want2) {
		tRunner.Errorf("%q does not contain %q", gotErr.Error(), want)
	}
}
