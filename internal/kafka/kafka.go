package kafka

import (
	"context"
	"encoding/json"
	kfk "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"github.com/wtkeqrf0/restService/internal/enricher"
	ce "github.com/wtkeqrf0/restService/internal/enricher/controller"
	"github.com/wtkeqrf0/restService/internal/postgres"
	cp "github.com/wtkeqrf0/restService/internal/postgres/controller"
)

// Kafka struct provides the ability to interact with the Apache Kafka service.
//
// Controller.Kafka interface and mock.MockKafka
// are generated and based on Kafka implementation.
//
//go:generate ifacemaker -f kafka.go -o controller/kafka.go -i Kafka -s Kafka -p controller -y "Controller describes methods, implemented by the kafka package."
//go:generate mockgen -package mock -source controller/kafka.go -destination controller/mock/mock_kafka.go
type Kafka struct {
	fioW       *kfk.Writer
	fioFailedW *kfk.Writer
}

// New checks the connection and initializes topics.
func New(addr string) (*Kafka, error) {
	conn, err := kfk.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if err = conn.CreateTopics(kfk.TopicConfig{
		Topic:             topicFIO,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}, kfk.TopicConfig{
		Topic:             topicFIOFailed,
		NumPartitions:     1,
		ReplicationFactor: 1,
	}); err != nil {
		return nil, err
	}

	return &Kafka{
		fioW: &kfk.Writer{
			Addr:     conn.RemoteAddr(),
			Topic:    topicFIO,
			Balancer: &kfk.LeastBytes{},
		},
		fioFailedW: &kfk.Writer{
			Addr:     conn.RemoteAddr(),
			Topic:    topicFIOFailed,
			Balancer: &kfk.LeastBytes{},
		},
	}, nil
}

const (
	// TopicFIO represents FIO topic.
	topicFIO = "FIO"
	// TopicFIOFailed represents FIO_FAILED topic.
	topicFIOFailed = "FIO_FAILED"
)

type FIO struct {
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic,omitempty"`
}

// Produce creates a new kafka message.
// When field `cause` is `nil`, message will be `success` status.
func (k *Kafka) Produce(ctx context.Context, fio FIO, causes map[string]string) (err error) {
	msg := kfk.Message{}

	msg.Key, err = json.Marshal(fio)
	if err != nil {
		return err
	}

	if causes == nil {
		return k.fioW.WriteMessages(ctx, msg)
	}

	msg.Value, err = json.Marshal(causes)
	if err != nil {
		return err
	}
	return k.fioFailedW.WriteMessages(ctx, msg)
}

// Consume starts to consume messages from topics `FIO` and `FIO_FAILED`.
//
// Method doesn't block the main goroutine.
func (k *Kafka) Consume(ctx context.Context, db cp.Postgres, enr ce.Enricher) {
	go consumeFIO(ctx, k.fioW.Addr.String(), db, enr)
	go consumeFIOFailed(ctx, k.fioFailedW.Addr.String())
}

// consumeFIO starts receiving FIO and processing them by `success` function.
func consumeFIO(ctx context.Context, addr string, db cp.Postgres, enr ce.Enricher) {
	r := kfk.NewReader(kfk.ReaderConfig{
		Brokers:  []string{addr},
		GroupID:  "success",
		Topic:    topicFIO,
		MaxBytes: 10e6,
		MinBytes: 10e3,
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.WithError(err).Warn("can't read message")
			continue
		}

		var fio FIO
		if err = json.Unmarshal(msg.Key, &fio); err != nil {
			log.WithError(err).Warn("can't unmarshal to FIO")
			continue
		}

		log.Debugf("received from %v: %+v", msg.Topic, fio)

		enrichedFIO, err := enr.EnrichFIO(ctx, enricher.FIO(fio))
		if err != nil {
			log.WithError(err).Warnf("can't enrich FIO: %+v", fio)
			continue
		}

		if err = db.SaveUser(ctx, postgres.EnrichedFIOWithCreationTime{
			EnrichedFIO:  postgres.EnrichedFIO(enrichedFIO),
			CreationTime: msg.Time,
		}); err != nil {
			log.WithError(err).Warnf("failed to save fio: %+v", enrichedFIO)
			continue
		}

		if err = r.CommitMessages(ctx, msg); err != nil {
			log.WithError(err).Warnf("failed to commit message from %s: %+v", msg.Topic, msg)
		}
	}
}

// consumeFIOFailed starts receiving FIO and processing them by `failure` function.
func consumeFIOFailed(ctx context.Context, addr string) {
	r := kfk.NewReader(kfk.ReaderConfig{
		Brokers:  []string{addr},
		GroupID:  "failure",
		Topic:    topicFIOFailed,
		MaxBytes: 10e6,
		MinBytes: 10e3,
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.WithError(err).Warn("can't read message")
			continue
		}

		var fio FIO
		if err = json.Unmarshal(msg.Key, &fio); err != nil {
			log.WithError(err).Warnf("can't unmarshal to type FIO")
			continue
		}

		var causes map[string]string
		if err = json.Unmarshal(msg.Value, &causes); err != nil {
			log.WithError(err).Warnf("can't unmarshal msg.Value to type map[string]string")
			continue
		}

		log.Debugf("received from %s: %+v for a reason(s): %s", msg.Topic, fio, causes)

		if err = r.CommitMessages(ctx, msg); err != nil {
			log.WithError(err).Warnf("failed to commit message from %s: %v", msg.Topic, err)
		}
	}
}
