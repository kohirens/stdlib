package web

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"io"
	"net/http"
	"strings"
)

// Request Serves as a medium between the different types of HTTP request that
// you may run into. Initially this works with Go's HTTP request and AWS Lambda
// function URL request. More will be added as encountered; or until a better
// solution is found.
// The main goal is to always use http.Request no matter what type of Request
// you are given, so that you code across projects is more consistent and highly
// reusable.
// This is just a wrapper around http.Request, as that is used under the hood.
// When you need another type use any of the Request.To* methods to convert to
// another type that is supported.
// While the main goal it not so server as a way to convert one type to another
// it does work out that way.
type Request struct {
	Request                  *http.Request
	LambdaFunctionURLRequest *events.LambdaFunctionURLRequest
}

// NewRequest Wrap an http.Request.
func NewRequest(r *http.Request) *Request {
	return &Request{
		Request: r,
	}
}

// NewRequestFromLambdaFunctionURLRequest Work with this type of request as though it were of type http.Request.
func NewRequestFromLambdaFunctionURLRequest(l *events.LambdaFunctionURLRequest) (*Request, error) {
	origin := GetHeader(l.Headers, "Origin")
	uRL := fmt.Sprintf(
		"%v%v?%v",
		origin,
		l.RawPath,
		l.RawQueryString,
	)

	bdy := convertBody(l.Body, l.IsBase64Encoded)
	r, e2 := http.NewRequest(l.RequestContext.HTTP.Method, uRL, bdy)
	if e2 != nil {
		return nil, fmt.Errorf(Stderr.NewRequest, e2)
	}

	r.Header = convertToHttpHeaders(l.Headers)

	return &Request{
		Request:                  r,
		LambdaFunctionURLRequest: l,
	}, nil
}

// ToLambdaFunctionURLRequest Get the Lambda function URL request you put in or a new one with properties set.
func (r *Request) ToLambdaFunctionURLRequest() *events.LambdaFunctionURLRequest {
	if r.LambdaFunctionURLRequest == nil {
		return r.LambdaFunctionURLRequest
	}

	l := &events.LambdaFunctionURLRequest{
		Cookies: convertCookiesToStringArray(r.Request.Cookies()),
	}

	return l
}

// getBody Convert string body to io.Reader
func convertBody(body string, isBase64 bool) io.Reader {
	if body == "" {
		return nil
	}

	if isBase64 {
		b, e1 := base64.StdEncoding.DecodeString(body)
		if e1 != nil {
			panic(fmt.Errorf(Stderr.DecodeBase64, e1))
		}

		return bytes.NewReader(b)
	}

	return strings.NewReader(body)
}

// convertCookiesToStringArray Convert http.Request.Cookies() to []string
// cookies.
// Returns an empty non-nil slice if there are no cookies in the request.
func convertCookiesToStringArray(rcs []*http.Cookie) []string {
	cookies := make([]string, len(rcs))

	if len(rcs) == 0 {
		return cookies
	}

	for i, cookie := range rcs {
		cookies[i] = cookie.String()
	}

	return cookies
}

// convertToHttpHeaders Convert a map of strings to http.Header's.
func convertToHttpHeaders(lh map[string]string) http.Header {
	headers := http.Header{}

	if len(lh) == 0 {
		return headers
	}

	for name, value := range lh {
		headers[name] = strings.Split(value, "\n")
	}

	return headers
}
