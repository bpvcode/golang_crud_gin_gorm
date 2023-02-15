package repositories

import (
	"log"

	"github.com/bpvcode/golang_crud_gin_gorm/db"
	"github.com/bpvcode/golang_crud_gin_gorm/models"
)

func CreateShipment(shipment models.Shipment) error {
	result := db.DB.Create(&shipment)
	if result.Error != nil {
		log.Println("Failed to find all shipments", result.Error)
		return result.Error
	}
	return nil
}

func GetAllShipments() ([]models.Shipment, error) {
	var shipments []models.Shipment
	result := db.DB.Find(&shipments)
	if result.Error != nil {
		log.Println("Failed to find all shipments", result.Error)
		return []models.Shipment{}, result.Error
	}
	return shipments, nil
}

func GetShipment(id string) (models.Shipment, error) {
	var shipment models.Shipment
	result := db.DB.First(&shipment, id)
	if result.Error != nil {
		log.Printf("Failed to find shipment id ID: %v\nERROR: %v\n", id, result.Error)
		return models.Shipment{}, result.Error
	}
	return shipment, nil
}

func UpdateShipment(id string, data *models.Shipment) error {
	var shipment models.Shipment
	result := db.DB.First(&shipment, id)
	if result.Error != nil {
		log.Printf("Failed to find shipment id ID: %v\nERROR: %v\n", id, result.Error)
		return result.Error
	}
	db.DB.Model(&shipment).Updates(data)
	return nil
}

func DeleteShipment(id string) error {
	result := db.DB.Delete(&models.Shipment{}, id)
	if result.Error != nil {
		log.Printf("Failed to delete shipment id ID: %v\nERROR: %v\n", id, result.Error)
		return result.Error
	}
	return nil
}
