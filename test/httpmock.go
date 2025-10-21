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

type MockResponseWriter struct {
	ExpectedBody       []byte
	ExpectedHeaders    http.Header
	Headers            http.Header
	ExpectedStatusCode int
}

func (m MockResponseWriter) Header() http.Header {
	if m.Headers == nil {
		m.Headers = make(http.Header)
	}
	return m.Headers
}

func (m MockResponseWriter) Write(bytes []byte) (int, error) {
	idx := 0
	if m.ExpectedBody != nil {
		var val byte
		for idx, val = range bytes {
			if m.ExpectedBody[idx] != val {
				panic(stderr.ExpectedBytes)
			}
		}
	}
	return idx, nil
}

func (m MockResponseWriter) WriteHeader(statusCode int) {
	if statusCode != m.ExpectedStatusCode {
		panic(stderr.ExpectedStatusCode)
	}
}
