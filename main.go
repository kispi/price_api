package main

import (
	"time"

	"github.com/kispi/price_api/controllers"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
  app := fiber.New()

  // Logger middleware
  app.Use(logger.New(logger.Config{
    TimeFormat: time.RFC3339,
  }))

  app.Get("/prices/bitcoin", controllers.Bitcoin)

  app.Listen(":" + ServerSettings.API_PORT)
}
