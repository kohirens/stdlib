package stdlib

import (
	"testing"
)

const (
	FIXTURES_DIR = "testdata"
	TEST_TMP     = "testtmp"
)

func TestPathExist(t *testing.T) {

	cases := []struct {
		name, path string
		want       bool
	}{
		{"existIsTrue", FIXTURES_DIR + "/file-exist-01.md", true},
		{"existIsTrue", FIXTURES_DIR + "/file-does-not-exist-01.md", false},
	}

	for _, sbj := range cases {
		got := PathExist(sbj.path)

		if got != sbj.want {
			t.Errorf("got %v, want %v", got, sbj.want)
		}
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
		{"dirExist", FIXTURES_DIR + "/dir_that_exist", true},
		{"isFileNotDir", FIXTURES_DIR + "/dir_that_exist/file_that_exists.md", false},
		{"doesNotExists", FIXTURES_DIR + "/dir_that_exist/dir-does-not-exist", false},
	}

	for _, sbj := range cases {
		got := DirExist(sbj.path)

		if got != sbj.want {
			t.Errorf("got %v, want %v", got, sbj.want)
		}
	}
}