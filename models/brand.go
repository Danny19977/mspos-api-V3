package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	UUID string `gorm:"not null;unique" json:"uuid"`

	Name         string   `gorm:"not null" json:"name"`
	CountryUUID  string   `json:"country_uuid" gorm:"type:varchar(255);not null"`
	Country      Country  `gorm:"foreignKey:CountryUUID;references:UUID"`
	ProvinceUUID string   `json:"province_uuid" gorm:"type:varchar(255);not null"`
	Province     Province `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	Signature    string   `json:"signature"`

	PosFormItems []PosFormItems `gorm:"foreignKey:BrandUUID;references:UUID"`
}
