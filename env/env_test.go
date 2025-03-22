package env

import (
	"bytes"
	"os"
	"reflect"
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
				return
			}
		})
	}
}

func TestParseEnvRE(t *testing.T) {
	cases := []struct {
		name    string
		data    []byte
		want    map[string]string
		wantErr bool
	}{
		{"singleline", []byte("TEST=1234\nTEST2=4321"), map[string]string{"TEST": "1234", "TEST2": "4321"}, false},
		{"multilineVar", []byte("TEST=1234\nTEST2=4321\n1234\n"), map[string]string{"TEST": "1234", "TEST2": "4321\n1234"}, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, e1 := parseEnv(bytes.NewReader(c.data))

			if e1 != nil && !c.wantErr {
				t.Errorf("parseEnv() got err = %v, wanted = %v", e1, c.wantErr)
			}

			if reflect.DeepEqual(got, c.want) == false {
				t.Errorf("got %v; want %v", got, c.want)
				return
			}
		})
	}

}
