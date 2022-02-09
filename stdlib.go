package stdlib

import (
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	PS = string(os.PathSeparator)
)

type FileExtChecker struct {
	excludes *[]string
	includes *[]string
}

// New Add file extensions for the file checker to exclude as being included
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

// DirExist Check if a string path exist.
func DirExist(path string) bool {
	fileObj, err := os.Stat(path)

	if os.IsNotExist(err) || !fileObj.IsDir() {
		return false
	}

	return true
}

// IsValid Returns true for files that match allowed extensions.
func (fec *FileExtChecker) IsValid(file string) (ret bool) {
	ret = false

	ext := strings.Trim(path.Ext(file), ".")

	if ext != "" {
		for _, t := range *fec.excludes {
			ret = true
			if t == ext {
				return false
			}
		}
		for _, t := range *fec.includes {
			if t == ext {
				ret = true
				break
			}
		}
	}

	return
}

// PathExist true for a directory/file and false otherwise.
func PathExist(filename string) bool {
	_, err := os.Stat(filename)

	if err == nil {
		return true
	}

	return false
}
