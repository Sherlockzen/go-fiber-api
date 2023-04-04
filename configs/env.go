package configs

import (
	"log"
	"os"
)

func EnvMongoURI() string {
	// err := godotenv.Load()
	uri := os.Getenv("MONGO_URL")
	if uri == "" {
		log.Fatal("Error loading .env file")
	}

	return uri
}
