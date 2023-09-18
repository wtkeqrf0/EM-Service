package api

import (
	ce "github.com/wtkeqrf0/restService/internal/enricher/controller"
	ck "github.com/wtkeqrf0/restService/internal/kafka/controller"
	cp "github.com/wtkeqrf0/restService/internal/postgres/controller"
	cr "github.com/wtkeqrf0/restService/internal/redis/controller"
)

// Controllers struct contains interfaces,
// that can realize server capabilities.
type Controllers struct {
	ce.Enricher
	ck.Kafka
	cp.Postgres
	cr.Redis
}

// Server struct represents the server capabilities.
type Server struct {
	ctrl Controllers
}

// NewServer creates a new server with given Controllers.
func NewServer(ctrl Controllers) *Server {
	newValidator(ctrl.Enricher)
	return &Server{
		ctrl: ctrl,
	}
}
