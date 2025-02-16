package session

var stderr = struct {
	EmptySessionID string
	ExpiredCookie  string
	NoSuchKey      string
}{
	EmptySessionID: "session ID is empty",
	ExpiredCookie:  "session ID cookie has expired on %v",
	NoSuchKey:      "the key %v was not found in the session",
}
