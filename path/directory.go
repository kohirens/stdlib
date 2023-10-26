package path

import (
	"fmt"
	"github.com/kohirens/stdlib/log"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// CopyDirToDir Copy files from 1 directory into another. The files of src are
// copied and not the directory itself to the root of the dst directory. Any
// files that exist in the destination directory are overwritten.
func CopyDirToDir(src, dst, ps string, mode os.FileMode) error {
	if !Exist(src) {
		return fmt.Errorf("directory %v does not exist", src)
	}

	// Recursively walk the directory.
	return filepath.Walk(src, func(sourceFile string, fi os.FileInfo, wErr error) error {
		if wErr != nil {
			return wErr
		}

		if sourceFile == src {
			return nil
		}

		log.Infof("file to copy: %v", sourceFile)

		copyTo := dst + strings.Replace(sourceFile, src, "", 1)
		log.Infof("copy to: %v", copyTo)

		// Skip directories.
		if fi.IsDir() {
			if e := os.MkdirAll(copyTo, mode); e != nil { // Make the directory before exiting.
				return fmt.Errorf("could not make directory %v because %v", copyTo, e.Error())
			}
			return nil
		}

		// Make the parent directory of the file.
		dir := filepath.Dir(copyTo)
		log.Infof("making directory %v", dir)
		if e := os.MkdirAll(dir, mode); e != nil {
			return fmt.Errorf("could not make directory %v because %v", dir, e.Error())
		}

		log.Logf("copy file %v to %v", sourceFile, copyTo)
		_, e1 := CopyToDir(sourceFile, dir, ps)

		return e1
	})
}

// CopyToDir Copy a file to a directory.
//
//	This uses os.Create in a way that will cause the file to be overwritten if
//	it exists.
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
