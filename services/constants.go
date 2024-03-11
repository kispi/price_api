package services

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v3/log"
)

type Country struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Emoji string `json:"emoji"`
}

type ConstantsService struct{}

var countries []*Country

func (s *ConstantsService) GetCountries() []*Country {
	if countries != nil {
		return countries
	}

	log.Debug("Loading countries")
	file, err := os.ReadFile("data/countries.json")
	if err != nil {
		log.Error(err)
		return nil
	}

	err = json.Unmarshal(file, &countries)
	if err != nil {
		log.Error(err)
		return nil
	}

	return countries
}
