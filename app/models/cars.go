package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Car struct {
	CarId     string  `gorm:"column:car_id;primary_key" json:"car_id"`
	CarName   string  `gorm:"column:car_name;unique" json:"car_name"`
	DayRate   float64 `gorm:"column:day_rate" json:"day_rate"`
	MonthRate float64 `gorm:"column:month_rate" json:"month_rate"`
	ImageCar  string  `gorm:"column:image_car" json:"image_car"`
	Status    bool    `gorm:"colum:status;default:true" json:"status"`
}

func (b *Car) BeforeCreate(tx *gorm.DB) (err error) {
	b.CarId = uuid.New().String()
	return
}

func GetAllCars(tx *gorm.DB) (*[]Car, error) {
	res := []Car{}
	err := tx.Debug().Table("cars").Scan(&res).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func GetCarById(tx *gorm.DB, request string) (*Car, error) {
	res := Car{}
	err := tx.Debug().Table("cars").Where("car_id = ?", request).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func InsertCar(tx *gorm.DB, request Car) error {

	err := tx.Debug().Create(&request).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCarStatus(tx *gorm.DB, request string) error {

	err := tx.Debug().Table("cars").Where("car_id = ?", request).Update("status", false).Error
	if err != nil {
		return err
	}
	return nil
}
