package core

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port     string
	MysqlUrl string
}

func NewEnv() *Env {
	godotenv.Load()

	env := Env{
		Port:     os.Getenv("PORT"),
		MysqlUrl: os.Getenv("MYSQL_URL"),
	}

	return &env
}
