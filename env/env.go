package env

import "os"

// Get An environment variable falling back to a default value when not set.
func Get(key, def string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return v
}
