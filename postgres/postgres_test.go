package postgres

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	postgresURL string
	realTest    bool
)

func init() {
	_ = godotenv.Load("../configs/.env")

	postgresURL = os.Getenv("POSTGRES_URL")

	if postgresURL != "" {
		realTest = true
	}
}
