package env

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	DSN         string
	Port        string
	LogFilePath string
	SQLLogLevel string
	BaseURL     string
}

func GetEnv() Env {
	var s Env
	err := envconfig.Process("graphql", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	return s
}
