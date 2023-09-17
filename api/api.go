package api

import (
	ce "github.com/wtkeqrf0/restService/enricher/controller"
	ck "github.com/wtkeqrf0/restService/kafka/controller"
	cp "github.com/wtkeqrf0/restService/postgres/controller"
	cr "github.com/wtkeqrf0/restService/redis/controller"
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
