# httpx

Added a method to send an HTTP request and also retry a specified
number of times when the request does not return a specified HTTP
status code.

```go
package main

import (
	"net/http"

	"io"

	"github.com/kohirens/stdlib/httpx"
)

func main() {
	body := []byte{}
	headers := map[string][]string{}
	codes := []int{http.StatusOK, http.StatusNoContent}
	// The idea here is to retry more than once and use the status code
	// to determine success.
	res, err := httpx.SendWithRetry(
		&http.Client{},
		"POST",
		"https://example.com/",
		body,
		headers,
		codes,
		3, // retry at least 3 times.
	)
	if err != nil {
		panic(err.Error())
	}
	
	body := io.ReadAll(res.Body)
	_ = res.Body.Close()
}
```