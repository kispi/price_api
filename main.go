package main

import (
	localpackage "github.com/kispi/price_api/localpackage"

	"github.com/gofiber/fiber/v3"
)

func main() {
  app := fiber.New()

  app.Get("/", func(c fiber.Ctx) error {
    return c.JSON(fiber.Map{
      "message": localpackage.Foo(),
    })
  })

  app.Listen(":" + PriceAPIServer.API_PORT)
}
