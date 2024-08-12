package CustomErrors

import (
	"errors"
)

var (
	ErrNotFound                 = errors.New("not found")
	ErrInternalServerError      = errors.New("internal server error")
	ErrUnexpectedSigningMehthod = errors.New("unexpected signing method")
	ErrUnauthorized             = errors.New("uauthorized")
	ErrBadRequest               = errors.New("bad request")
)
