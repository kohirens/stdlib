package slice

import (
	"bytes"
	"testing"
)

func TestPrependByte(tr *testing.T) {
	testCases := []struct {
		name string
		item byte
		ary  []byte
		want []byte
	}{
		{"abc", byte('a'), []byte("bc"), []byte("abc")},
	}

	for _, tc := range testCases {
		tr.Run(tc.name, func(t *testing.T) {
			got := PrependByte(tc.item, tc.ary)

			if bytes.Compare(got, tc.want) != 0 {
				t.Errorf("got %v but want %v", got, tc.want)
			}
		})
	}
}

func TestPrependInt(tr *testing.T) {
	testCases := []struct {
		name string
		item int
		ary  []int
		want []int
	}{
		{"123", 1, []int{2, 3}, []int{1, 2, 3}},
	}

	for _, tc := range testCases {
		tr.Run(tc.name, func(t *testing.T) {
			got := PrependInt(tc.item, tc.ary)

			for i, v := range got {
				if v != tc.want[i] {
					t.Errorf("got %v but want %v", got, tc.want)
					break
				}
			}
		})
	}
}
