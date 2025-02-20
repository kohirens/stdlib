package session

import (
	"fmt"
	"time"
)

type ExpiredError struct {
	exp time.Time
}

func (e ExpiredError) Error() string {
	return fmt.Sprintf(stderr.ExpiredCookie, e.exp.UTC().Format(time.RFC3339))
}
