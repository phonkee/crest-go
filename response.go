package crest

import (
	"net/http"
	"reflect"
)

type Responder interface {
	Write(w http.ResponseWriter) error
	StatusCode() int
	Header() http.Header
}

// Response returns full response that can be used in scenarios where you want to control which http status you want
// to return. Also when using this response you can add custom headers
func Response[T any](status int, object T, headers ...string) Responder {
	result := &resp[T]{
		status: status,
		object: object,
		header: make(http.Header),
	}

	return result
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
	if reflect.ValueOf(o.object).IsNil() {
		return nil
	}
	return nil
}

func (o *resp[T]) StatusCode() int {
	return o.status
}
