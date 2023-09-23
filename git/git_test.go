package git

import (
	"github.com/kohirens/stdlib/internal/test"
	"github.com/kohirens/stdlib/path"
	"os"
	"testing"
)

const ps = string(os.PathSeparator)

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
			args{"repo-01", test.TestTmp, test.FixtureDir, ps},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CloneFromBundle(tt.args.bundleName, test.TestTmp, tt.args.bundleDir, tt.args.ps)
			if path.Exist(got) != tt.want {
				t.Errorf("CloneFromBundle() = %v, want %v", got, tt.want)
			}
		})
	}
}
