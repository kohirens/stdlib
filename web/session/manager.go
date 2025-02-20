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

// ID Of the session as an HTTP cookie with secure and http-only (cannot be read by JavaScript) enabled.
// The domain parameter is optional, and only set when it is not an emptry string.
func (m *Manager) ID(cookiePath, domain string) *http.Cookie {
	c := &http.Cookie{
		Expires:  m.data.Expiration,
		Name:     IDKey,
		Path:     cookiePath,
		Secure:   true,
		HttpOnly: true,
		Value:    m.data.Id,
	}

	if domain != "" {
		c.Domain = domain
	}

	return c
}

// Restore Restores the session by ID as a string.
func (m *Manager) Restore(id string) error {
	if id == "" {
		return fmt.Errorf(stderr.EmptySessionID)
	}

	// Load from storage.
	data, e1 := m.storage.Load(id)
	if e1 != nil {
		return e1
	}

	m.data = data

	return nil
}

// RestoreFromCookie Restore a session by ID from an HTTP cookie as long as the
// cookie has not expired and the data can be pulled from storage.
func (m *Manager) RestoreFromCookie(sidCookie *http.Cookie) error {
	if sidCookie == nil || sidCookie.Value == "" {
		return fmt.Errorf(stderr.EmptySessionID)
	}

	// Verify the cookie has not expired.
	if !sidCookie.Expires.IsZero() && time.Now().UTC().After(sidCookie.Expires.UTC()) {
		return ExpiredError{sidCookie.Expires}
	}

	// Load the session from storage, no matter the storage medium this should always return JSON as a byte array.
	data, e1 := m.storage.Load(sidCookie.Value)
	if e1 != nil {
		return e1
	}

	m.data = data
	// extend the session a bit more since data was recently accessed.
	m.data.Expiration.Add(ExtendTime)

	return nil
}

// NewManager Initialize a new session manager to handle session save, restore, get, and set.
func NewManager(storage Storage, expiration time.Duration) *Manager {
	return &Manager{
		data: &Data{
			GenerateID(),
			time.Now().Add(expiration),
			make(Store, 100),
		},
		storage:    storage,
		hasUpdates: false,
	}
}

// Save Writes session data to its storage. This is no-op if Set was not previously called.
func (m *Manager) Save() error {
	if m.hasUpdates {
		return m.storage.Save(m.data)
	}

	return nil
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

// Set Store data in the session.
func (m *Manager) Set(key string, value []byte) {
	m.hasUpdates = true
	m.data.Items[key] = value
}
