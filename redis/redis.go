package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// Redis struct provides the ability to
// interact with the cache database.
//
// controller.Redis interface and mock.MockRedis
// are generated and based on Redis implementation.
//
//go:generate ifacemaker -f redis.go -o controller/redis.go -i Redis -s Redis -p controller -y "Controller describes methods, implemented by the redis package."
//go:generate mockgen -package mock -source controller/redis.go -destination controller/mock/mock_redis.go
type Redis struct {
	cl *redis.Client
}

// New creates new Redis client.
func New(addr, password string) *Redis {
	return &Redis{cl: redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DialTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 5,
	})}
}

// Save `ret` by key. Key must be cached by func CacheKey.
func (r *Redis) Save(ctx context.Context, key string, ret any) error {
	jsonRet, err := json.Marshal(ret)
	if err != nil {
		return err
	}

	return r.cl.Set(ctx, key, jsonRet, time.Second*15).Err()
}

// Get data of type `want` by `key`.
//
// Key must be cached by func CacheKey. `want` must be a pointer.
func (r *Redis) Get(ctx context.Context, key string, want any) error {
	err := r.cl.Get(ctx, key).Scan(anyUnmarshaler{val: want})
	if err == redis.Nil {
		return nil
	}
	return err
}

type anyUnmarshaler struct {
	val any
}

func (a anyUnmarshaler) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a.val)
}

func (*Redis) CacheKey(obj any) (string, error) {
	jsonGet, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", jsonGet), nil
}

func (r *Redis) Close() error {
	return r.cl.Close()
}
