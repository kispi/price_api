package controllers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kispi/price_api/services"
)

func GetCountries(c fiber.Ctx) error {
	constantsService := &services.ConstantsService{}
	c.JSON(constantsService.GetCountries())
	return nil
}
