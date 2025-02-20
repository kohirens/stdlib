package session

var stderr = struct {
	DecodeJSON     string
	EmptySessionID string
	ExpiredCookie  string
	NoStorage      string
	NoSuchKey      string
}{
	DecodeJSON:     "could not decode json data: %v",
	EmptySessionID: "session ID is empty",
	ExpiredCookie:  "session has expired at %v",
	NoStorage:      "storage has not been set",
	NoSuchKey:      "the key %v was not found in the session",
}
