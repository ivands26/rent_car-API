package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ivands26/rent_car-API/app/config"
	"github.com/ivands26/rent_car-API/app/models"
)

func GetAllCars(c *gin.Context) {
	res, err := models.GetAllCars(config.DB)
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

func GetCarById(c *gin.Context) {
	param, bool := c.Params.Get("id")
	if !bool {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to get request",
		})
		return
	}

	res, err := models.GetCarById(config.DB, param)
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

func InsertCar(c *gin.Context) {
	var req models.Car
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to get request",
		})
		return
	}

	err = models.InsertCar(config.DB, req)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to insert data",
				"error":   "Car Name Already Exist",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to insert data",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "success insert data",
	})
	return
}
