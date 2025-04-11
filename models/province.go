package models

import "gorm.io/gorm"

type Province struct {
	gorm.Model
	UUID string `gorm:"type:text;not null;unique" json:"uuid"`

	Name        string  `json:"name"`
	CountryUUID string  `json:"country_uuid" gorm:"type:varchar(255);not null"`
	Country     Country `gorm:"foreignKey:CountryUUID;references:UUID"`
	Signature   string  `json:"signature"`

	Users    []User    `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	Areas    []Area    `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	SubAreas []SubArea `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	Communes []Commune `gorm:"foreignKey:ProvinceUUID;references:UUID"`

	Brands   []Brand   `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	RoutePlan []RoutePlan `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	Posforms []PosForm `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	Pos      []Pos     `gorm:"foreignKey:ProvinceUUID;references:UUID"`

	Asms   []Asm   `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	Sups   []Sup   `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	Drs    []Dr    `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	Cyclos []Cyclo `gorm:"foreignKey:ProvinceUUID;references:UUID"`
}
