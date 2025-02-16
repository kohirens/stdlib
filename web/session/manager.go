package session

import (
	"fmt"
	"github.com/kohirens/stdlib/log"
	"net/http"
	"time"
)

// Manager This is the container/interface for your session. Needed to make a
// new session or restore an existing one.
type Manager struct {
	cookiePath string
	data       Store
	expires    time.Time
	id         string
	storage    Storage
	// To save on network traffic default this to false, and set to true when
	// Set is called. Save is no-up if this is false.
	hasUpdates bool
	log        log.Logger
}

// Get Retrieve data from the session.
func (m *Manager) Get(key string) string {
	value, ok := m.data[key]
	if ok {
		return value
	}

	return ""
}

// ID Of the session as an HTTP cookie with secure and http-only (cannot be read by JavaScript) enabled.
func (m *Manager) ID() *http.Cookie {
	return &http.Cookie{
		Expires:  m.expires,
		Name:     IDKey,
		Path:     m.cookiePath,
		Secure:   true,
		HttpOnly: true,
		Value:    m.id,
	}
}

// Init Restore a session by ID as a string.
func (m *Manager) Init(id string) error {
	if id == "" {
		return fmt.Errorf(stderr.EmptySessionID)
	}

	// Load from storage.
	data, e1 := m.storage.Load(id)
	if e1 != nil {
		return e1
	}

	m.data = data.Data
	m.id = data.Id
	m.expires = data.Expiration.Add(ExtendTime)

	return nil
}

// RestoreFromCookie Restore a session by ID from an HTTP cookie.
func (m *Manager) RestoreFromCookie(sidCookie *http.Cookie, res http.ResponseWriter) error {
	if sidCookie == nil || sidCookie.Value == "" {
		return fmt.Errorf(stderr.EmptySessionID)
	}

	if !sidCookie.Expires.IsZero() && time.Now().UTC().After(sidCookie.Expires.UTC()) {
		return fmt.Errorf(stderr.ExpiredCookie, sidCookie.Expires.UTC())
	}

	// Load the session from storage.
	data, e1 := m.storage.Load(sidCookie.Value)
	if e1 != nil {
		return e1
	}

	m.data = data.Data
	m.id = data.Id
	m.expires = data.Expiration.Add(ExtendTime)

	return nil
}

// NewManager Initialize a new session manager to handle session save, restore, get, and set.
func NewManager(storage Storage, expiration time.Duration, logger log.Logger) *Manager {
	return &Manager{
		data:       make(Store, 100),
		expires:    time.Now().Add(expiration),
		id:         GenerateID(),
		storage:    storage,
		hasUpdates: false,
		log:        logger,
	}
}

// Save Writes session data to its storage. This is no-op if Set was not previously called.
// Logs a no-op message at level info.
func (m *Manager) Save() error {
	if m.hasUpdates {
		return m.storage.Save(m.id, m.data, m.expires)

	}

	if m.log != nil {
		m.log.Infof("no session updates to save")
	}

	return nil
}

// Remove data from a session
func (m *Manager) Remove(key string) error {
	// verify the key exists
	_, ok := m.data[key]
	if !ok {
		return fmt.Errorf(stderr.NoSuchKey, key)
	}

	// Indicate the session data needs to be saved.
	m.hasUpdates = true

	// Remove the key
	delete(m.data, key)

	return nil
}

// Set Store data in the session.
func (m *Manager) Set(key, value string) {
	m.hasUpdates = true
	m.data[key] = value
}
