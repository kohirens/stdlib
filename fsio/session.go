package fsio

import (
	"encoding/json"
	"fmt"
	"github.com/kohirens/stdlib/web/session"
	"os"
	"path/filepath"
)

// NewStorageLocal Initialize local session storage.
func NewStorageLocal(workDir string) *LocalStorage {
	return &LocalStorage{
		WorkDir: workDir,
	}
}

type LocalStorage struct {
	WorkDir string
}

// Load session data from local file storage.
func (ls *LocalStorage) Load(id string) (*session.Data, error) {
	f := filepath.Join(ls.WorkDir, id)
	if !Exist(f) {
		return nil, fmt.Errorf("file %v not found", f)
	}

	content, e1 := os.ReadFile(f)
	if e1 != nil {
		return nil, fmt.Errorf(Stderr.ReadFile, f, e1.Error())
	}

	data := &session.Data{}

	if e := json.Unmarshal(content, data); e != nil {
		return nil, fmt.Errorf(Stderr.DecodeJSON, f, e)
	}

	return data, nil
}

// Save session data to a local file for storage.
func (ls *LocalStorage) Save(data *session.Data) error {
	f := filepath.Join(ls.WorkDir, data.Id)

	fh, e1 := os.OpenFile(f, os.O_CREATE|os.O_RDWR, DefaultFilePerms)
	if e1 != nil {
		return fmt.Errorf(Stderr.OpenFile, f, e1)
	}

	content, e2 := json.Marshal(data)
	if e2 != nil {
		return fmt.Errorf(Stderr.EncodeJSON, e2)
	}

	_, e3 := fh.Write(content)
	if e3 != nil {
		return fmt.Errorf(Stderr.WriteFile, e3)
	}

	return nil
}
