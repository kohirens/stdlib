package fsio

import (
	"fmt"
	"github.com/kohirens/stdlib/git"
	"github.com/kohirens/stdlib/internal/test"
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.M) {
	test.MainSetup(t)
}

func TestCopyDirToDir(runner *testing.T) {
	td := test.TmpDir + PS + runner.Name()
	_ = os.MkdirAll(td, 0744)

	tests := []struct {
		name    string
		src     string
		dst     string
		ps      string
		wantErr bool
		want    []string
	}{
		{
			"copy all files 01",
			test.FixtureDir + PS + "dir-to-dir-01",
			td,
			PS,
			false,
			[]string{td + PS + "README.md"},
		},
		{
			"copy all files 02",
			test.FixtureDir + PS + "dir-to-dir-02",
			td,
			PS,
			false,
			[]string{
				td + PS + "README.md",
				td + PS + "sub-01" + PS + "README.md",
				td + PS + "sub-02" + PS + "file-01.txt",
				td + PS + "sub-02" + PS + "file-02.txt",
				td + PS + "sub-03" + PS + "README.md",
				td + PS + "sub-03" + PS + "sub-sub-04" + PS + "README-04.md",
			},
		},
	}

	for _, tt := range tests {
		runner.Run(tt.name, func(t *testing.T) {
			if err := CopyDirToDir(tt.src, tt.dst, tt.ps, test.FileMode); (err != nil) != tt.wantErr {
				t.Errorf("CopyDirToDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, f := range tt.want {
				if !Exist(f) {
					t.Errorf("file not found %v", f)
					return
				}
			}
		})
	}
}

func TestCopyDirToDirSrcDoesNotExist(runner *testing.T) {
	td := test.TmpDir + PS + runner.Name()
	_ = os.MkdirAll(td, 0744)

	tests := []struct {
		name    string
		src     string
		dst     string
		ps      string
		wantErr bool
	}{
		{
			"src error",
			test.FixtureDir + PS + "fake-dir-01",
			td,
			PS,
			true,
		},
	}
	for _, tt := range tests {
		runner.Run(tt.name, func(t *testing.T) {
			if err := CopyDirToDir(tt.src, tt.dst, tt.ps, test.FileMode); (err != nil) != tt.wantErr {
				t.Errorf("CopyDirToDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// The destination (Dst) exist, and it should be overwritten.
func TestCopyDirToDirDstOverwrite(runner *testing.T) {
	td := test.TmpDir + PS + runner.Name()
	_ = os.MkdirAll(td, 0744)

	f := "dir-to-dir-03"

	tests := []struct {
		name    string
		src     string
		dst     string
		ps      string
		wantErr bool
	}{
		{
			"success",
			test.FixtureDir + PS + f,
			td,
			PS,
			false,
		},
	}
	for _, tt := range tests {
		runner.Run(tt.name, func(t *testing.T) {
			d := git.CloneFromBundle(f, td, test.FixtureDir, PS)

			fmt.Printf("d = %v", d)
			if err := CopyDirToDir(tt.src, d, tt.ps, test.FileMode); (err != nil) != tt.wantErr {
				t.Errorf("CopyDirToDir() error = %v, wantErr %v", err, tt.wantErr)
			}

			b, e1 := os.ReadFile(d + PS + ".config/config.yml")
			if e1 != nil {
				t.Errorf("failed to read file: %v", e1.Error())
				return
			}

			if !strings.Contains(string(b), "2.1") {
				t.Errorf("did not get expected content from file.")
				return
			}
		})
	}
}

func TestCopyToDir(t *testing.T) {
	testCases := []struct {
		name, source, dest string
		want               int64
	}{
		{"canCopyFile", test.FixtureDir + "/copy-file-01.txt", test.TmpDir, 26},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CopyToDir(tc.source, tc.dest, "/")

			if err != nil {
				t.Errorf("got an unexpected error copying file %q to %q", tc.source, tc.dest)
				return
			}

			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
				return
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
			return
		}
	}
}
