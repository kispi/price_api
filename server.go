package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type PriceAPIServer struct {
	API_PORT string
}

var ServerSettings = &PriceAPIServer{
	API_PORT: "3000",
}

func loadServerConfig() {
	if os.Getenv("API_PORT") != "" {
		ServerSettings.API_PORT = os.Getenv("API_PORT")
	}
}

func init() {
	loadServerConfig()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
