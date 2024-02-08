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
	Date string `json:"d"`
}

type ServiceRequestBitcoin struct {
	Limit int64
	Offset int64
	Timeframe string
}

type PriceService struct {}

var priceData []*PriceRow

// Once the price data is loaded, it is cached in the priceData.
func loadBitcoinPriceData() []*PriceRow {
	if priceData != nil {
		return priceData
	}

	log.Debug("Loading bitcoin price data")
	file, err := os.ReadFile("data/price_bitcoin.json")
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

func (s *PriceService) Bitcoin(r *ServiceRequestBitcoin) *fiber.Map {
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
		"data":  result,
		"total": total,
		"limit": r.Limit,
		"offset": r.Offset,
	}
	
	return response
}