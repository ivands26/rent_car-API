package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ivands26/rent_car-API/app/config"
	"github.com/ivands26/rent_car-API/routes"
)

func main() {

	config.InitDB()

	router := gin.Default()
	routes.SetupRoute(router)
	router.Run()
}
