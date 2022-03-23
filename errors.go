package crest

import (
	"errors"
	"net/http"

	"github.com/phonkee/go-response/v2"
)

// HttpError custom interface that gives you ability to provide status code. Otherwise, Internal server error will be
// returned.
type HttpError interface {
	error
	// StatusCode returns http status code for given error
	StatusCode() int
	// Unwrap adds support for unwrapping errors (errors.Is)
	Unwrap() error
}

var (
	ErrInvalidJsonBody = errors.New("invalid json body")
)

func init() {
	response.RegisterError(ErrInvalidJsonBody, http.StatusBadRequest)
}

// Error wraps error and returns error with http status code
func Error(source error, statusCode int) HttpError {
	return &err{
		source:     source,
		statusCode: statusCode,
	}
}

type err struct {
	source     error
	statusCode int
}

func (e *err) Error() string {
	return e.source.Error()
}

func (e *err) StatusCode() int {
	return e.statusCode
}

func (e *err) Unwrap() error {
	return e.source
}
