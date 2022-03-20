package crest

import (
	"net/http"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Run("test create", func(t *testing.T) {
		type Serializer struct {
			Name  string
			Value int
		}
		type Object struct {
			ID    string
			Name  string
			Value int
		}

		// test create function
		Create(func(r *http.Request, serializer Serializer) (object Object, err error) {
			return Object{}, nil
		})

	})
}
