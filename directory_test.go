package stdlib

import (
	"github.com/kohirens/stdlib/internal/test"
	"testing"
)

func TestMain(t *testing.M) {
	test.TestMainSetup(t)
}

func TestCopyToDir(t *testing.T) {
	testCases := []struct {
		name, source, dest string
		want               int64
	}{
		{"canCopyFile", test.FixtureDir + "/copy-file-01.txt", test.TestTmp, 26},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CopyToDir(tc.source, tc.dest, "/")

			if err != nil {
				t.Errorf("got an unexpected error copying file %q to %q", tc.source, tc.dest)
			}

			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestNormalizePath(t *testing.T) {
	testCases := []struct {
		name string
		path string
		want string
	}{
		{"windowsOnUnix", "\\windows\\path\\on\\uUnix", PS + "windows" + PS + "path" + PS + "on" + PS + "uUnix"},
		{"unixOnWindows", "/unix/on/windows", PS + "unix" + PS + "on" + PS + "windows"},
		{"mixedBag", "/i\\dont\\care/really", PS + "i" + PS + "dont" + PS + "care" + PS + "really"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := NormalizePath(tc.path)

			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
