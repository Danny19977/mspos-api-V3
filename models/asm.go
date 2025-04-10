package models

import "gorm.io/gorm"

type Asm struct {
	gorm.Model
	UUID string `gorm:"type:text;not null;unique" json:"uuid"`

	CountryUUID  string   `json:"country_uuid" gorm:"type:varchar(255);not null"`
	Country      Country  `gorm:"foreignKey:CountryUUID;references:UUID"`
	ProvinceUUID string   `json:"province_uuid" gorm:"type:varchar(255);not null"`
	Province     Province `gorm:"foreignKey:ProvinceUUID;references:UUID"`

	Signature string `json:"signature"`

	UserUUID string `json:"user_id" gorm:"type:varchar(255);not null"`
	// User   User `gorm:"foreignKey:UserUUID"`

	Sups   []Sup   `gorm:"foreignKey:AsmUUID;references:UUID"`
	Drs    []Dr    `gorm:"foreignKey:AsmUUID;references:UUID"`
	Cyclos []Cyclo `gorm:"foreignKey:AsmUUID;references:UUID"`

	Pos      []Pos     `gorm:"foreignKey:AsmUUID;references:UUID"`
	PosForms []PosForm `gorm:"foreignKey:AsmUUID;references:UUID"`
}
