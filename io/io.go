package io

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// ReadCloser Return an io.ReadCloser from a string or filepath.
//
// Deprecated This was quickly added and now is replaced to add a new feature.
// Use NewReadCloser as a replacement since it add the feature to indicate if
// the content is a string to be used in the buffer or a path to a file to load
// into the buffer.
func ReadCloser(filepath string) (io.ReadCloser, error) {
	b, e := os.ReadFile(filepath)

	if e != nil {
		return nil, fmt.Errorf("could not read %s: %v", filepath, e.Error())
	}

	return io.NopCloser(bytes.NewBuffer(b)), nil
}

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
