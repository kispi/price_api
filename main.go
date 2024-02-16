package main

import (
	"os"
	"time"

	"github.com/kispi/price_api/controllers"

	"github.com/gofiber/fiber/v3"
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

	// Logger middleware
	app.Use(logger.New(defaultConfig))

	app.Get("/prices/bitcoin", controllers.Bitcoin)

	app.Listen(":" + ServerSettings.API_PORT)
}
