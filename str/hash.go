package str

import (
	"crypto/sha256"
)

// Sha256 returns SHA 256 hash for a string.
func Sha256(str string) ([]byte, error) {
	// Create a new SHA256 hash
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(str))

	// Get the hash sum as a byte slice
	return h.Sum(nil), nil
}
