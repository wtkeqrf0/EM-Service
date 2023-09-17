package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
)

// Set a Decoder instance as a package global, because it caches
// meta-data about structs, and an instance can be shared safely.
var decoder = schema.NewDecoder()

// API implements the [http.Handler] interface with a template representation of the request structure and
// response. With this handler, you can wrap business logic functions and not perform
// marshaling and unmarshalling of request structures and URLs. For example:
//
// type Server struct {}
//
// type Request struct {
// Name string `schema:"name"`
// }
//
// type Response struct {
// Greeting string `json:"greeting"`
// }
//
// func (s *Server) Hello(ctx context.Context, r Request) (*Response, error) { ... }
//
//		func main() {
//			var s Server
//			http.Handle("/", API[Request, Response](s.Hello))
//			http.ListenAndServe("localhost:8080", nil)
//	}
//
// The handler returns the status 200 and a JSON response with the fields:
//
// struct {
// Status string `json:"status"`
// Message string `json:"message,omitempty"`
// Payload any `json:"payload"`
// }
//
// where:
// - status - "success" or "error", depending on how the handler call ended.
// - message is err.Error() in the case when the handler call is completed with an error.
// - payload is a JSON type of Response template or a JSON type of error structure.
//
// When a request is received in the Handler, an attempt is made to decode the http request body into a
// typed Request parameter. The error is not checked. Further, using the library
// https://github.com/gorilla/schema additionally, an attempt is made to apply the URL keys of the request
// to the Request. Thus, the Request by priority is first filled in from the request body, then
// from the URL arguments.
//
// Then an attempt is made to "validate" the request. Response is checked for the presence of the method
// Validate() error. If the method is implemented, it is called. If a validation error has occurred,
// the handler fills in the message and payload fields and returns the result with an error.
//
// If the request is successfully validated or there is no implementation of the Validate method, the Request executes
// the handler function. To which the completed Request fields are passed.
//
// Depending on the result of the handler execution, the structure fields and its status are filled in.
// The response is sent as JSON with the status 200.
type API[Request, Response any] func(context.Context, Request) (Response, error)

// HandlerFunc returns [http.HandlerFunc] to simplify routing in some situations.
func (h API[Request, Response]) HandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

// ServeHTTP implements the [http.Handler] interface and calls the handler.
func (h API[Request, Response]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req    Request
		resp   Response
		err    error
		output struct {
			Error   string `json:"error,omitempty"`
			Payload any    `json:"payload"`
		}
		errCode = http.StatusBadRequest
	)

	_ = json.NewDecoder(r.Body).Decode(&req)
	_ = decoder.Decode(&req, r.URL.Query())

	if v, ok := any(&req).(interface{ Validate() error }); ok {
		err = v.Validate()
	}

	if err == nil {
		resp, err = h(r.Context(), req)
		errCode = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		output.Error = err.Error()
		output.Payload = err
		w.WriteHeader(errCode)
	} else {
		output.Payload = resp
	}

	raw, _ := json.Marshal(output)
	_, _ = w.Write(raw)
}
