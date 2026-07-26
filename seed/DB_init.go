package seed

import (
	"github.com/nihil_sum/Golang_IM/confs"
	"github.com/nihil_sum/Golang_IM/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbc, errno := confs.LoadDBConfig()
	if errno != nil {
		panic("Failed to load the database configurations!")
	}
	dsn := dbc.DSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database failed to load")
	}
	db.AutoMigrate(&models.UserBasic{})

	return db
}
