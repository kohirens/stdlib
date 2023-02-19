package stdlib

import (
	"io"
	"os"
	"strings"
)

// CopyToDir Copy a file to a directory.
func CopyToDir(sourcePath, destDir, separator string) (int64, error) {
	sFile, err1 := os.Open(sourcePath)
	if err1 != nil {
		return 0, err1
	}

	fileStats, err2 := os.Stat(sourcePath)
	if err2 != nil {
		return 0, err2
	}

	dstFile := destDir + separator + fileStats.Name()
	dFile, err3 := os.Create(dstFile)
	if err3 != nil {
		return 0, err3
	}

	return io.Copy(dFile, sFile)
}

// NormalizePath Normalize the path separator.
func NormalizePath(p string) string {
	str := strings.ReplaceAll(p, "/", PS)
	return strings.ReplaceAll(str, "\\", PS)
}
