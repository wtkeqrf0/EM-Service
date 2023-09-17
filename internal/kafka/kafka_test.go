package kafka

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	kafkaAddr string
	realTest  bool
)

func init() {
	_ = godotenv.Load("../configs/.env")

	kafkaAddr = os.Getenv("KAFKA_ADDR")

	if kafkaAddr != "" {
		realTest = true
	}
}
