package path

import (
	"io"
	"os"
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

// DirExist Those rare times when you just need to check if a string path is a
// directory and not a file.
func DirExist(path string) bool {
	fileObj, err := os.Stat(path)

	if os.IsNotExist(err) || !fileObj.IsDir() {
		return false
	}

	return true
}