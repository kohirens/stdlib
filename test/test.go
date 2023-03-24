package test

import (
	"net/http"
	"os"
)

type MockHttpClient struct {
	DoHandler  func(r *http.Request) (*http.Response, error)
	GetHandler func(url string) (*http.Response, error)
}

func (m *MockHttpClient) Do(r *http.Request) (*http.Response, error) {
	return m.DoHandler(r)
}

func (m *MockHttpClient) Get(url string) (*http.Response, error) {
	return m.GetHandler(url)
}

func GetHttpResponseFromFile(filepath string) (*http.Response, error) {
	f, e1 := os.Open(filepath)
	if e1 != nil {
		return nil, e1
	}
	return &http.Response{Body: f}, nil
}
