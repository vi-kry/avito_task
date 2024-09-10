package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env      string         `yaml:"env" env-default:"local"`
	HTTP     HTTPConfig     `yaml:"http"`
	Postgres PostgresConfig `yaml:"postgres"`
}

type HTTPConfig struct {
	Port int `yaml:"port"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name"`
	SslMode  string `yaml:"ssl_mode"`
}

func InitConfig() Config {
	configPath := getConfigPath()

	if configPath == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return cfg
}

func getConfigPath() string {
	var path string

	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
