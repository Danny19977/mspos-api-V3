package models

import (
	// "github.com/lib/pq"
	"gorm.io/gorm"
)

type Area struct {
	gorm.Model

	UUID string `gorm:"not null;unique" json:"uuid"`
	Name string `gorm:"not null" json:"name"`

	CountryUUID  string   `json:"country_uuid" gorm:"type:varchar(255);not null"`
	Country      Country  `gorm:"foreignKey:CountryUUID;references:UUID"`
	ProvinceUUID string   `json:"province_uuid" gorm:"type:varchar(255);not null"`
	Province     Province `gorm:"foreignKey:ProvinceUUID;references:UUID"`

	Signature string `json:"signature"`

	SubAreas []SubArea `gorm:"foreignKey:AreaUUID;references:UUID"`
	Communes []Commune `gorm:"foreignKey:AreaUUID;references:UUID"`

	// Sups   []Sup   `gorm:"foreignKey:AreaUUID;references:UUID"`
	// Drs    []Dr    `gorm:"foreignKey:AreaUUID;references:UUID"`
	// Cyclos []Cyclo `gorm:"foreignKey:AreaUUID;references:UUID"`

	Pos      []Pos     `gorm:"foreignKey:AreaUUID;references:UUID"`
	PosForms []PosForm `gorm:"foreignKey:AreaUUID;references:UUID"`

	RoutePlans []RoutePlan `gorm:"foreignKey:AreaUUID;references:UUID"`

	Users []User `gorm:"foreignKey:AreaUUID;references:UUID"`
}
