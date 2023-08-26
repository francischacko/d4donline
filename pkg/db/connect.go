package db

import (
	"github.com/francischacko/d4donline/pkg/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Connectdb() {

	var err error
	dsn := configs.EnvConf.DbConnect

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed")
	}
}

// &gorm.Config{
//     NamingStrategy: schema.NamingStrategy{
//         SingularTable: true,
//     },
