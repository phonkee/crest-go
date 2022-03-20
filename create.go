package crest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/phonkee/go-response/v2"
)

// Create converts function to http handler function
func Create[S any, O any](fn func(ctx context.Context, serializer S) (object O, err error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// unmarshal body into serializer
		into := new(S)
		if err := json.NewDecoder(r.Body).Decode(&into); err != nil {
			response.Error(fmt.Errorf("%w: %v", ErrInvalidJsonBody, err.Error())).Write(r, w)
			return
		}

	}
}
