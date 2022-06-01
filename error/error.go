package error

import "errors"

var (
	ErrNotFound       = errors.New("record not found")
	ErrInternalServer = errors.New("internal server error")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrBadRequest     = errors.New("bad request")
)
