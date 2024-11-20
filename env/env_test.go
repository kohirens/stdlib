package env

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	fixtureEnvName := "TEST_FIXTURE_1"
	_ = os.Setenv(fixtureEnvName, "1234")
	defer func() {
		_ = os.Unsetenv(fixtureEnvName)
	}()

	cases := []struct {
		name string
		key  string
		def  string
		want string
	}{
		{"canGetEnvironmentValue", fixtureEnvName, "abc", "1234"},
		{"environmentValueNotSet", fixtureEnvName + "_", "abc", "abc"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Get(c.key, c.def); got != c.want {
				t.Errorf("Get() = %v, want %v", got, c.want)
			}
		})
	}
}
