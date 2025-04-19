package models

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	UUID string `gorm:"type:text;not null;unique" json:"uuid"`

	Name      string `gorm:"not null" json:"name"`
	Signature string `json:"signature"`

	Provinces []Province `gorm:"foreignKey:CountryUUID;references:UUID"`
	Areas     []Area     `gorm:"foreignKey:CountryUUID;references:UUID"`
	SubAreas  []SubArea  `gorm:"foreignKey:CountryUUID;references:UUID"`
	Communes  []Commune  `gorm:"foreignKey:CountryUUID;references:UUID"`

	Managers []Manager `gorm:"foreignKey:CountryUUID;references:UUID"`
	Asms     []Asm     `gorm:"foreignKey:CountryUUID;references:UUID"`
	Sups     []Sup     `gorm:"foreignKey:CountryUUID;references:UUID"`
	Drs      []Dr      `gorm:"foreignKey:CountryUUID;references:UUID"`
	Cyclos   []Cyclo   `gorm:"foreignKey:CountryUUID;references:UUID"`

	Brands []Brand `gorm:"foreignKey:CountryUUID;references:UUID"`
	Pos    []Pos   `gorm:"foreignKey:CountryUUID;references:UUID"`
	PosForms []PosForm `gorm:"foreignKey:CountryUUID;references:UUID"`

	Users      []User      `gorm:"foreignKey:CountryUUID;references:UUID"`
	RoutePlans []RoutePlan `gorm:"foreignKey:CountryUUID;references:UUID"`
}
