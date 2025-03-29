package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Orders struct {
	OrderId         string    `gorm:"column:order_id;primary_key"`
	CarId           string    `gorm:"column:car_id" json:"car_id"`
	OrderDate       time.Time `gorm:"column:order_date;default:CURRENT_TIMESTAMP()" json:"order_date,string"`
	PickupDate      time.Time `gorm:"column:pickup_date" json:"pickup_date,string"`
	DropoffDate     time.Time `gorm:"column:dropoff_date" json:"dropoff_date,string"`
	PickupLocation  string    `gorm:"column:pickup_location" json:"pickup_location"`
	DropoffLocation string    `gorm:"column:dropoff_location" json:"dropoff_location"`
}

type RequestOrders struct {
	OrderId         string `gorm:"column:order_id;primary_key"`
	CarId           string `gorm:"column:car_id" json:"car_id"`
	OrderDate       string `gorm:"column:order_date;default:CURRENT_TIMESTAMP()" json:"order_date"`
	PickupDate      string `gorm:"column:pickup_date" json:"pickup_date"`
	DropoffDate     string `gorm:"column:dropoff_date" json:"dropoff_date"`
	PickupLocation  string `gorm:"column:pickup_location" json:"pickup_location"`
	DropoffLocation string `gorm:"column:dropoff_location" json:"dropoff_location"`
}

func (b *Orders) BeforeCreate(tx *gorm.DB) (err error) {
	b.OrderId = uuid.New().String()
	return
}

func MakeOrder(tx *gorm.DB, request Orders) error {
	err := tx.Debug().Table("orders").Create(&request).Error
	if err != nil {
		return err
	}

	return nil
}

func GetOrderById(tx *gorm.DB, request string) (*Orders, error) {
	res := Orders{}
	err := tx.Debug().Table("orders").Where("order_id = ?", request).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}
