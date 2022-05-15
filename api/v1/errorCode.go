package v1

import (
	goplaylist "mini-clean/error"
	"net/http"
)

func GetErrorStatus(err error) int {
	switch err {
	case goplaylist.ErrBadRequest:
		return http.StatusBadRequest
	case goplaylist.ErrInternalServer:
		return http.StatusInternalServerError
	case goplaylist.ErrNotFound:
		return http.StatusNotFound
	case goplaylist.ErrUnauthorized:
		return http.StatusUnauthorized
	}
	return http.StatusOK
}
