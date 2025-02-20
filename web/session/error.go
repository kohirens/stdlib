package session

import (
	"errors"
	"fmt"
	"time"
)

var StorageError = errors.New(stderr.NoStorage)

type ExpiredError struct {
	exp time.Time
}

func (e ExpiredError) Error() string {
	return fmt.Sprintf(stderr.ExpiredCookie, e.exp.UTC().Format(time.RFC3339))
}
