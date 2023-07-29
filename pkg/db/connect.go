package db

import (
	"github.com/francischacko/d4donline/pkg/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connectdb() {

	var err error
	dsn := configs.EnvConf.DbConnect

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed")
	}
}
