package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
	"time"
)

type Config struct {
	App struct {
		Port string `yaml:"port"`
	} `yaml:"app"`

	PostgreSQL struct {
		Username string `yaml:"username" env:"POSTGRES_USER"`
		Password string `yaml:"password" env:"POSTGRES_PASS"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"postgreSQL"`

	Auth struct {
		Salt string `yaml:"salt" env:"PASS_SALT"`
		JWT  struct {
			SigningKey      string        `env-required:"true" yaml:"signingKey" env:"SIGNING_KEY"`
			AccessTokenTTL  time.Duration `env-required:"true" yaml:"accessTokenTTL"`
			RefreshTokenTTL time.Duration `env-required:"true" yaml:"refreshTokenTTL"`
		} `yaml:"jwt"`
	} `yaml:"auth"`
}

var cfg = &Config{}
var once sync.Once

func GetConfig() *Config {

	once.Do(func() {
		err := cleanenv.ReadConfig("config/config.yml", cfg)
		if err != nil {
			return
		}
	})

	return cfg
}
