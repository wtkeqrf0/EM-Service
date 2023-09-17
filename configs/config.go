package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// Config struct contains all the program configuration data.
type Config struct {
	// `false` value enables `debug` mode.
	Prod bool   `env:"PROD" env-default:"true"`
	Port uint16 `env:"PORT"`

	ServiceURLs struct {
		Age     string `env:"AGE_SERVICE_URL"`
		Gender  string `env:"GENDER_SERVICE_URL"`
		Country string `env:"COUNTRY_SERVICE_URL"`
	}

	Connections struct {
		KafkaAddr   string `env:"KAFKA_ADDR"`
		PostgresURL string `env:"POSTGRES_URL"`
		Redis       struct {
			Addr     string `env:"REDIS_ADDR"`
			Password string `env:"REDIS_PASSWORD"`
		}
	}
}

var cfg *Config

// Build builds the Config struct by environment variables and returns it.
func Build() *Config {
	if cfg == nil {
		if err := godotenv.Load("configs/.env"); err != nil {
			log.WithError(err).Warn()
		}

		cfg = new(Config)

		if err := cleanenv.ReadEnv(cfg); err != nil {
			log.WithError(err).Error("error occurred while reading config file")
			help, _ := cleanenv.GetDescription(cfg, nil)
			log.Fatal(help)
		}

		if !cfg.Prod {
			log.SetLevel(log.DebugLevel)
		}

		log.Debugf("%+v", *cfg)
	}

	return cfg
}
