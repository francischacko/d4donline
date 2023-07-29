package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestGetDataByCountry(t *testing.T) {

	app := fiber.New()

	tests := []struct {
		country          string
		request          string
		expectedStatus   int
		expectedResponse string
	}{
		{
			country:          "success",
			request:          "/getoffers/US",
			expectedStatus:   fiber.StatusOK,
			expectedResponse: `[{"OfferID": 307470, "ClientID": 1001, "Country": "US", "Image": "image1.jpg", "ImageWidth": 800, "ImageHeight": 600}]`,
		},
		{
			country:          "invalid country",
			request:          "/getoffers/USSR",
			expectedStatus:   fiber.StatusInternalServerError,
			expectedResponse: `{"error": "Invalid country parameter"}`,
		},
	}

	// Loop through the test cases
	for _, tt := range tests {
		t.Run(tt.country, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, tt.request, nil)
			resp, err := app.Test(req)
			require.NoError(t, err)

			// Check the actual status code against the expected one
			require.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Read the response body
			body, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)

			// Check the actual response body against the expected one
			require.JSONEq(t, tt.expectedResponse, string(body))
		})
	}
}
