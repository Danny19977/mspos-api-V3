package models

import "gorm.io/gorm"

type Sup struct {
	gorm.Model
	UUID string `gorm:"type:text;not null;unique" json:"uuid"`

	CountryUUID  string   `json:"country_uuid" gorm:"type:varchar(255);not null"`
	Country      Country  `gorm:"foreignKey:CountryUUID;references:UUID"`
	ProvinceUUID string   `json:"province_uuid" gorm:"type:varchar(255);not null"`
	Province     Province `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	AreaUUID     string   `json:"area_uuid" gorm:"type:varchar(255);not null"`
	Area         Area     `gorm:"foreignKey:AreaUUID;references:UUID"`

	AsmUUID string `json:"asm_uuid" gorm:"type:varchar(255);not null"`
	Asm     Asm    `gorm:"foreignKey:AsmUUID;references:UUID"`

	UserUUID string `json:"user_uuid" gorm:"type:varchar(255);not null"`
	// User   User `gorm:"foreignKey:UserUUID"`

	Signature string `json:"signature"`

	Drs    []Dr    `gorm:"foreignKey:SupUUID;references:UUID"`
	Cyclos []Cyclo `gorm:"foreignKey:SupUUID;references:UUID"`

	PosForms []PosForm `gorm:"foreignKey:SupUUID;references:UUID"`
	Pos      []Pos     `gorm:"foreignKey:SupUUID;references:UUID"`
}
