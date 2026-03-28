package errors

import "errors"

var (
	ErrDogNotFound = errors.New("dog not found")
	ErrHumanNotFound = errors.New("human not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrInternalServer = errors.New("internal server error")
	ErrDatabase = errors.New("database error")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden = errors.New("forbidden")
	ErrConflict = errors.New("conflict")
	ErrBadRequest = errors.New("bad request")
)