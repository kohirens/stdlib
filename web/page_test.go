package web

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestGetHeader(t *testing.T) {
	tests := []struct {
		name    string
		headers map[string]string
		header  string
		want    string
	}{
		{"get-lowercase-host-header", map[string]string{"Host": "example.com"}, "host", "example.com"},
		{"get-uppercase-host-header", map[string]string{"HOST": "example.com"}, "host", "example.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetHeader(tt.headers, tt.header)

			if got != tt.want {
				t.Errorf("GetHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPageType(t *testing.T) {
	tests := []struct {
		name    string
		headers map[string]string
		want    string
	}{
		{"html-type", map[string]string{"content-type": "text/html"}, "text/html"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPageType(tt.headers); got != tt.want {
				t.Errorf("GetPageType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPageTypeByExt(t *testing.T) {
	tests := []struct {
		name     string
		pagePath string
		want     string
	}{
		{"html", "page.html", "text/html;charset=utf-8"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetPageTypeByExt(tt.pagePath)

			if got != tt.want {
				t.Errorf("GetPageTypeByExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRespond301Or308(t *testing.T) {
	fixedResponse := &events.LambdaFunctionURLResponse{
		StatusCode: 301,
		Headers: map[string]string{
			"Content-Type": "text/html;charset=utf-8",
			"Location":     "https://www.example.com",
		},
	}
	fixed302Response := &events.LambdaFunctionURLResponse{
		StatusCode: 302,
		Headers: map[string]string{
			"Content-Type": "text/html;charset=utf-8",
			"Location":     "https://www.example.com",
		},
	}
	fixed308Response := &events.LambdaFunctionURLResponse{
		StatusCode: 308,
		Headers: map[string]string{
			"Content-Type": "text/html;charset=utf-8",
			"Location":     "https://www.example.com",
		},
	}
	tests := []struct {
		name     string
		call     func(string) *Response
		location string
		want     *events.LambdaFunctionURLResponse
		status   string
	}{
		{"301", Respond301, "https://www.example.com", fixedResponse, "Moved Permanently"},
		{"302", Respond302, "https://www.example.com", fixed302Response, "Found"},
		{"308", Respond308, "https://www.example.com", fixed308Response, "Permanent Redirect"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.call(tt.location)

			if got.StatusCode != tt.want.StatusCode {
				t.Errorf("Respond%v() = %v, want %v", tt.name, got.StatusCode, tt.want.StatusCode)
			}

			if !strings.Contains(got.Body, tt.status) {
				t.Errorf("Respond%v() = does not contain %v", tt.name, tt.status)
			}

			if !reflect.DeepEqual(got.Headers, tt.want.Headers) {
				t.Errorf("Respond%v() = %v, want %v", tt.name, got.Headers, tt.want.Headers)
			}
		})
	}
}

// This suite of test ensure any refactoring of these methods leave the
// required HTTP status code and recommended status message are left intact.
// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401
// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404
// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500
// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/501
func TestRespondXXX(t *testing.T) {
	tests := []struct {
		call       func() *Response
		wantCode   int
		wantStatus string
	}{
		{Respond401, 401, http.StatusText(401)},
		{Respond404, 404, http.StatusText(404)},
		{Respond500, 500, http.StatusText(500)},
		{Respond501, 501, http.StatusText(501)},
	}
	for _, tt := range tests {
		t.Run(tt.wantStatus, func(t *testing.T) {
			got := tt.call()

			if got.StatusCode != tt.wantCode {
				t.Errorf("Respond%v() = %v, want %v", tt.wantCode, got.StatusCode, tt.wantCode)
			}

			if !strings.Contains(got.Body, tt.wantStatus) {
				t.Errorf("Respond%v() = does not contain %v", tt.wantCode, tt.wantStatus)
			}
		})
	}
}

// This suite of test ensure any refactoring of these methods leave the
// required HTTP status code and recommended status message are left intact.
// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405
func TestRespond405(t *testing.T) {
	tests := []struct {
		name       string
		methods    string
		wantCode   int
		wantStatus string
	}{
		{"405", "GET, HEAD, POST,", 405, "405"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Respond405(tt.methods)

			if got.StatusCode != tt.wantCode {
				t.Errorf("Respond%v() = %v, want %v", tt.name, got.StatusCode, tt.wantCode)
			}

			if !strings.Contains(got.Body, tt.wantStatus) {
				t.Errorf("Respond%v() = does not contain %v", tt.name, tt.wantStatus)
			}
		})
	}
}

func TestRespondJSONOG(t *testing.T) {
	type jsonMsg struct {
		Msg string `json:"msg"`
	}

	fixedBody := &jsonMsg{"Salam"}

	tests := []struct {
		name     string
		content  *jsonMsg
		wantBody string
		wantErr  bool
	}{
		{"can-encode", fixedBody, `{"msg":"Salam"}`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, e := RespondJSON(tt.content)

			if (e != nil) != tt.wantErr {
				t.Errorf("RespondJSON() = %v, want %v", e, tt.wantErr)
			}

			if got.Body != tt.wantBody {
				t.Errorf("RespondJSON() = %v, want %v", got.Body, tt.wantBody)
			}
		})
	}
}

func TestRespondDebug(t *testing.T) {
	tests := []struct {
		name    string
		message string
		footer  string
		code    int
	}{
		{"Debug200", "status ok", "Acme", 401},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RespondDebug(tt.code, tt.message, tt.footer)

			if !strings.Contains(got.Body, fmt.Sprintf("%v", tt.code)) {
				t.Errorf("RespondDebug() = does not contain %v", tt.code)
			}

			if !strings.Contains(got.Body, fmt.Sprintf("%v", tt.message)) {
				t.Errorf("RespondDebug() = does not contain %v", tt.footer)
			}

			if !strings.Contains(got.Body, fmt.Sprintf("%v", tt.footer)) {
				t.Errorf("RespondDebug() = does not contain %v", tt.footer)
			}
		})
	}
}
