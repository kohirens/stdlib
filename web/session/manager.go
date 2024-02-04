package session

import (
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

func (m *Manager) Init(id string) error {
	if id == "" { // continue with the default values.
		return nil
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
