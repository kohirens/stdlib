package github

import (
	"bytes"
	"fmt"
	"os"
)

func bodyFromFile(filepath string) (*bytes.Reader, error) {
	body, errBody := os.ReadFile(filepath)
	if errBody != nil {
		return nil, fmt.Errorf(stderr.CouldNotReadFile, filepath, errBody.Error())
	}

	return bytes.NewReader(body), nil
}
