package io

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func ReadCloser(filepath string) (io.ReadCloser, error) {
	b, e := os.ReadFile(filepath)

	if e != nil {
		return nil, fmt.Errorf("could not read %s: %v", filepath, e.Error())
	}

	return io.NopCloser(bytes.NewBuffer(b)), nil
}
