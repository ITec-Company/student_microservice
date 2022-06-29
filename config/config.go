package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"student_microservice/internal/logging"
	"sync"
)

type Config struct {
	Server struct {
		Type string `env:"LISTEN_TYPE" env-default:"port"`
		Port string `env:"SERVER_PORT" env-default:"8080"`
		Host string `env:"SERVER_HOST" env-default:"127.0.0.1"`
	}
	DB struct {
		Host     string `env:"DB_HOST" env-default:"127.0.0.1"`
		Port     string `env:"POSTGRES_PORT" env-default:"5432"`
		Username string `env:"POSTGRES_USER" env-default:"postgres"`
		Password string `env:"POSTGRES_PASSWORD" env-default:"postgres"`
		DBName   string `env:"POSTGRES_DB" env-default:"postgres"`
		SSLMode  string `env:"POSTGRES_SSL_MODE" env-default:"disable"`
	}
	Parser struct {
		Host string `env:"PARSER_HOST" env-default:"127.0.0.1"`
		Port string `env:"PARSER_PORT" env-default:"8030"`
	}
}

var instance *Config
var once sync.Once

func GetConfig(logger logging.Logger) *Config {
	once.Do(func() {
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig(".env", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
