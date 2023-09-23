package path

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

func TestDirExist(t *testing.T) {
	cases := []struct {
		name, path string
		want       bool
	}{
		{"dirExist", test.FixtureDir + "/dir_that_exist", true},
		{"isFileNotDir", test.FixtureDir + "/dir_that_exist/file_that_exists.md", false},
		{"doesNotExists", test.FixtureDir + "/dir_that_exist/dir-does-not-exist", false},
	}

	for _, sbj := range cases {
		got := DirExist(sbj.path)

		if got != sbj.want {
			t.Errorf("got %v, want %v", got, sbj.want)
		}
	}
}
