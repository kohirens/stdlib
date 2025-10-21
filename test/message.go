package test

var stderr = struct {
	ExpectedBytes,
	ExpectedStatusCode string
}{
	ExpectedBytes:      "expected bytes found in the body",
	ExpectedStatusCode: "unexpected status code",
}
