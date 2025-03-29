package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ivands26/rent_car-API/app/controllers"
)

func SetupRoute(router *gin.Engine) {
	//CAR
	car := router.Group("/car")
	car.GET("/get", controllers.GetAllCars)
	car.GET("/get/:id", controllers.GetCarById)
	car.POST("/insert", controllers.InsertCar)

	//ORDER
	order := router.Group("/order")
	order.POST("/insert", controllers.MakeOrder)
	order.GET("/get/:id", controllers.GetOrderById)
}
