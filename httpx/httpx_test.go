package httpx

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/kohirens/stdlib/test"
)

const (
	fixtureDir = "testdata"
	//tmpDir     = "tmp"
)

//func TestMain(m *testing.M) {
//	test.ResetDir(tmpDir, 0777)
//
//	os.Exit(m.Run())
//}

// HttpClient Methods needed to make HTTP request.
type MockClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func TestSendWithRetry(t *testing.T) {
	tests := []struct {
		name    string
		uri     string
		client  MockClient
		wantErr bool
	}{
		{
			"200_response",
			"https://test.local/oauth2/v3/token",
			&test.MockHttpClient{
				DoHandler: func(r *http.Request) (*http.Response, error) {
					b, _ := os.ReadFile(fixtureDir + "/test-token-01.json")
					return &http.Response{
						Body:       io.NopCloser(bytes.NewReader(b)),
						StatusCode: 200,
					}, nil
				},
			},
			false,
		},
		{
			"200_response",
			"https://test.local/oauth2/v3/token",
			&test.MockHttpClient{
				DoHandler: func(r *http.Request) (*http.Response, error) {
					b, _ := os.ReadFile(fixtureDir + "/test-token-01.json")
					return &http.Response{
						Body:       io.NopCloser(bytes.NewReader(b)),
						StatusCode: 404,
					}, nil
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SendWithRetry(
				tt.client,
				"POST",
				tt.uri,
				[]byte{},
				map[string][]string{},
				[]int{http.StatusOK},
				3,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendWithRetry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
