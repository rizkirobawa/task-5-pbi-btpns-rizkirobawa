package database

import (
	"github.com/rizkirobawa/task-5-pbi-btpns-rizki/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/task5_pbi?parseTime=true"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Photo{})

	DB = db
}
