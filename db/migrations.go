package db

import "github.com/bpvcode/golang_crud_gin_gorm/models"

func MigrateTables() {
	db := GetDB()
	db.AutoMigrate(&models.Shipment{}) // Migrates the table with no data
}

func MigrateShipmentsMockData() {
	ship := models.Shipment{
		CompanyID:      "00001",
		TrackingNumber: "FDHX00001",
	}

	ship1 := models.Shipment{
		CompanyID:      "00002",
		TrackingNumber: "FDHX00002",
	}

	ship2 := models.Shipment{
		CompanyID:      "00003",
		TrackingNumber: "FDHX00003",
	}

	db := GetDB()
	db.Save(&ship)  // Migrates shipment data
	db.Save(&ship1) // Migrates shipment data
	db.Save(&ship2) // Migrates shipment data
}
