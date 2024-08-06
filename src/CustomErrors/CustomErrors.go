package customErrors

import (
	"errors"
)

var (
	ErrNotFound = errors.New("Not Found")
)