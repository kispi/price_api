package services

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type PriceRow struct {
	Price float64 `json:"o"`
	Date  string  `json:"d"`
}

type ServiceRequestBitcoinPrice struct {
	Limit     int64
	Offset    int64
	Timeframe string
}

type BitcoinService struct{}

var priceData []*PriceRow

type QuoteRow struct {
	Name   string `json:"name"`
	Quotes []struct {
		Text   string `json:"text"`
		Source string `json:"source"`
		Date   string `json:"date"`
	} `json:"quotes"`
}

var quotesData []*QuoteRow

// Once the price data is loaded, it is cached in the priceData.
// Source: https://www.investing.com/crypto/bitcoin/historical-data
// API isn't available, so hardcoded json file is used.
func loadBitcoinPriceData() []*PriceRow {
	if priceData != nil {
		return priceData
	}

	log.Debug("Loading bitcoin price data")
	file, err := os.ReadFile("data/bitcoin_price.json")
	if err != nil {
		log.Error(err)
		return nil
	}

	err = json.Unmarshal(file, &priceData)
	if err != nil {
		log.Error(err)
		return nil
	}

	return priceData
}

func (s *BitcoinService) Price(r *ServiceRequestBitcoinPrice) *fiber.Map {
	priceData := loadBitcoinPriceData()

	result := make([]*PriceRow, 0)
	for _, row := range priceData {
		t, err := time.Parse("2006-01-02", row.Date)
		if err != nil {
			log.Error(err)
			continue
		}

		switch r.Timeframe {
		case "year":
			_, month, day := t.Date()
			if month == 1 && day == 1 {
				result = append(result, row)
			}
		case "month":
			_, _, day := t.Date()
			if day == 1 {
				result = append(result, row)
			}
		case "week":
			weekday := t.Weekday()
			if weekday == time.Monday {
				result = append(result, row)
			}
		default:
			result = append(result, row)
		}
	}

	total := len(result)

	if int(r.Offset) >= total {
		result = []*PriceRow{}
	} else {
		result = result[int(r.Offset):]
	}

	if len(result) > int(r.Limit) {
		result = result[:r.Limit]
	}

	// Prepare response map
	response := &fiber.Map{
		"data":   result,
		"total":  total,
		"limit":  r.Limit,
		"offset": r.Offset,
	}

	return response
}

func loadBitcoinQuotes() []*QuoteRow {
	if quotesData != nil {
		return quotesData
	}

	log.Debug("Loading bitcoin quotes")
	file, err := os.ReadFile("data/bitcoin_quotes.json")
	if err != nil {
		log.Error(err)
		return nil
	}

	err = json.Unmarshal(file, &quotesData)
	if err != nil {
		log.Error(err)
		return nil
	}

	return quotesData
}

func (s *BitcoinService) Quotes() []*QuoteRow {
	return loadBitcoinQuotes()
}
