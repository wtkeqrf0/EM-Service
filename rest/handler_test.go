package rest

import (
	"context"
	"net/http"
	"time"
)

type ValidationError struct {
	Message string `json:"error"`
}

func (ve ValidationError) Error() string {
	return "request validation failed"
}

type GreetingRequest struct {
	Name string `json:"name" schema:"name"`
}

func (r GreetingRequest) Validate() error {
	if r.Name != "" {
		return nil
	}

	return &ValidationError{
		Message: "field 'name' must be specified",
	}
}

type GreetingResponse struct {
	Message string `json:"greeting"`
}

type GreetingError struct {
	Message string `json:"error"`
}

func (e *GreetingError) Error() string {
	return "greeting call failed"
}

type Server struct{}

func (s *Server) Greeting(_ context.Context, r GreetingRequest) (GreetingResponse, error) {
	if r.Name == "Luri" {
		return GreetingResponse{}, &GreetingError{Message: "user Luri already registered"}
	}

	return GreetingResponse{
		Message: "Hello, " + r.Name,
	}, nil
}

func ExampleHandlerFunc() {
	var s Server

	http.Handle("/", API[GreetingRequest, GreetingResponse](s.Greeting))
	go func() {
		_ = http.ListenAndServe("localhost:8080", nil)
	}()

	<-time.NewTimer(time.Minute).C
}
