package fsio

import (
	"os"
	"strings"
)

const (
	PS = string(os.PathSeparator)
)

// Normalize Switch the path separator to same as the os.PathSeparator
func Normalize(p string) string {
	str := strings.ReplaceAll(p, "/", PS)
	return strings.ReplaceAll(str, "\\", PS)
}

// Exist true for a directory/file and false otherwise.
func Exist(filename string) bool {
	_, e1 := os.Stat(filename)

	return e1 == nil || os.IsExist(e1)
}
