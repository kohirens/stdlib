package io

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// NewReadCloser Return an io.ReadCloser from a string or content from a file.
func NewReadCloser(content string, isFile bool) (io.ReadCloser, error) {
	if isFile {
		b, e := os.ReadFile(content)

		if e != nil {
			return nil, fmt.Errorf("could not read %s: %v", content, e.Error())
		}

		return io.NopCloser(bytes.NewBuffer(b)), nil
	}

	return io.NopCloser(bytes.NewBuffer([]byte(content))), nil
}
