package models

import "gorm.io/gorm"

type Dr struct {
	gorm.Model
	UUID string `gorm:"not null;unique" json:"uuid"`

	Title string `gorm:"not null;default:''" json:"title"`
	Fullname    string `gorm:"-" json:"fullname"`
	AsmFullname string `gorm:"-" json:"asm_fullname"`
	SupFullname string `gorm:"-" json:"sup_fullname"`

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

	Users  []User  `gorm:"foreignKey:DrUUID;references:UUID"`
	Cyclos []Cyclo `gorm:"foreignKey:DrUUID;references:UUID"`

	PosForms []PosForm `gorm:"foreignKey:DrUUID;references:UUID"`
	Pos      []Pos     `gorm:"foreignKey:DrUUID;references:UUID"`

}
