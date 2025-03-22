package fsio

import (
	"github.com/kohirens/stdlib"
	"github.com/kohirens/stdlib/internal/test"
	"testing"
)

func TestExist(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{"yes", "./", true},
		{"no", "./does-not-exist", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exist(tt.filename); got != tt.want {
				t.Errorf("Exist() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
func TestExistFiles(runner *testing.T) {

	cases := []struct {
		name, path string
		want       bool
	}{
		{"existIsTrue", test.FixtureDir + "/file-exist-01.md", true},
		{"existIsFalse", test.FixtureDir + "/file-does-not-exist-01.md", false},
		{"invalidPathIsFalse", test.FixtureDir + "https://github.com/kohirens/tmpl-go-web/archive/refs/heads/main.zip\\template.json", false},
	}

	for _, sbj := range cases {
		runner.Run(sbj.name, func(t *testing.T) {
			got := Exist(sbj.path)

			if got != sbj.want {
				t.Errorf("got %v, want %v", got, sbj.want)
				return
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
		{"windowsOnUnix", "\\windows\\path\\on\\uUnix", stdlib.PS + "windows" + stdlib.PS + "path" + stdlib.PS + "on" + stdlib.PS + "uUnix"},
		{"unixOnWindows", "/unix/on/windows", stdlib.PS + "unix" + stdlib.PS + "on" + stdlib.PS + "windows"},
		{"mixedBag", "/i\\dont\\care/really", stdlib.PS + "i" + stdlib.PS + "dont" + stdlib.PS + "care" + stdlib.PS + "really"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Normalize(tc.path)

			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
				return
			}
		})
	}
}
