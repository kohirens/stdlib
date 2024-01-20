package web

import "github.com/aws/aws-lambda-go/events"

// Response Has standard response fields and should be easily convertable.
type Response struct {
	Body            string            `json:"body"`
	Cookies         []string          `json:"cookies"`
	Headers         map[string]string `json:"headers"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	StatusCode      int               `json:"statusCode"`
}

func (res *Response) ToLambdaResponse() *events.LambdaFunctionURLResponse {
	return &events.LambdaFunctionURLResponse{
		StatusCode:      res.StatusCode,
		Headers:         res.Headers,
		Body:            res.Body,
		IsBase64Encoded: res.IsBase64Encoded,
		Cookies:         res.Cookies,
	}
}
