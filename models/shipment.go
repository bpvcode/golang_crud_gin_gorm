package models

import "gorm.io/gorm"

type Shipment struct {
	gorm.Model
	CompanyID      string `gorm:"unique;not null;type:varchar(100);default:null"`
	TrackingNumber string `gorm:"unique;not null;type:varchar(100);default:null"`
}
