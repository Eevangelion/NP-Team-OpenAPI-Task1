package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	VK_APP_ID       string
	VK_SECRET_KEY   string
	VK_REDIRECT_URL string
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

var conf *Config = nil

func GetConfig() *Config {
	if conf == nil {
		conf = &Config{
			VK_APP_ID:       getEnv("VK_APP_ID", ""),
			VK_SECRET_KEY:   getEnv("VK_SECRET_KEY", ""),
			VK_REDIRECT_URL: getEnv("VK_REDIRECT_URL", ""),
		}
	}
	return conf
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
