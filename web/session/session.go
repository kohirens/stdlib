// Package session works to extend HTTP State Management Mechanism beyond the
// HTTP cookie header storage. Such as using files on servers that have file
// storage access or database storage for those without. The latter options
// can add more latency, so please consider your options and use or build an
// implementation according to your use case. For clarity on the subject of
// HTTP State Management please review the RFC at
// https://datatracker.ietf.org/doc/html/rfc6265
package session

import (
	"github.com/google/uuid"
	"time"
)

const (
	IDKey = "_sid_"
)

// Handler Handle adding, getting, and removing data from a session.
type Handler interface {
	Get(key string) string
	Remove(key string)
	Set(key, value string)
}

// OfflineStore Model for long term storage
type OfflineStore struct {
	Id         string            `bson:"session_id"`
	Expiration time.Time         `bson:"expiration"`
	Data       map[string]string `bson:"session_data"`
}

// Storage an interface to a medium for storing the session data out-side of an
// HTTP cookie header. Especially for sensitive data pertaining to a users
// session. Use this to implement storage for mediums like File, Database,
// In-memory cache, etc.
type Storage interface {
	// Load The session from storage.
	Load(id string) (*OfflineStore, error)

	// Save The session to storage.
	Save(id string, store Store, expiration time.Time) error
}

// Store Model for short term storage in memory (not intended for long
// term storage, see OfflineStore for that purpose).
type Store map[string]string

var (
	// ExtendTime How much time the session is extended when a user loads a
	// page after the initial start of the session
	ExtendTime = 5 * time.Minute
)

// GenerateID A unique session ID
func GenerateID() string {
	return uuid.NewString()
}
