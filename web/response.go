package web

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"strings"
)

// Response Serves as a middle ground between types such as http.Response and events.LambdaFunctionURLResponse;
// with methods to easily convert to each type. In addition, it also works as a http.ResponseWriter.
type Response struct {
	Body            string              `json:"body"`
	Headers         map[string][]string `json:"headers"`
	IsBase64Encoded bool                `json:"isBase64Encoded"`
	StatusCode      int                 `json:"statusCode"`
}

// ToLambdaResponse Convert to a Lambda function URL response.
func (res *Response) ToLambdaResponse() *events.LambdaFunctionURLResponse {
	return &events.LambdaFunctionURLResponse{
		StatusCode:      res.StatusCode,
		Headers:         res.convertToLambdaHeaders(),
		Body:            res.Body,
		IsBase64Encoded: res.IsBase64Encoded,
		Cookies:         res.Cookies(),
	}
}

// Cookies Get all the cookies.
func (res *Response) Cookies() []string {
	return res.Headers["Set-Cookie"]
}

// Base64Encode Encodes the response body.
func (res *Response) Base64Encode() {
	res.IsBase64Encoded = true
	res.Body = base64.StdEncoding.EncodeToString([]byte(res.Body))
}

// ToHttpResponse Convert to an HTTP response.
func (res *Response) ToHttpResponse() *http.Response {
	r := &http.Response{
		StatusCode: res.StatusCode,
		Header:     res.Headers,
	}

	b := bytes.NewBufferString(res.Body)
	if e := r.Write(b); e != nil {
		panic(e)
	}

	return r
}

// Convert the http.Response style headers map[string][]string to map[string]string.
func (res *Response) convertToLambdaHeaders() map[string]string {
	headers := make(map[string]string, len(res.Headers))
	for k, v := range res.Headers {
		headers[k] = strings.Join(v, "\n")
	}
	return headers
}

// Header Part of the http.ResponseWriter interface.
func (res *Response) Header() http.Header {
	return res.Headers
}

// WriteHeader Part of the http.ResponseWriter interface.
func (res *Response) Write(b []byte) (int, error) {
	res.Body = string(b)
	return len(res.Body), nil
}

// WriteHeader Part of the http.ResponseWriter interface.
func (res *Response) WriteHeader(statusCode int) {
	res.StatusCode = statusCode
}
