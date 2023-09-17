package enricher

import (
	"github.com/joho/godotenv"
	"os"
)

var (
	ageBaseURL     string
	genderBaseURL  string
	countryBaseURL string
	realTest       bool
)

func init() {
	_ = godotenv.Load("../configs/.env")

	ageBaseURL = os.Getenv("AGE_SERVICE_URL")
	genderBaseURL = os.Getenv("GENDER_SERVICE_URL")
	countryBaseURL = os.Getenv("COUNTRY_SERVICE_URL")

	if ageBaseURL != "" && genderBaseURL != "" && countryBaseURL != "" {
		realTest = true
	}
}
