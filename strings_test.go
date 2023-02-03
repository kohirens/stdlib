package stdlib

import "testing"

func TestStrToCamel(runner *testing.T) {
	cases := []struct {
		name,
		subject,
		separator string
		pascal bool
		want   string
	}{
		{"success", "test-me", "-", false, "testMe"},
		{"capitalized", "Test-me", "-", false, "testMe"},
		{"snake", "Test_me_Too", "_", false, "testMeToo"},
		{"pascal", "me_three", "_", true, "MeThree"},
	}

	for _, tc := range cases { // tc stands for test case
		runner.Run(tc.name, func(t *testing.T) {
			got, err := StrToCamel(tc.subject, tc.separator, tc.pascal)

			if err != nil {
				t.Errorf("got an unexpected error: %v", err.Error())
			}

			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
