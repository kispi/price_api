package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/kispi/price_api/services"
)

const (
	queryLimit = 200
)

func Price(c fiber.Ctx) error {
	timeframe := c.Query("timeframe")
	if timeframe != "" && timeframe != "month" && timeframe != "week" && timeframe != "year" {
		c.Status(400).JSON(fiber.Map{
			"error": "timeframe must be 'month' or 'week' or 'year'",
		})
		return c.Context().Err()
	}

	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	if limit > queryLimit {
		c.Status(400).JSON(fiber.Map{
			"error": "limit must be less than " + strconv.FormatInt(queryLimit, 10) + " or empty",
		})
		return c.Context().Err()
	}

	if limit == 0 {
		limit = 20
	}

	offset, _ := strconv.ParseInt(c.Query("offset"), 10, 64)
	serviceRequestBitcoinPrice := &services.ServiceRequestBitcoinPrice{
		Limit:     limit,
		Timeframe: timeframe,
		Offset:    offset,
	}

	bitcoinService := &services.BitcoinService{}
	c.JSON(bitcoinService.Price(serviceRequestBitcoinPrice))
	return nil
}

func Quotes(c fiber.Ctx) error {
	bitcoinService := &services.BitcoinService{}
	c.JSON(bitcoinService.Quotes())
	return nil
}
