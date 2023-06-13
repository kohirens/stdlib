package cli

import (
	"testing"
)

func TestAmendStringAry(t *testing.T) {
	tests := []struct {
		name string
		ce   []string
		env  map[string]string
		want []string
	}{
		{
			"overwriteAndAmend",
			[]string{"GOARCH=amd64", "GOOS=windows"},
			map[string]string{"GOOS": "linux", "test": "1234"},
			[]string{"GOARCH=amd64", "GOOS=linux", "test=1234"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AmendStringAry(tt.ce, tt.env)
			if got[0] != tt.want[0] {
				t.Errorf("AmendStringAry()[0] = %v, want %v", got[0], tt.want[0])
			}
			if got[1] != tt.want[1] {
				t.Errorf("AmendStringAry()[1] = %v, want %v", got[1], tt.want[1])
			}
			if got[2] != tt.want[2] {
				t.Errorf("AmendStringAry()[2] = %v, want %v", got[2], tt.want[2])
			}
		})
	}
}
