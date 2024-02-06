package session

var stderr = struct {
	EmptySessionID string
	ExpiredCookie  string
}{
	EmptySessionID: "session ID is empty",
	ExpiredCookie:  "session ID cookie has expired on %v",
}
