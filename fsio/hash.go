package fsio

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// Sha256 returns SHA 256 hash for a file.
func Sha256(filePath string) ([]byte, error) {
	file, e1 := os.Open(filePath)
	if e1 != nil {
		return nil, fmt.Errorf(stderr.OpenFile, e1)
	}

	defer func() {
		file.Close()
	}()

	hash := sha256.New()
	if _, e2 := io.Copy(hash, file); e2 != nil {
		return nil, fmt.Errorf(stderr.HashFile, e2)
	}

	return hash.Sum(nil), nil
}
