package CustomErrors

import (
	"errors"
)

var (
	ErrNotFound = errors.New("Not Found")
)