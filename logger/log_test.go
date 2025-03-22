package logger

import (
	"os"
	"testing"
)

func Test_verboseF(t *testing.T) {
	nf, _ := os.CreateTemp(os.TempDir(), "stdlib")
	defer nf.Close()

	os.Stdout = nf

	tests := []struct {
		name string
	}{
		{"verboseF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			verboseF(-1, "%v", "hi")

			got := make([]byte, 3)
			want := "hi\n"

			_, _ = nf.ReadAt(got, 0)

			if string(got) != want {
				t.Errorf("got %s, want %s", got, want)
				return
			}
		})
	}
}
