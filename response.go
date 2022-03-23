package crest

import "net/http"

type Responder interface {
	Write(w http.ResponseWriter) error
	StatusCode() int
	Header() http.Header
}

// Response returns full response that can be used
func Response[T any](status int, object T) Responder {
	return &resp[T]{
		status: status,
		object: object,
		header: make(http.Header),
	}
}

type resp[T any] struct {
	status int
	object T
	header http.Header
}

func (o *resp[T]) Header() http.Header {
	return o.header
}

func (o *resp[T]) Write(w http.ResponseWriter) error {
	return nil
}

func (o *resp[T]) StatusCode() int {
	return o.status
}
