package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	prod := os.Getenv("PROD")

	if prod != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}
	uri := os.Getenv("MONGO_URL")
	if uri == "" {
		log.Fatal("Error loading .env file")
	}

	return uri
}
