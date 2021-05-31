package stdlib

import (
	"os"
	"path"
)

const (
	PS = string(os.PathSeparator)
)

var textFileTypes = [4]string{
	".json",
	".md",
	".txt",
	".xml",
}

// DirExist Check if a string path exist.
func DirExist(path string) bool {
	fileObj, err := os.Stat(path)

	if os.IsNotExist(err) || !fileObj.IsDir() {
		return false
	}

	return true
}

// IsTextFile Returns true for files that match the text extensions.
func IsTextFile(file string) (ret bool) {
	ret = false

	ext := path.Ext(file)

	if ext != "" {
	txtcompare: // sorry, I just wanted to play with this so I get used to it. Even though this is single loop or I could just use return. I like to be explicit.
		for _, t := range textFileTypes {
			if t == ext {
				ret = true
				break txtcompare
			}
		}
	}

	return
}

// PathExist true if it is a directory/file and false otherwise.
func PathExist(filename string) bool {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return true
}
