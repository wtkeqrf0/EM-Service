package api

import (
	ce "github.com/wtkeqrf0/restService/internal/enricher/controller"
	ck "github.com/wtkeqrf0/restService/internal/kafka/controller"
	cp "github.com/wtkeqrf0/restService/internal/postgres/controller"
	cr "github.com/wtkeqrf0/restService/internal/redis/controller"
)

type Controllers struct {
	ce.Enricher
	ck.Kafka
	cp.Postgres
	cr.Redis
}

type Server struct {
	ctrl Controllers
}

func NewServer(ctrl Controllers) *Server {
	newValidator(ctrl.Enricher)
	return &Server{
		ctrl: ctrl,
	}
}
