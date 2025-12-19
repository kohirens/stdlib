package fsio

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadByJson(t *testing.T) {
	type mockFixture struct {
		Test int `json:"test"`
	}
	tests := []struct {
		name     string
		filename string
		v        *mockFixture
		want     *mockFixture
		wantErr  bool
	}{
		{
			name:     "load structure by json",
			filename: "testdata/load_by.json",
			v:        &mockFixture{},
			want: &mockFixture{
				Test: 1234,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadByJson(tt.filename, tt.v); (err != nil) != tt.wantErr {
				t.Errorf("LoadByJson() error = %v, wantErr %v", err, tt.wantErr)
			}
			if reflect.DeepEqual(tt.v, tt.want) == false {
				t.Errorf("LoadByJson() got = %v, want %v", tt.v, tt.want)
			}
		})
	}
}

func ExampleLoadByJson() {
	type url struct {
		Host string `json:"host"`
		Path string `json:"path"`
	}
	u := new(url)

	if e := LoadByJson("testdata/url.json", u); e != nil {
		fmt.Println("error:", e)
		return
	}
	fmt.Println(u.Host + "/" + u.Path)
	// Output:
	// example.com/api/greeting
}
