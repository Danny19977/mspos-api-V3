package models

import (
	"gorm.io/gorm"
)

type SubArea struct {
	gorm.Model
	UUID string `gorm:"not null;unique" json:"uuid"`

	Name string `gorm:"not null" json:"name"`

	CountryUUID  string   `json:"country_uuid" gorm:"type:varchar(255);not null"`
	Country      Country  `gorm:"foreignKey:CountryUUID;references:UUID"`
	ProvinceUUID string   `json:"province_uuid" gorm:"type:varchar(255);not null"`
	Province     Province `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	AreaUUID     string   `json:"area_uuid" gorm:"type:varchar(255);not null"`
	Area         Area     `gorm:"foreignKey:AreaUUID;references:UUID"`

	Signature string `json:"signature"`

	Communes []Commune `gorm:"foreignKey:SubAreaUUID;references:UUID"`

	Pos      []Pos     `gorm:"foreignKey:SubAreaUUID;references:UUID"`
	Posforms []PosForm `gorm:"foreignKey:SubAreaUUID;references:UUID"`

	// Drs    []Dr    `gorm:"foreignKey:SubAreaUUID;references:UUID"`
	// Cyclos []Cyclo `gorm:"foreignKey:ProvinceUUID;references:UUID"`

	RoutePlan []RoutePlan `gorm:"foreignKey:SubAreaUUID;references:UUID"`

	Users []User `gorm:"foreignKey:SubAreaUUID;references:UUID"`
}
