package config

import (
	"log"

	"github.com/joho/godotenv"
)




func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot Load .env File")
	}

	

	
}