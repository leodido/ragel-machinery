package ragelmachinery

import (
	"errors"
)

var (
	// ErrNotFound represents a situation in which the needle we were looking for has not been found.
	ErrNotFound = errors.New("needle not found")
)
