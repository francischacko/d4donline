package controllers

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/francischacko/d4donline/pkg/db"
	"github.com/francischacko/d4donline/pkg/models"
	"github.com/gofiber/fiber/v2"
)

var countryDataMap = sync.Map{}

func HelloWorld(c *fiber.Ctx) {
	c.JSON(fiber.Map{
		"message": "Hello World",
	})
}

type OffersDetails interface {
	GetDataByCountry()
}

func GetDataByCountry(c *fiber.Ctx) error {
	fmt.Println("function got called")
	// Get the country parameter from the URL
	country := c.Params("country")
	// SQL query with the USE INDEX clause
	query := `
        SELECT * FROM offer_company  FORCE INDEX (idx_country,idx_valid_from)
        WHERE country = ? AND valid_from <= ? AND valid_to >= ? AND flag = ?
        ORDER BY priority
    `
	rows, err := db.DB.Raw(query, country, time.Now(), time.Now(), 1).Rows()
	if err != nil {

		// error handling
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch data",
		})
		return err
	}
	defer rows.Close()

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Processing the query results concurrently using goroutines
	// var offers []models.OfferCompany
	for rows.Next() {
		var offer models.OfferCompany
		if err := db.DB.ScanRows(rows, &offer); err != nil {
			log.Fatalf("Error scanning row:%v", err)
		}

		// Incrementing the WaitGroup counter before starting a new goroutine
		wg.Add(1)
		go func(offer models.OfferCompany) {
			// Decrementing the WaitGroup counter when the goroutine is done
			defer wg.Done()

			// Processing the offer data

			countryDataMap.Store(country, append(getCountryData(country), offer))
		}(offer)
	}

	// Waiting for all goroutines to finish before proceeding
	wg.Wait()
	// Retrieving the updated data from countryDataMap
	updatedOffers := getCountryData(country)
	if updatedOffers != nil {
		c.JSON(updatedOffers)
	} else {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No data found for the specified country",
		})
	}

	return nil
}

func getCountryData(country string) []models.OfferCompany {
	data, ok := countryDataMap.Load(country)
	if !ok {
		return nil
	}

	offers, ok := data.([]models.OfferCompany)
	if !ok {
		return nil
	}

	return offers
}
