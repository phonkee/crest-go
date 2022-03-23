package create

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"crest"

	"github.com/phonkee/go-response/v2"
)

// Handler converts function to http handler
// This function is generic over 2 parameters
// 	* S - serializer (request body json)
//	* O - object returned
func Handler[S any, O any](
	fn func(r *http.Request, serializer S) (object O, err error),
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// unmarshal body into serializer
		into := new(S)
		if err := json.NewDecoder(r.Body).Decode(into); err != nil {
			response.Error(fmt.Errorf("%w: %v", crest.ErrInvalidJsonBody, err.Error())).Write(r, w)
			return
		}

		result, err := fn(r, *into)

		if err != nil {
			response.Error(err).Write(r, w)
			return
		}

		// check for resp now
		switch responder := any(result).(type) {
		case crest.Responder:
			w.WriteHeader(responder.StatusCode())
			_ = responder.Write(w)
		default:
			re := response.New(http.StatusCreated)
			if reflect.ValueOf(result).IsNil() {
				re = re.Result(result)
				return
			}
			re.Write(r, w)
			return
		}
	})
}
