package bootstrap

import (
	"github.com/joho/godotenv"
)

// Returns Environment variables as a map
func NewEnv() *map[string]string {
	EnvFile := ".env"
	Env, err := godotenv.Read(EnvFile)
	if err != nil {
		panic(err)
	}
	return &Env
}
