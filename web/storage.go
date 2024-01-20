package web

import (
	"fmt"
	"github.com/kohirens/stdlib/log"
	"github.com/kohirens/stdlib/path"
	"os"
)

var Log log.Logger = &log.StdLogger{}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

type Storage interface {
	Load(pagePath string) ([]byte, error)
	Save(b []byte, key string) error
}

type LocalStorage struct {
	WorkDir string
}

func (ls *LocalStorage) Save(content []byte, filename string) (int, error) {
	fh, e1 := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0774)
	if e1 != nil {
		return 0, fmt.Errorf(Stderr.CannotOpenFile, filename, e1.Error())
	}

	return fh.Write(content)
}

// Load a file from local storage.
func (ls *LocalStorage) Load(pagePath string) ([]byte, error) {
	Log.Infof(Stdout.LoadPage, pagePath)

	if !path.Exist(pagePath) {
		return nil, fmt.Errorf("file %v not found", pagePath)
	}

	contents, e1 := os.ReadFile(pagePath)
	if e1 != nil {
		return nil, fmt.Errorf(Stderr.CannotReadFile, pagePath, e1.Error())
	}

	Log.Dbugf(Stdout.BytesRead, pagePath, len(contents))

	return contents, nil
}
