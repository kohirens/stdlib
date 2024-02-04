package session

import (
	"fmt"
	"github.com/kohirens/stdlib/log"
	"net/http"
	"time"
)

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

func (m *Manager) Get(key string) string {
	value, ok := m.data[key]
	if ok {
		return value
	}

	return ""
}

func (m *Manager) ID() *http.Cookie {
	return &http.Cookie{Name: IDKey, Value: m.id, Path: m.cookiePath, Secure: true}
}

// Init Loads an already-in-session session from storage by a string. In
// contrast to RestoreFromCookie, which you will likely only use 1 of these
// methods.
func (m *Manager) Init(id string) error {
	if id == "" {
		return fmt.Errorf(stderr.EmptySessionID)
	}

	// otherwise load from wherever storage keeps it.
	data, e1 := m.storage.Load(id)
	if e1 != nil {
		return e1
	}

	m.data = data.Data
	m.id = data.Id
	m.expires = data.Expiration.Add(ExtendTime)

	return nil
}

// RestoreFromCookie Get the session ID from a cookie then load it from storage.
func (m *Manager) RestoreFromCookie(sidCookie *http.Cookie, res http.ResponseWriter) error {
	if sidCookie.Value == "" {
		return fmt.Errorf(stderr.EmptySessionID)
	}

	// otherwise load from wherever storage keeps it.
	data, e1 := m.storage.Load(sidCookie.Value)
	if e1 != nil {
		return e1
	}

	m.data = data.Data
	m.id = data.Id
	m.expires = data.Expiration.Add(ExtendTime)

	return nil
}

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

// Save calls the storage.Save, this is no-up is Set was never called.
// Logs (at level info) when no-up.
func (m *Manager) Save() error {
	if m.hasUpdates {
		return m.storage.Save(m.id, m.data, m.expires)

	}

	if m.log != nil {
		m.log.Infof("no session updates to save")
	}

	return nil
}

func (m *Manager) Remove(key string) {
	delete(m.data, key)
}

func (m *Manager) Set(key, value string) {
	m.hasUpdates = true
	m.data[key] = value
}
