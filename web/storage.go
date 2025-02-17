package web

import (
	"fmt"
	"github.com/kohirens/stdlib/fsio"
	"github.com/kohirens/stdlib/logger"
	"os"
)

var log = &logger.StdLogger{}

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
	log.Infof(Stdout.LoadPage, pagePath)

	if !fsio.Exist(pagePath) {
		return nil, fmt.Errorf("file %v not found", pagePath)
	}

	contents, e1 := os.ReadFile(pagePath)
	if e1 != nil {
		return nil, fmt.Errorf(Stderr.CannotReadFile, pagePath, e1.Error())
	}

	log.Dbugf(Stdout.BytesRead, pagePath, len(contents))

	return contents, nil
}
