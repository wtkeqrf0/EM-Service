package redis

import (
	"context"
	"encoding/json"
	"testing"
	"time"
)

type fio struct {
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic,omitempty"`
}

func TestRedis_Get(t *testing.T) {
	if !realTest {
		t.Skip()
	}

	cl := New(redisAddr, redisPassword)
	defer cl.cl.Close()

	data := fio{
		Name:    "Matvey",
		Surname: "Sizov",
	}

	jsonRet, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	key := "2"

	cl.cl.Set(ctx, key, jsonRet, time.Second*5)

	var res fio
	if err = cl.Get(ctx, key, &res); err != nil {
		t.Fatal(err)
	}

	if data != res {
		t.Fatalf("Want: %+v, Got: %+v", data, res)
	}

	t.Logf("%+v", res)
}

func TestRedis_Save(t *testing.T) {
	if !realTest {
		t.Skip()
	}

	cl := New(redisAddr, redisPassword)
	defer cl.cl.Close()

	data := fio{
		Name:    "Matvey",
		Surname: "Sizov",
	}

	ctx := context.Background()
	key := "1"

	if err := cl.Save(ctx, key, data); err != nil {
		t.Fatal(err)
	}

	var res fio
	err := cl.cl.Get(ctx, key).Scan(anyUnmarshaler{val: &res})
	if err != nil {
		t.Fatal(err)
	}

	if data != res {
		t.Fatalf("Want: %+v, Got: %+v", data, res)
	}

	t.Logf("%+v", res)
}
