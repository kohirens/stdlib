package slice

import (
	"bytes"
	"reflect"
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

func TestPrepend(t *testing.T) {
	type testCase[V comparable] struct {
		name string
		ary  []V
		item V
		want []V
	}
	tests := []testCase[int]{
		{"int", []int{2, 3}, 1, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prepend(tt.ary, tt.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prepend() = %v, want %v", got, tt.want)
			}
		})
	}
}
