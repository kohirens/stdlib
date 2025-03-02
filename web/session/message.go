package session

var stderr = struct {
	DecodeJSON,
	EmptySessionID,
	ExpiredCookie,
	NoStorage,
	NoSuchKey,
	SessionStrange string
}{
	DecodeJSON:     "could not decode json data: %v",
	EmptySessionID: "session ID is empty",
	ExpiredCookie:  "session has expired at %v",
	NoStorage:      "storage has not been set",
	NoSuchKey:      "the key %v was not found in the session",
	SessionStrange: "strangeness detected, the session is out of sync. expiring the current session cookie, the user will have to start a new session",
}

var stdout = struct {
	IDSet,
	Restored string
}{
	IDSet:    "setting a session ID cookie now",
	Restored: "session restored",
}
