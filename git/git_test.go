package git

import (
	"fmt"
	"github.com/kohirens/stdlib/fsio"
	"github.com/kohirens/stdlib/test"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	FixtureDir = "testdata"
	ps         = string(os.PathSeparator)
	TmpDir     = "tmp"
)

func TestMain(t *testing.M) {
	test.ResetDir(TmpDir, 0774)
	os.Exit(t.Run())
}

func TestCloneFromBundle(t *testing.T) {
	type args struct {
		bundleName string
		tmpDir     string
		bundleDir  string
		ps         string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"bundle-eixists",
			args{"repo-01", TmpDir, FixtureDir, ps},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CloneFromBundle(tt.args.bundleName, TmpDir, tt.args.bundleDir, tt.args.ps)
			if fsio.Exist(got) != tt.want {
				t.Errorf("CloneFromBundle() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Clone a repo
func TestGitClone(tester *testing.T) {
	var testCases = []struct {
		name      string
		repo      string
		outPath   string
		branch    string
		shouldErr bool
		wantHash  string
	}{
		{
			"cloneRepo1",
			"repo-01.git",
			TmpDir + ps + "repo-01-refs-heads-main",
			"refs/heads/main",
			false,
			"b7e42844c597d2beaf774eddfdcb653a2a4b0050",
		},
	}

	for _, tc := range testCases {
		tester.Run(tc.name, func(t *testing.T) {
			repoPath := CloneFromBundle(tc.repo, TmpDir, FixtureDir, ps)

			gotPath, gotHash, err := Clone(repoPath, tc.outPath, tc.branch)

			if tc.shouldErr == true && err == nil {
				t.Error("did not get expected err")
			}

			if tc.shouldErr == false && err != nil {
				t.Errorf("got an unexpected err: %s", err)
			}

			if gotHash != tc.wantHash {
				t.Errorf("got %v, want %v", gotHash, tc.wantHash)
			}

			if gotPath != tc.outPath {
				t.Errorf("got %v, want %v", gotPath, tc.outPath)
			}
		})
	}
}

// Clone a repo
func TestGitCannotClone(tester *testing.T) {
	var testCases = []struct {
		name      string
		repo      string
		outPath   string
		branch    string
		shouldErr bool
		wantHash  string
	}{
		{"clone404", "does-not-exist.git", "", "", true, ""},
	}

	for _, tc := range testCases {
		tester.Run(tc.name, func(t *testing.T) {
			repoPath, _ := filepath.Abs(TmpDir)
			repoPath += ps + tc.repo

			gotPath, gotHash, err := Clone(repoPath, tc.outPath, tc.branch)

			if tc.shouldErr == true && err == nil {
				t.Error("did not get expected err")
			}

			if tc.shouldErr == false && err != nil {
				t.Errorf("got an unexpected err: %s", err)
			}

			if gotHash != tc.wantHash {
				t.Errorf("got %v, want %v", gotHash, tc.wantHash)
			}

			if gotPath != tc.outPath {
				t.Errorf("got %v, want %v", gotPath, tc.outPath)
			}
		})
	}
}

func TestGetRepoDir2(tester *testing.T) {
	absTestTmp, _ := filepath.Abs(TmpDir)
	var testCases = []struct {
		name     string
		bundle   string
		branch   string
		want     string
		wantHash string
	}{
		{"localFullBranchRefArg", "repo-02", "refs/remotes/origin/third-commit", absTestTmp + ps + "repo-02", "bfeb0a45c027420e4df286dc089965599e350bf9"},
	}

	for _, tc := range testCases {
		repoPath := CloneFromBundle(tc.bundle, TmpDir, FixtureDir, ps)

		tester.Run(tc.name, func(t *testing.T) {
			gotRepo, gotHash, gotErr := Checkout(repoPath, tc.branch)

			if gotErr != nil {
				t.Errorf("unexpected error in test %q", gotErr.Error())
			}

			if gotRepo != tc.want {
				t.Errorf("got %v, want %v", gotRepo, tc.want)
			}

			if gotHash != tc.wantHash {
				t.Errorf("got %v, want %v", gotHash, tc.wantHash)
			}
		})
	}
}

func TestGetLatestTag(tester *testing.T) {
	var testCases = []struct {
		name   string
		bundle string
		want   string
	}{
		{"found", "repo-03", "0.1.0"},
	}

	for i, tc := range testCases {
		repoPath := CloneFromBundle(tc.bundle, TmpDir, FixtureDir, ps)

		tester.Run(fmt.Sprintf("%v.%v", i+1, tc.name), func(t *testing.T) {
			got, gotErr := LatestTag(repoPath)

			if gotErr != nil {
				t.Errorf("unexpected error in test %q", gotErr.Error())
			}

			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGetLatestTagError(tester *testing.T) {
	var testCases = []struct {
		name   string
		bundle string
		want   string
	}{
		{"doesNotExist", "repo-dne", ""},
	}

	for i, tc := range testCases {
		repoPath, _ := filepath.Abs("tmp")
		repoPath += ps + tc.bundle
		tester.Run(fmt.Sprintf("%v.%v", i+1, tc.name), func(t *testing.T) {
			got, gotErr := LatestTag(repoPath)

			if gotErr == nil {
				t.Errorf("unexpected error in test %q", gotErr.Error())
			}

			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGetRemoteTags(tester *testing.T) {
	var testCases = []struct {
		name      string
		bundle    string
		want      string
		shouldErr bool
	}{
		{"hasTags", "repo-04", "1.0.0,0.2.0,0.1.1,0.1.0", false},
		{"noTags", "repo-05", "", true},
	}

	for i, tc := range testCases {
		repoPath := CloneFromBundle(tc.bundle, TmpDir, FixtureDir, ps)

		tester.Run(fmt.Sprintf("%v.%v", i+1, tc.name), func(t *testing.T) {
			got, gotErr := RemoteTags(repoPath)

			if !tc.shouldErr && gotErr != nil {
				t.Errorf("unexpected error in test %q", gotErr.Error())
			}

			t1 := strings.Join(got, ",")
			if got != nil && t1 != tc.want {
				t.Errorf("got %v, want %v", t1, tc.want)
			}
		})
	}
}
