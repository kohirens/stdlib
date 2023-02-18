package stdlib

import (
	"github.com/kohirens/stdlib/internal/test"
	"testing"
)

func TestPathExist(runner *testing.T) {

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
			got := PathExist(sbj.path)

			if got != sbj.want {
				t.Errorf("got %v, want %v", got, sbj.want)
			}
		})
	}
}

func TestIsTextFile(tester *testing.T) {

	cases := []struct {
		name, path string
		want       bool
	}{
		{"IsATextFile", "text-file-01.txt", true},
		{"notATextFile", "text-file-02.jpg", false},
		{"notATextFile", "text-file-03.gif", false},
		{"notATextFile", "text-file-04.png", false},
		{"notATextFile", "text-file-05.json", true},
		{"notATextFile", "text-file-06.md", true},
		{"notATextFile", "text-file-07.xml", true},
	}

	el := []string{"jpg", "gif", "png", "pdf"}
	in := []string{"txt", "json", "md", "xml"}
	sbj, _ := NewFileExtChecker(&el, &in)

	for _, fxtr := range cases {
		tester.Run(fxtr.name, func(t *testing.T) {

			got := sbj.IsValid(fxtr.path)

			if got != fxtr.want {
				t.Errorf("got %v, want %v, for %v", got, fxtr.want, fxtr.path)
			}
		})
	}

	cases = []struct {
		name, path string
		want       bool
	}{
		{"notInTheExcludeList", "text-file-07.xml", true},
	}
	fxtr := cases[0]
	el = []string{"jpg", "gif", "png", "pdf"}
	in = []string{}
	sbj, _ = NewFileExtChecker(&el, &in)
	tester.Run(fxtr.name, func(t *testing.T) {

		got := sbj.IsValid(fxtr.path)

		if got != fxtr.want {
			t.Errorf("got %v, want %v, for %v", got, fxtr.want, fxtr.path)
		}
	})

	cases = []struct {
		name, path string
		want       bool
	}{
		{"notInTheIncludeList", "file-07.jpg", false},
	}
	fxtr = cases[0]
	el = []string{}
	in = []string{"txt", "json", "md", "xml"}
	sbj, _ = NewFileExtChecker(&el, &in)
	tester.Run(fxtr.name, func(t *testing.T) {

		got := sbj.IsValid(fxtr.path)

		if got != fxtr.want {
			t.Errorf("got %v, want %v, for %v", got, fxtr.want, fxtr.path)
		}
	})
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
