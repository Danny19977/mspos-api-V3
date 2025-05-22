package models

import "gorm.io/gorm"

type Cyclo struct {
	gorm.Model
	UUID string `gorm:"type:text;not null;unique" json:"uuid"`

	CountryUUID  string   `json:"country_uuid" gorm:"type:varchar(255);not null"`
	Country      Country  `gorm:"foreignKey:CountryUUID;references:UUID"`
	ProvinceUUID string   `json:"province_uuid" gorm:"type:varchar(255);not null"`
	Province     Province `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	AreaUUID     string   `json:"area_uuid" gorm:"type:varchar(255);not null"`
	Area         Area     `gorm:"foreignKey:AreaUUID;references:UUID"`
	SubAreaUUID  string   `json:"subarea_uuid" gorm:"type:varchar(255);not null"`
	SubArea      SubArea  `gorm:"foreignKey:SubAreaUUID;references:UUID"`
	CommuneUUID  string   `json:"commune_uuid" gorm:"type:varchar(255);not null"`
	Commune      Commune  `gorm:"foreignKey:CommuneUUID;references:UUID"`

	// AsmUUID string `json:"asm_uuid" gorm:"type:varchar(255);not null"`
	// Asm     Asm    `gorm:"foreignKey:AsmUUID;references:UUID"`
	// SupUUID string `json:"sup_uuid" gorm:"type:varchar(255);not null"`
	// Sup     Sup    `gorm:"foreignKey:SupUUID;references:UUID"`
	// DrUUID  string `json:"dr_uuid" gorm:"type:varchar(255);not null"` //Dr does not display in the list of cyclo
	// Dr      Dr     `gorm:"foreignKey:DrUUID;references:UUID"`         //Dr does not display in the list of cyclo

	Signature string `json:"signature"`

	// UserUUID string `json:"user_uuid" gorm:"type:varchar(255);not null"`
	// User   User `gorm:"foreignKey:UserUUID"`

	Users []User `gorm:"foreignKey:CycloUUID;references:UUID"`

	PosForms []PosForm `gorm:"foreignKey:CycloUUID;references:UUID"`
	// Pos      []Pos     `gorm:"foreignKey:CycloUUID;references:UUID"`
}
