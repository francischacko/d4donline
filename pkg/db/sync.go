package db

import (
	"github.com/francischacko/d4donline/pkg/models"
)

func Syncdb() {
	DB.AutoMigrate(&models.OfferCompany{})
}
