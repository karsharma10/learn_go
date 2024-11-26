package config

import (
	"github.com/joho/godotenv"
	"os"
)

const (
	OpenAIKey = "OPENAI_KEY"
)

type Config struct {
	OpenAIKey string
}

func InitConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	return &Config{
		OpenAIKey: os.Getenv(OpenAIKey),
	}, nil
}
