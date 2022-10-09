package main

import (
	"assignment2/config"
	"assignment2/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/orders", inDB.CreateOrder)
	router.GET("/orders", inDB.GetOrders)
	router.PUT("/orders/:orderId", inDB.UpdateOrder)
	router.DELETE("/orders/:orderId", inDB.DeleteOrder)
	router.Run(":8080")

}
