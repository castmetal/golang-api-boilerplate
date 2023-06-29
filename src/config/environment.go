package config

import "github.com/vrischmann/envconfig"

var Env struct {
	Logger struct {
		Env string `envconfig:"default=production"`
	}

	Server struct {
		Type string `envconfig:"default=http"`
		Port string `envconfig:"default=8088"`
	}

	ApiName    string `envconfig:"default=API - Example"`
	ApiVersion string `envconfig:"default=v1"`

	Environment string `envconfig:"default=dev"`

	DB struct {
		URL string `envconfig:"default=postgres://postgres:supersecret@localhost:5432/postgres?sslmode=disable"`
	}

	RateLimit int `envconfig:"default=60"`

	Cache struct {
		Addr     string `envconfig:"default=0.0.0.0"`
		Port     string `envconfig:"default=6379"`
		Database string `envconfig:"default=0"`
		Password string `envconfig:"default= "`
	}

	RabbitMq struct {
		URL string `envconfig:"default=amqp://guest:guest@localhost:5672"`
	}
}

type EnvStruct struct {
	Logger struct {
		Env string
	}

	Server struct {
		Type string
		Port string
	}

	ApiName    string
	ApiVersion string

	Environment string

	DB struct {
		URL string
	}

	RateLimit int

	Cache struct {
		Addr     string
		Port     string
		Database string
		Password string
	}

	RabbitMq struct {
		URL string
	}
}

func Init() error {
	return envconfig.InitWithPrefix(&Env, "CF")
}
