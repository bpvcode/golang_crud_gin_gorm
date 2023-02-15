package controllers

import "github.com/gin-gonic/gin"

func Router(ge *gin.Engine) {

	apiRoutes := ge.Group("/api")

	shipment := apiRoutes.Group("/shipment")
	shipment.POST("", CreateShipment)
	shipment.GET("", GetAllShipments)
	shipment.GET("/:id", GetShipment)
	shipment.PUT("/:id", UpdateShipment)
	shipment.DELETE("/:id", DeleteShipment)
}
