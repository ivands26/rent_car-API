package config

import (
	"log"

	"github.com/ivands26/rent_car-API/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "host=localhost user=root password=P@ssw0rd dbname=project port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect Database")
	}

	DB.AutoMigrate(&models.Car{})
	DB.AutoMigrate(&models.Orders{})
}
