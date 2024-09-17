package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string `env:"ENV" env-default:"local"`
	HTTP     HTTPConfig
	Postgres PostgresConfig
}

type HTTPConfig struct {
	Address string `env:"SERVER_ADDRESS" env-default:":8080"`
}

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port     string `env:"POSTGRES_PORT" env-default:"5432"`
	Username string `env:"POSTGRES_USERNAME" env-default:"postgres"`
	DbName   string `env:"POSTGRES_DATABASE" env-default:"postgres"`
	SslMode  string `env:"SSL_MODE" env-default:"disable"`
	Password string `env:"PASSWORD" env-default:"postgres"`
}

func InitConfig() Config {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic("cannot read env: " + err.Error())
	}
	return cfg
}
