package stdlib

import (
	"testing"
)

func TestIsTextFile(tester *testing.T) {
	cases := []struct {
		name, path string
		want       bool
	}{
		{"txt", "text-file-01.txt", true},
		{"jpg", "text-file-02.jpg", false},
		{"gif", "text-file-03.gif", false},
		{"png", "text-file-04.png", false},
		{"json", "text-file-05.json", true},
		{"md", "text-file-06.md", true},
		{"xml", "text-file-07.xml", true},
		{"hiddenDirectory", ".hidden/file.txt", true},
		{"noExtension", "config", true},
	}

	el := []string{"jpg", "gif", "png", "pdf"}
	in := []string{"txt", "json", "md", "xml", "config"}
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
}
