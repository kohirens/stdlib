// Package session works to extend HTTP State Management Mechanism beyond the
// HTTP cookie header storage. Such as using files on servers that have file
// storage access or database storage for those without. The latter options
// can add more latency, so please consider your options and use or build an
// implementation according to your use case. For clarity on the subject of
// HTTP State Management please review the RFC at
// https://datatracker.ietf.org/doc/html/rfc6265
package session

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// Handler Handle adding, getting, and removing data from a session.
type Handler interface {
	Get(key string) string
	Remove(key string)
	Set(key, value string)
}

// Storage An interface medium for storing the session data to anyplace an
// implementer see fit. An implementor should especially take into consideration
// sensitive data pertaining to the clients session. This simple interface does
// implementation of encryption for Save and decryption for Load. Use this
// to implement storage for mediums like File, Database, In-memory cache, etc.
type Storage interface {
	// Load The session from storage.
	// No matter the storage medium this should always return JSON as a byte array.
	Load(id string) (*Data, error)

	// Save The session data to the storage medium.
	Save(data *Data) error
}

type Data struct {
	Id         string    `json,bson:"session_id"`
	Expiration time.Time `json,bson:"expiration"`
	Items      Store     `json,bson:"session_data"`
}

// Store Model for short term storage in memory (not intended for long
// term storage).
type Store map[string][]byte

const (
	IDKey = "_sid_"
)

var (
	// ExtendTime How much time the session is extended when a user loads a
	// page after the initial start of the session
	ExtendTime = 5 * time.Minute
)

// GenerateID A unique session ID
func GenerateID() string {
	return uuid.NewString()
}

// Init Will return previous session or a new session if retrieval fails.
// When there is an error, then a new session will be returned.
func Init(c *http.Cookie, s Storage, d time.Duration) (*Manager, error) {
	sm := NewManager(s, d)

	if c != nil { // Look for previous session ID in an HTTP cookie.
		var expErr *ExpiredError
		if e := sm.RestoreFromCookie(c); !errors.As(e, &expErr) {
			return sm, e
		}
	}

	return sm, nil
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
