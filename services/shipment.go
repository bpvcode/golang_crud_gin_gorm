package services

import (
	"github.com/bpvcode/golang_crud_gin_gorm/models"
	"github.com/bpvcode/golang_crud_gin_gorm/repositories"
)

func CreateShipment(body models.Shipment) error {
	shipment := models.Shipment{
		CompanyID:      body.CompanyID,
		TrackingNumber: body.TrackingNumber,
	}

	return repositories.CreateShipment(shipment)
}

func GetAllShipments() ([]models.Shipment, error) {
	return repositories.GetAllShipments()
}

func GetShipment(id string) (models.Shipment, error) {
	return repositories.GetShipment(id)
}

func UpdateShipment(id string, shipment models.Shipment) error {
	return repositories.UpdateShipment(id, &shipment)
}

func DeleteShipment(id string) error {
	return repositories.DeleteShipment(id)
}
