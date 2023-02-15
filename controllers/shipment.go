package controllers

import (
	"net/http"

	"github.com/bpvcode/golang_crud_gin_gorm/models"
	"github.com/bpvcode/golang_crud_gin_gorm/services"
	"github.com/gin-gonic/gin"
)

func CreateShipment(c *gin.Context) {
	var body struct {
		CompanyID      string
		TrackingNumber string
	}

	err := c.Bind(&body)

	err = services.CreateShipment(models.Shipment{
		CompanyID:      body.CompanyID,
		TrackingNumber: body.TrackingNumber,
	})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Shipment created successfully",
	})
}

func GetAllShipments(c *gin.Context) {
	shipments, err := services.GetAllShipments()
	if len(shipments) == 0 {
		c.AbortWithError(http.StatusNotFound, err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
		return
	}
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, shipments)
}

func GetShipment(c *gin.Context) {
	id := c.Param("id") // GET PARAM FROM URL

	shipment, err := services.GetShipment(id)
	if shipment.ID == 0 {
		c.AbortWithError(http.StatusNotFound, err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
		return
	}
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, shipment)
}

func UpdateShipment(c *gin.Context) {
	id := c.Param("id") // GET PARAM FROM URL

	var body struct {
		CompanyID      string
		TrackingNumber string
	}

	c.Bind(&body)

	shipment := models.Shipment{
		CompanyID:      body.CompanyID,
		TrackingNumber: body.TrackingNumber,
	}

	err := services.UpdateShipment(id, shipment)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Shipments updated successfully",
	})
}

func DeleteShipment(c *gin.Context) {
	id := c.Param("id") // GET PARAM FROM URL

	err := services.DeleteShipment(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Shipment deleted successfully",
	})
}
