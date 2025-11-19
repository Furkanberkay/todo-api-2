package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPAddr   string
	SQLitePath string
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		HTTPAddr:   getEnv("HTTP_ADDR", ":8080"),
		SQLitePath: getEnv("SQLITE_PATH", "./todo.db"),
	}

	log.Printf("[config] HTTP_ADDR=%s SQLITE_PATH=%s", cfg.HTTPAddr, cfg.SQLitePath)

	return cfg
}
