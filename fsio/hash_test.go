package fsio

import (
	"encoding/base64"
	"testing"
)

func TestSha256(t *testing.T) {
	cases := []struct {
		name     string
		filePath string
		want     string
		wantErr  bool
	}{
		{"hello-hash", "testdata/hi-world.txt", "waaFf8Nx4eylK9s/TraYL26cyS5jQ2W2TB0WJQF2Q4g=", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotHash, err := Sha256(tc.filePath)
			got := base64.StdEncoding.EncodeToString(gotHash)

			if (err != nil) != tc.wantErr {
				t.Errorf("Sha256() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if got != tc.want {
				t.Errorf("Sha256() got = %v, want %v", got, tc.want)
				return
			}
		})
	}
}
