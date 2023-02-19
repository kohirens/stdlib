package stdlib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	PS = string(os.PathSeparator)
)

// FileExtChecker A store for file extensions to exclude and included.
type FileExtChecker struct {
	excludes *[]string
	includes *[]string
}

// NewFileExtChecker Initialize a new FileExtChecker instance.
func NewFileExtChecker(el, in *[]string) (*FileExtChecker, error) {
	var err error

	if el == nil && in == nil {
		err := fmt.Errorf("you did not add provide any file extensions to include or exclude")
		return nil, err
	}

	if el == nil {
		el = &[]string{}
	}

	if in == nil {
		in = &[]string{}
	}

	fce := FileExtChecker{
		excludes: el,
		includes: in,
	}

	return &fce, err
}

func InitFileExtChecker() *FileExtChecker {
	return &FileExtChecker{
		excludes: &[]string{},
		includes: &[]string{},
	}
}

func NewFileExtCheckerStr(el, in []string) (*FileExtChecker, error) {
	return NewFileExtChecker(&el, &in)
}

// DirExist Check if a string path exist.
func DirExist(path string) bool {
	fileObj, err := os.Stat(path)

	if os.IsNotExist(err) || !fileObj.IsDir() {
		return false
	}

	return true
}

// IsValid Returns true for files that match allowed or excluded extensions.
// Passing a full path only checks the basename.
// Default to include all files.
// If a file has no extension, then use its basename.
func (fec *FileExtChecker) IsValid(file string) bool {
	basename := filepath.Base(file) // account for hidden directories
	ext := strings.Trim(filepath.Ext(basename), ".")
	// when there is no extension (unix/linux/mac) use the basename
	if len(ext) == 0 && len(basename) > 1 {
		ext = basename
	}

	if ext != "" {
		for _, t := range *fec.excludes {
			if t == ext {
				return false
			}
		}
		for _, t := range *fec.includes {
			if t == ext {
				return true
			}
		}
	}

	return true
}

// PathExist true for a directory/file and false otherwise.
func PathExist(filename string) bool {
	_, err := os.Stat(filename)

	if err == nil {
		return true
	}

	return false
}
