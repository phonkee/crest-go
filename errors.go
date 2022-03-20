package crest

import (
	"errors"
	"net/http"

	"github.com/phonkee/go-response/v2"
)

var (
	ErrInvalidJsonBody = errors.New("invalid json body")
)

func init() {
	response.RegisterError(ErrInvalidJsonBody, http.StatusBadRequest)
}
