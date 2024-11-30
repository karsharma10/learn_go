package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	OpenAIKey = "OPENAI_KEY"
)

type Configs struct {
	OpenAIKey string
	Host      string
	Port      string
	User      string
	Password  string
	Dbname    string
}

type Config func(*Configs)

func WithOpenAI() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	key := os.Getenv(OpenAIKey)
	return func(configs *Configs) {
		configs.OpenAIKey = key
	}
}

func WithDb() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	return func(configs *Configs) {
		configs.Host = os.Getenv("DB_HOST")
		configs.Port = os.Getenv("DB_PORT")
		configs.User = os.Getenv("DB_USER")
		configs.Password = os.Getenv("DB_PASSWORD")
		configs.Dbname = os.Getenv("DB_NAME")
	}
}
