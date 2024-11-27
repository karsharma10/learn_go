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
