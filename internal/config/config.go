package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	TgBotToken string
	TgChatID   int64
}

func Load() (*Config, error) {
	_ = godotenv.Load(".env.local", ".env")

	cfg := &Config{
		TgBotToken: os.Getenv("TG_BOT_TOKEN"),
	}

	var err error
	cfg.TgChatID, err = getEnvInt64("TG_CHAT_ID")
	if err != nil {
		return nil, err
	}

	if cfg.TgBotToken == "" {
		return nil, errors.New("TG_BOT_TOKEN is not set")
	}

	return cfg, nil
}

func getEnvInt64(key string) (int64, error) {
	v := os.Getenv(key)
	if v == "" {
		return 0, fmt.Errorf("%s is not set", key)
	}
	return strconv.ParseInt(v, 10, 64)
}
