package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var PriceAPIServer struct {
  API_PORT string
}

func loadServerConfig() {
	PriceAPIServer.API_PORT = "3000"
	if os.Getenv("API_PORT") != "" {
		PriceAPIServer.API_PORT = os.Getenv("API_PORT")
	}
}

func init() {
	loadServerConfig()
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}