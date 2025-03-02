package session

import (
	"fmt"
	"net/http"
	"time"
)

// Manager This is the container/interface for your session. Needed to make a
// new session or restore an existing one.
type Manager struct {
	data    *Data
	storage Storage
	// To save on network traffic default this to false, and set to true when
	// Set is called. Save is no-up if this is false.
	hasUpdates bool
}

// Get Retrieve data from the session.
func (m *Manager) Get(key string) []byte {
	value, ok := m.data.Items[key]
	if ok {
		return value
	}

	return nil
}

// Expiration Retrieve expiration time.
func (m *Manager) Expiration() time.Time {
	return m.data.Expiration
}

// ID Of the session as an HTTP cookie with secure and http-only (cannot be read by JavaScript) enabled.
// The domain parameter is optional, and only set when it is not an emptry string.
func (m *Manager) ID() string {
	return m.data.Id
}

func (m *Manager) IDCookie(cookiePath, domain string) *http.Cookie {
	c := &http.Cookie{
		Expires:  m.data.Expiration,
		Name:     IDKey,
		Path:     cookiePath,
		Secure:   true,
		HttpOnly: true,
		Value:    m.data.Id,
		SameSite: http.SameSiteStrictMode,
	}

	if domain != "" {
		c.Domain = domain
	}

	return c
}

// Remove data from a session
func (m *Manager) Remove(key string) error {
	// verify the key exists
	_, ok := m.data.Items[key]
	if !ok {
		return fmt.Errorf(stderr.NoSuchKey, key)
	}

	// Indicate the session data needs to be saved.
	m.hasUpdates = true

	// Remove the key
	delete(m.data.Items, key)

	return nil
}

// Restore Restores the session by ID as a string.
func (m *Manager) Restore(id string) error {
	if id == "" {
		return fmt.Errorf(stderr.EmptySessionID)
	}

	if m.storage == nil {
		return StorageError
	}

	// Load from storage.
	data, e1 := m.storage.Load(id)
	if e1 != nil {
		return e1
	}

	// Verify the session has not expired.
	if time.Now().UTC().After(data.Expiration.UTC()) {
		return ExpiredError{data.Expiration}
	}

	m.data = data
	// extend the session a bit more since data was recently accessed.
	m.data.Expiration.Add(ExtendTime)

	return nil
}

// Save Writes session data to its storage. This is no-op if Set was not previously called.
func (m *Manager) Save() error {
	if m.hasUpdates {
		return m.storage.Save(m.data)
	}

	return nil
}

// Set Store data in the session.
func (m *Manager) Set(key string, value []byte) {
	m.hasUpdates = true
	m.data.Items[key] = value
}

// SetSessionIDCookie Set the session ID cookie. Will only set if there is no
// existing cookie or current session expires.
func (m *Manager) SetSessionIDCookie(w http.ResponseWriter, r *http.Request) {
	sidCookie, _ := r.Cookie(IDKey)

	if sidCookie != nil && !time.Now().UTC().After(sidCookie.Expires) {
		Log.Infof(stdout.SessionExist, sidCookie.Value)
		// Expire the cookie immediately if the ID does not match.
		if sidCookie.Value != m.ID() {
			Log.Errf(stderr.SessionStrange)
			sidCookie.Expires = time.Now().UTC()
		}
		return
	}

	// ONLY set a new cookie when there is no session, or it has expired.
	Log.Infof(stdout.SessionSet)
	sidCookie = m.IDCookie("/", "")
	http.SetCookie(w, sidCookie)
}
