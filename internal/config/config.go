package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	TelegramToken      string
	OpenAIToken        string
	Debug              bool
	ContextPoolTimeout int
	DSN                string
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is required")
	}

	openaiToken := os.Getenv("OPENAI_TOKEN")
	if openaiToken == "" {
		log.Fatal("OPENAI_TOKEN is required")
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG_MODE"))
	if err != nil {
		debug = false
	}

	contextPoolTimeout, err := strconv.Atoi(os.Getenv("CONTEXT_POOL_TIMEOUT"))
	if err != nil {
		contextPoolTimeout = 5
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = "database.db"
	}

	return &Config{
		TelegramToken:      token,
		OpenAIToken:        openaiToken,
		Debug:              debug,
		ContextPoolTimeout: contextPoolTimeout,
		DSN:                dsn,
	}
}
