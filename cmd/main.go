package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/wtkeqrf0/restService/api"
	"github.com/wtkeqrf0/restService/configs"
	"github.com/wtkeqrf0/restService/graph"
	"github.com/wtkeqrf0/restService/internal/enricher"
	"github.com/wtkeqrf0/restService/internal/kafka"
	"github.com/wtkeqrf0/restService/internal/postgres"
	"github.com/wtkeqrf0/restService/internal/redis"
	"github.com/wtkeqrf0/restService/rest"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
}

func main() {
	cfg := configs.Build()

	//-----------------------Initialize controllers-----------------------

	var ctrl api.Controllers

	httpClient := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}, Timeout: time.Second * 5}

	// TODO: postgres migration (without schema creation)
	ctrl.Postgres = postgres.New(cfg.Connections.PostgresURL)
	ctrl.Redis = redis.New(cfg.Connections.Redis.Addr, cfg.Connections.Redis.Password)
	ctrl.Enricher = enricher.New(
		httpClient,
		cfg.ServiceURLs.Age,
		cfg.ServiceURLs.Gender,
		cfg.ServiceURLs.Country,
	)

	var err error
	ctrl.Kafka, err = kafka.New(cfg.Connections.KafkaAddr)
	if err != nil {
		log.WithError(err).Fatal("failed to initialize Kafka service")
	}

	srv := api.NewServer(ctrl)

	//-------------------------Setup handlers-------------------------

	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer, middleware.Timeout(60*time.Second))

	// REST handlers
	r.Route("/fio", func(r chi.Router) {
		r.Post("/", rest.API[api.CreateFIORequest, api.CreateFIOResponse](srv.CreateFIO).HandlerFunc())
		r.Patch("/", rest.API[api.UpdateEnrichedFIORequest, api.UpdateEnrichedFIOResponse](srv.UpdateEnrichedFIO).HandlerFunc())
		r.Get("/", rest.API[api.GetEnrichedFIORequest, api.GetEnrichedFIOResponse](srv.GetEnrichedFIO).HandlerFunc())
		r.Delete("/", rest.API[api.DeleteEnrichedFIORequest, api.DeleteEnrichedFIOResponse](srv.DeleteEnrichedFIO).HandlerFunc())
	})

	// GraphQL handler
	r.Route("/query", func(r chi.Router) {
		r.Handle("/", handler.NewDefaultServer(
			graph.NewExecutableSchema(
				graph.Config{
					Resolvers: &graph.Resolver{
						Server: srv,
					},
				},
			),
		),
		)
	})

	//-------------------Create and start http httpServer-------------------

	httpSrv := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Port),
		Handler:        r,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx := context.Background()
	ctrl.Kafka.Consume(ctx, ctrl.Postgres, ctrl.Enricher)

	go func() {
		if err = httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatalf("error occurred while running http httpSrv")
		}
	}()

	log.Infof("connect to http://localhost:%d/ for GraphQL playground", cfg.Port)

	<-quit

	log.Info("server is shutting down ...")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err = httpSrv.Shutdown(ctx); err != nil {
		log.WithError(err).Error("server shutdown failed")
	}

	if err = ctrl.Postgres.Close(); err != nil {
		log.WithError(err).Error("postgres close failed")
	}

	if err = ctrl.Redis.Close(); err != nil {
		log.WithError(err).Error("redis close failed")
	}

	log.Info("server exited properly")
}
