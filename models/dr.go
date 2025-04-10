package models

import "gorm.io/gorm"

type Dr struct {
	gorm.Model
	UUID string `gorm:"not null;unique" json:"uuid"`

	CountryUUID  string   `json:"country_uuid" gorm:"type:varchar(255);not null"`
	Country      Country  `gorm:"foreignKey:CountryUUID;references:UUID"`
	ProvinceUUID string   `json:"province_uuid" gorm:"type:varchar(255);not null"`
	Province     Province `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	AreaUUID     string   `json:"area_uuid" gorm:"type:varchar(255);not null"`
	Area         Area     `gorm:"foreignKey:AreaUUID;references:UUID"`
	SubAreaUUID  string   `json:"subarea_uuid" gorm:"type:varchar(255);not null"`
	SubArea      SubArea  `gorm:"foreignKey:SubAreaUUID;references:UUID"`

	AsmUUID string `json:"asm_uuid" gorm:"type:varchar(255);not null"`
	Asm     Asm    `gorm:"foreignKey:AsmUUID;references:UUID"`
	SupUUID string `json:"sup_uuid" gorm:"type:varchar(255);not null"`
	Sup     Sup    `gorm:"foreignKey:SupUUID;references:UUID"`

	Signature string `json:"signature"`

	UserUUID string `json:"user_id" gorm:"type:varchar(255);not null"`
	// User   User `gorm:"foreignKey:UserUUID"`

	PosForms []PosForm `gorm:"foreignKey:DrUUID;references:UUID"`
	Pos      []Pos     `gorm:"foreignKey:DrUUID;references:UUID"`

	Cyclos []Cyclo `gorm:"foreignKey:DrUUID;references:UUID"`
}
