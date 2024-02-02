package session

import (
	"testing"
	"time"
)

func TestManager(t *testing.T) {
	tests := []struct {
		name       string
		storage    Storage
		expiration time.Duration
	}{
		{"new", &MockStorage{}, 5 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mngr := NewManager(tt.storage, tt.expiration)

			// use default with no error
			if e := mngr.Init(""); e != nil {
				t.Errorf("Manager.Init() = %v, want %v", e.Error(), nil)
			}

			// can load a session from storage
			_ = mngr.Init("abcdefg")
			if got := mngr.Get("test2"); got != "54321" {
				t.Errorf("Manager.Init() = %v, want %v", got, "54321")
			}

			// can set and get an item from the session
			mngr.Set("test", "1245")
			if got := mngr.Get("test"); got != "1245" {
				t.Errorf("Manager.Get() = %v, want %v", got, "1245")
			}

			// can remove an item from the session
			mngr.Remove("test")
			if got := mngr.Get("test"); got != "" {
				t.Errorf("Manager.Get() = %v, want %v", got, "")
			}
		})
	}
}

type MockStorage struct {
	Storage
}

func (ms *MockStorage) Load(id string) (*OfflineStore, error) {
	if id == "abcdefg" {
		return &OfflineStore{Data: Store{"test2": "54321"}}, nil
	}
	return nil, nil
}
func (ms *MockStorage) Save(id string, store Store, expiration time.Time) error {
	return nil
}
