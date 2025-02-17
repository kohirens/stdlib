package session

var stderr = struct {
	DecodeJSON     string
	EmptySessionID string
	ExpiredCookie  string
	NoSuchKey      string
}{
	DecodeJSON:     "could not decode json data: %v",
	EmptySessionID: "session ID is empty",
	ExpiredCookie:  "session ID cookie has expired on %v",
	NoSuchKey:      "the key %v was not found in the session",
}
