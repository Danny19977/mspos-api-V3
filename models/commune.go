package models

import "gorm.io/gorm"

type Commune struct {
	gorm.Model
	UUID string `gorm:"not null;unique" json:"uuid"`

	Name string `gorm:"not null" json:"name"`

	CountryUUID  string   `json:"country_uuid" gorm:"type:varchar(255);not null"`
	Country      Country  `gorm:"foreignKey:CountryUUID;references:UUID"`
	ProvinceUUID string   `json:"province_uuid" gorm:"type:varchar(255);not null"`
	Province     Province `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	AreaUUID     string   `json:"area_uuid" gorm:"type:varchar(255);not null"`
	Area         Area     `gorm:"foreignKey:AreaUUID;references:UUID"`
	SubAreaUUID  string   `json:"subarea_uuid" gorm:"type:varchar(255);not null"`
	SubArea      SubArea  `gorm:"foreignKey:SubAreaUUID;references:UUID"`

	Signature string `json:"signature"`

	// Cyclos []Cyclo `gorm:"foreignKey:CommuneUUID;references:UUID"`

	RouthePlans []RoutePlan `gorm:"foreignKey:CommuneUUID;references:UUID"`
	Pos         []Pos       `gorm:"foreignKey:CommuneUUID;references:UUID"`
	PosForms    []PosForm   `gorm:"foreignKey:CommuneUUID;references:UUID"`

	Users []User `gorm:"foreignKey:CommuneUUID;references:UUID"`
}
