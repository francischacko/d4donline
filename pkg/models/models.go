package models

import "time"

type OfferCompany struct {
	OfferID            int       `json:"offerId"`
	ClientID           int       `json:"clientId"`
	Country            string    `json:"country"`
	Image              string    `json:"image"`
	ImageWidth         int       `json:"imageWidth"`
	ImageHeight        int       `json:"imageHeight"`
	TextLocale         string    `json:"textHocale"`
	ValidityTextLocale string    `json:"validityTextLocale"`
	Position           int       `json:"position"`
	ValidFrom          time.Time `json:"validFrom"`
	ShowFrom           time.Time `json:"showFrom"`
	ValidTo            time.Time `json:"validTo"`
	Flag               uint      `json:"flag"`
	PageCount          uint      `json:"pageCount"`
	StoreURL           string    `json:"storeUrl"`
	StoreURLTitle      string    `json:"storeUrlTitle"`
	OfferHome          int       `json:"offerHome"`
}
