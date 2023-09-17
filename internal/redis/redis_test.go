package redis

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	redisAddr     string
	redisPassword string
	realTest      bool
)

func init() {
	_ = godotenv.Load("../configs/.env")

	redisAddr = os.Getenv("REDIS_ADDR")
	redisPassword = os.Getenv("REDIS_PASSWORD")

	if redisAddr != "" && redisPassword != "" {
		realTest = true
	}
}
