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
