package kafka

import (
	"context"
	"encoding/json"
	kfk "github.com/segmentio/kafka-go"
	"github.com/wtkeqrf0/restService/enricher"
	me "github.com/wtkeqrf0/restService/enricher/controller/mock"
	mp "github.com/wtkeqrf0/restService/postgres/controller/mock"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestKafka_Consume(t *testing.T) {
	if !realTest {
		t.Skip()
	}

	k, err := New(kafkaAddr)
	if err != nil {
		t.Fatal(err)
	}

	t.Parallel()

	data := FIO{
		Name:    "Matvey1",
		Surname: "Sizov",
	}

	ctrl := gomock.NewController(t)
	ctx := context.Background()
	ch := make(chan FIO, 1)

	enr := me.NewMockEnricher(ctrl)
	enr.EXPECT().EnrichFIO(gomock.Any(), gomock.AssignableToTypeOf(enricher.FIO{})).
		DoAndReturn(func(_ any, fio enricher.FIO) (enricher.EnrichedFIO, error) {
			ch <- FIO(fio)
			return enricher.EnrichedFIO{}, nil
		}).MinTimes(1).MaxTimes(3)

	db := mp.NewMockPostgres(ctrl)
	db.EXPECT().SaveUser(gomock.Any(), gomock.Any()).Return(nil).MinTimes(1).MaxTimes(3)

	jsonFio, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	if err = (&kfk.Writer{
		Addr:     k.fioW.Addr,
		Topic:    topicFIO,
		Balancer: &kfk.LeastBytes{},
	}).WriteMessages(ctx, kfk.Message{
		Key: jsonFio,
	}); err != nil {
		t.Error(err)
	}

	k.Consume(ctx, db, enr)

	select {
	case v := <-ch:
		if v != data {
			t.Fatalf("Want: %+v, Got: %+v", data, v)
		}
		t.Logf("%+v", v)
	case <-time.Tick(time.Second * 35):
		t.Fatal("message not found")
	}
}

func TestKafka_Produce(t *testing.T) {
	if !realTest {
		t.Skip()
	}

	k, err := New(kafkaAddr)
	if err != nil {
		t.Fatal(err)
	}
	t.Parallel()

	data := FIO{
		Name:    "Matvey2",
		Surname: "Sizov",
	}

	ctx := context.Background()

	if err = k.Produce(ctx, data, map[string]string{}); err != nil {
		t.Fatal(err)
	}

	ch := make(chan FIO, 1)

	go func() {
		r := kfk.NewReader(kfk.ReaderConfig{
			Brokers:  []string{kafkaAddr},
			GroupID:  "failed",
			Topic:    topicFIOFailed,
			MaxBytes: 10e6,
			MinBytes: 10e3,
		})

		var msg kfk.Message
		msg, err = r.ReadMessage(ctx)
		if err != nil {
			t.Error(err)
			return
		}

		var fio FIO
		if err = json.Unmarshal(msg.Key, &fio); err != nil {
			t.Error(err)
			return
		}

		ch <- fio

		if err = r.CommitMessages(ctx, msg); err != nil {
			t.Error(err)
		}
	}()

	select {
	case v := <-ch:
		if v != data {
			t.Fatalf("Want: %+v, Got: %+v", data, v)
		}
		t.Logf("%+v", v)
	case <-time.Tick(time.Second * 35):
		t.Fatal("message not found")
	}
}
