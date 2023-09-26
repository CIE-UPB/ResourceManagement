package config

import (
	"log"

	"github.com/caarlos0/env/v9"
)

type Env struct {
	DBName         string `env:"DATABASE_NAME" envDefault:"cie-db"`
	DBUsername     string `env:"DATABASE_USER" envDefault:"postgres"`
	DBPassword     string `env:"DATABASE_PASSWORD" envDefault:"postgres"`
	DBHost         string `env:"DATABASE_HOST" envDefault:"postgresql"`
	DBPort         string `env:"DATABASE_PORT" envDefault:"5432"`
	DBSslmode      string `env:"DATABASE_SSLMODE" envDefault:"disable"`
	DBLogmode      string `env:"DATABASE_LOGMODE" envDefault:"false"`
	GinMode        string `env:"GIN_MODE" envDefault:"debug"`
	ServerPort     string `env:"SERVER_PORT" envDefault:"8080"`
	Environment    string `env:"ENVIRONMENT" envDefault:"development"`
	LogOutput      string `env:"LOG_OUTPUT" envDefault:""`
	LogLevel       string `env:"LOG_LEVEL" envDefault:"debug"`
	ContextTimeout int    `env:"CONTEXT_TIMEOUT" envDefault:"5"`
}

var environ Env

func init() {
	if err := env.Parse(&environ); err != nil {
		log.Fatalf("Error al leer las variables de entorno: %v", err)
	}
}

func NewEnv() Env {
	return environ
}
