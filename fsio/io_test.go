package fsio

import (
	"io"
	"testing"
)

func TestReadCloser(t *testing.T) {
	tests := []struct {
		name    string
		content string
		isFile  bool
		want    string
		wantErr bool
	}{
		{"can-read-file", "testdata/salam.txt", true, "Salam", false},
		{"can-read-file", "Salam", false, "Salam", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewReadCloser(tt.content, tt.isFile)

			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCloser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotBytes, _ := io.ReadAll(got)
			if string(gotBytes) != tt.want {
				t.Errorf("ReadCloser() got = %v, want %v", string(gotBytes), tt.want)
			}
		})
	}
}
