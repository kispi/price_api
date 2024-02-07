package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

var PriceAPIServer struct {
  API_PORT string
}

func init() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  PriceAPIServer.API_PORT = "3000"
  if os.Getenv("API_PORT") != "" {
    PriceAPIServer.API_PORT = os.Getenv("API_PORT")
  }
}

func main() {
  app := fiber.New()

  app.Get("/", func(c fiber.Ctx) error {
    return c.JSON(fiber.Map{
      "message": "Hello, World!",
    })
  })

  app.Listen(":" + PriceAPIServer.API_PORT)
}
