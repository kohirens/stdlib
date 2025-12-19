package fsio

import (
	"encoding/json"
	"os"
)

// LoadByJson Initialize a Go struct from a JSON file.
func LoadByJson(filename string, v any) error {
	content, e1 := os.ReadFile(filename)
	if e1 != nil {
		return e1
	}

	if e := json.Unmarshal(content, &v); e != nil {
		return e
	}

	return nil
}
