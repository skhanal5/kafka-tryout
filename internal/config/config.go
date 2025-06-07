package config

import (
	"log"
	"os"
)

type Config struct {
	KafkaBroker string
	KafkaTopic  string
}

func NewConfig() Config {
	return Config{
		KafkaBroker: getEnvOrFail("KAFKA_BROKER"),
		KafkaTopic:  getEnvOrFail("KAFKA_TOPIC"),
	}
}

func getEnvOrFail(key string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("required environment variable %s is not set", key)
	}
	return val
}
