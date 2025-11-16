package httpx

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// HttpClient Methods needed to make HTTP request.
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// SendWithRetry Make an HTTP request, retrying up to so many times.
// NOTE: A response is only returned on success. Otherwise, it will be nil when
// the expected status code is not met; so be careful to set the correct codes.
//
// Also, an error can be returned even on success. This is to expose any
// unexpected status code(s) that may have not been accounted for. So any
// responses caused by previous attempts are combined into a single error and
// returned; along with a possible http.Response object given at least 1 attempt
// succeeds.
// The idea here is to retry more than once and use the status code to determine
// success.
func SendWithRetry(
	httpClient HttpClient,
	method, url string,
	data []byte,
	headers http.Header,
	codes []int,
	retries int,
) (*http.Response, error) {
	body := bytes.NewBuffer(data)

	req, e1 := http.NewRequest(method, url, body)
	if e1 != nil {
		return nil, fmt.Errorf(stderr.BuildRequest, e1.Error())
	}

	req.Header = headers
	var lastResponse *http.Response
	var errMessage string

	for attempt := 1; attempt <= retries; attempt++ {
		res, err := httpClient.Do(req)
		if err != nil {
			errMessage += fmt.Sprintf(stderr.RetryRequest, err.Error())
			continue
		}

		if IntInArray(res.StatusCode, codes) {
			lastResponse = res
			break
		}

		// condition where response is not the expected status code but err is
		// nil
		// We throw this back at the app to let the dev know they should handle
		// this particular case.
		resBody, _ := io.ReadAll(res.Body)
		_ = res.Body.Close()

		errMessage += fmt.Sprintf(stderr.UnexpectedCode, attempt, url, res.StatusCode, string(resBody))

		res = nil
	}

	if errMessage != "" {
		return lastResponse, fmt.Errorf("%v", errMessage)
	}

	return lastResponse, nil
}

func IntInArray(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
