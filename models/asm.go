package models

import "gorm.io/gorm"

type Asm struct {
	gorm.Model
	UUID  string `gorm:"type:text;not null;unique" json:"uuid"`
	Title string `gorm:"not null;default:''" json:"title"`

	Fullname     string `gorm:"-" json:"fullname"` // Ce champ n'est pas pris en charge par gorm
	CountryName  string `gorm:"-" json:"country_name"`
	ProvinceName string `gorm:"-" json:"province_name"`

	CountryUUID  string   `json:"country_uuid" gorm:"type:varchar(255);not null"`
	Country      Country  `gorm:"foreignKey:CountryUUID;references:UUID"`
	ProvinceUUID string   `json:"province_uuid" gorm:"type:varchar(255);not null"`
	Province     Province `gorm:"foreignKey:ProvinceUUID;references:UUID"`

	Signature string `json:"signature"`

	Users []User `gorm:"foreignKey:AsmUUID;references:UUID"`

	Sups   []Sup   `gorm:"foreignKey:AsmUUID;references:UUID"`
	Drs    []Dr    `gorm:"foreignKey:AsmUUID;references:UUID"`
	Cyclos []Cyclo `gorm:"foreignKey:AsmUUID;references:UUID"`

	Pos      []Pos     `gorm:"foreignKey:AsmUUID;references:UUID"`
	PosForms []PosForm `gorm:"foreignKey:AsmUUID;references:UUID"`
}
