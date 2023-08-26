package models

import "time"

// type OfferCompany struct {
// 	OfferID            int       `json:"offerId"`
// 	ClientID           int       `json:"clientId"`
// 	Country            string    `json:"country"`
// 	Image              string    `json:"image"`
// 	ImageWidth         int       `json:"imageWidth"`
// 	ImageHeight        int       `json:"imageHeight"`
// 	TextLocale         string    `json:"textHocale"`
// 	ValidityTextLocale string    `json:"validityTextLocale"`
// 	Position           int       `json:"position"`
// 	ValidFrom          time.Time `json:"validFrom"`
// 	ShowFrom           time.Time `json:"showFrom"`
// 	ValidTo            time.Time `json:"validTo"`
// 	Flag               uint      `json:"flag"`
// 	PageCount          uint      `json:"pageCount"`
// 	StoreURL           string    `json:"storeUrl"`
// 	StoreURLTitle      string    `json:"storeUrlTitle"`
// 	OfferHome          int       `json:"offerHome"`
// }

type OfferCompany struct {
	Idoffer_company         uint64    `json:"idoffercompany"`
	Idcompany               uint64    `json:"idcompany"`
	Idbrand                 uint64    `json:"idbrand"`
	Idfirm_sub_category     uint64    `json:"idfirmsubcategory"`
	Idproduct_main_category uint64    `json:"idproductmaincategory"`
	Country                 string    `json:"country"`
	Thumb_url               string    `json:"thumburl"`
	Image_url               string    `json:"imageurl"`
	Crop_url                string    `json:"crop_url"`
	Original_url            string    `json:"originalurl"`
	Thumb_url_ng            string    `json:"thumburlng"`
	Image_url_ng            string    `json:"imageurlng"`
	Thumb_width             int       `json:"thumbwidth"`
	Thumb_height            int       `json:"thumbheight"`
	Image_width             int       `json:"imagewidth"`
	Image_height            int       `json:"imageheight"`
	Text_eng                string    `json:"texteng"`
	Text_ar                 string    `json:"textar"`
	Validity_text           string    `json:"validitytext"`
	Validity_text_ar        string    `json:"validitytextar"`
	Priority                int       `json:"priority"`
	Valid_from              time.Time `json:"validfrom"`
	Show_from               time.Time `json:"showfrom"`
	Valid_to                time.Time `json:"validto"`
	Visitors                int       `json:"visitors"`
	Click_count             int       `json:"clickcount"`
	Featured                int       `json:"featured"`
	Category                uint64    `json:"category"`
	Type                    int       `json:"type"`
	Flag                    uint      `json:"flag"`
	Group_no                uint      `json:"groupno"`
	Auto_hidden             []byte    `json:"autohidden"`
	View_multiplier         float64   `json:"viewmultiplier"`
	Products_available      []byte    `json:"productsavailable"`
	Page_count              uint      `json:"pagecount"`
	Adl_header              string    `json:"adlheader"`
	Description             string    `json:"description"`
	Store_url               string    `json:"storeurl"`
	Store_url_text          int       `json:"storeurltext"`
	Offer_home              int       `json:"offerhome"`
	Sponsored_from          time.Time `json:"sponsoredfrom"`
	Sponsored_to            time.Time `json:"sponsoredto"`
	Qa_answer               string    `json:"qaanswer"`
	Nearest                 uint      `json:"nearest"`
	LinkText                string    `json:"linktext"`
	LastEditDate            time.Time `json:"LastEditDate"`
	CreationDate            time.Time `json:"CreationDate"`
	Offer_id                uint64    `json:"offerid"`
	Client_id               uint64    `json:"clientid"`
}
