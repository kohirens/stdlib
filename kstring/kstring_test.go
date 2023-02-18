package kstring

import "testing"

func TestStringInArray(tester *testing.T) {
	testCases := []struct {
		name  string
		item  string
		items []string
		want  bool
	}{
		{"canFind", "item1", []string{"item1", "item2"}, true},
		{"cannotFind", "item3", []string{"item1", "item2"}, false},
	}

	for _, tc := range testCases {
		tester.Run(tc.name, func(t *testing.T) {
			if got := InArray(tc.item, tc.items); got != tc.want {
				t.Errorf("got %v, but want %v", got, tc.want)
			}
		})
	}
}

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
