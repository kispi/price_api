package main

import (
	"os"
	"time"

	"github.com/kispi/price_api/controllers"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

var defaultConfig = logger.Config{
	Next:          nil,
	Done:          nil,
	Format:        "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${queryParams}\n",
	TimeFormat:    time.RFC3339,
	TimeZone:      "Local",
	TimeInterval:  500 * time.Millisecond,
	Output:        os.Stdout,
	DisableColors: false,
}

func main() {
	app := fiber.New(fiber.Config{
		ProxyHeader: fiber.HeaderXForwardedFor,
	})

	app.Use(logger.New(defaultConfig))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://coinsect.io, https://btc.coinsect.io, http://localhost:4001",
	}))

	app.Get("/bitcoin/price", controllers.Price)
	app.Get("/bitcoin/quotes", controllers.Quotes)

	app.Listen(":" + ServerSettings.API_PORT)
}
