package httpx

var stderr = struct {
	BuildRequest,
	NoResponse,
	RetryRequest,
	UnexpectedCode string
}{
	BuildRequest:   "cannot build the request: %v",
	NoResponse:     "no response",
	RetryRequest:   "request with retry %v",
	UnexpectedCode: "attempt %v to url %v has returned HTTP status code %v with body %v",
}
