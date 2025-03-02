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
	items := *m.data.Items
	value, ok := items[key]
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
	items := *m.data.Items
	_, ok := items[key]
	if !ok {
		return fmt.Errorf(stderr.NoSuchKey, key)
	}

	// Indicate the session data needs to be saved.
	m.hasUpdates = true

	// Remove the key
	delete(*m.data.Items, key)

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
	items := *m.data.Items
	items[key] = value
}

// Load Will begin a new session, or restore an unexpired session, store the
// session ID in an HTTP cookie to use on the next request.
func (m *Manager) Load(w http.ResponseWriter, r *http.Request) {
	idCookie, _ := r.Cookie(IDKey)

	// ONLY set a new cookie when there is no session, or it has expired.
	if idCookie == nil {
		idCookie = m.IDCookie(IDCookiePath, IDCookieDomain)
		Log.Infof(stdout.IDSet)
		http.SetCookie(w, idCookie)
	}

	if e := m.Restore(idCookie.Value); e != nil {
		Log.Errf(e.Error())
	} else {
		// When we successfully restore a session, we extend it a bit.
		// Have the cookie also reflect this extended time.
		Log.Infof(stdout.Restored)
	}

	// Expire the cookie immediately if the ID does not match (tampering).
	if idCookie.Value != m.ID() {
		Log.Errf(stderr.SessionStrange)
		idCookie.Expires = time.Now().UTC()
		m.RemoveAll()
	}
}

// RemoveAll When you need to scrub the data from the session and fast.
func (m *Manager) RemoveAll() {
	m.data.Items = &Store{}
}
