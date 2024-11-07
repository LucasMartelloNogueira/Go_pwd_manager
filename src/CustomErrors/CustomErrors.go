package CustomErrors

import (
	"errors"
)

var (
	ErrNotFound                 = errors.New("not found")
	ErrInternalServerError      = errors.New("internal server error")
	ErrUnexpectedSigningMehthod = errors.New("unexpected signing method")
	ErrUnauthorized             = errors.New("unauthorized")
	ErrBadRequest               = errors.New("bad request")
)
