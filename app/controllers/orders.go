package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ivands26/rent_car-API/app/config"
	"github.com/ivands26/rent_car-API/app/models"
)

func MakeOrder(c *gin.Context) {
	var req models.RequestOrders
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to get request",
		})
		return
	}

	resCar, err := models.GetCarById(config.DB, req.CarId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get data",
		})
		return
	}

	if !resCar.Status {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to make order",
			"error":   fmt.Sprintf("%s Doesnt Avaiable", resCar.CarName),
		})
		return

	}
	parsePickup, _ := time.Parse("2006-01-02", req.PickupDate)
	parseDrop, _ := time.Parse("2006-01-02", req.DropoffDate)

	err = models.MakeOrder(config.DB, models.Orders{
		CarId:           req.CarId,
		OrderDate:       time.Now(),
		PickupDate:      parsePickup,
		DropoffDate:     parseDrop,
		PickupLocation:  req.PickupLocation,
		DropoffLocation: req.DropoffLocation,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to make order",
		})
		return
	}
	err = models.UpdateCarStatus(config.DB, req.CarId)

	c.JSON(http.StatusCreated, gin.H{
		"message": "success make order",
	})
	return
}

func GetOrderById(c *gin.Context) {
	param, bool := c.Params.Get("id")
	if !bool {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to get request",
		})
		return
	}

	res, err := models.GetOrderById(config.DB, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get data",
		"data":    res,
	})
	return
}
