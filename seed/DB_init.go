package seed

import (
	"fmt"

	"github.com/nihil_sum/Golang_IM/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(config_path string) {
	dsn := "nihil:050903@tcp(192.168.137.1:3306)/devdb"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database failed to load")
	}
	db.AutoMigrate(&models.UserBasic{})
	user := &models.UserBasic{}
	user.Name = "Dylan Liu"
	db.Create(user)
	fmt.Println(db.First(user, 1))
	db.Model(user).Update("Password", "050903")
}
