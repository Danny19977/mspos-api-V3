package models

import (
	"time"

	"gorm.io/gorm"
)

type RoutePlan struct {
	UUID string `gorm:"type:text;not null;unique;primaryKey" json:"uuid"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	UserUUID string `json:"user_uuid" gorm:"type:varchar(255);not null"`
	User     User   `gorm:"foreignKey:UserUUID;references:UUID"`

	CountryUUID string  `json:"country_uuid" gorm:"type:varchar(255);not null"`
	Country     Country `gorm:"foreignKey:CountryUUID;references:UUID"`

	ProvinceUUID string   `json:"province_uuid" gorm:"type:varchar(255);not null"`
	Province     Province `gorm:"foreignKey:ProvinceUUID;references:UUID"`

	AreaUUID string `json:"area_uuid" gorm:"type:varchar(255);not null"`
	Area     Area   `gorm:"foreignKey:AreaUUID;references:UUID"`

	SubAreaUUID string  `json:"subarea_uuid" gorm:"type:varchar(255);not null"`
	SubArea     SubArea `gorm:"foreignKey:SubAreaUUID;references:UUID"`

	CommuneUUID string  `json:"commune_uuid" gorm:"type:varchar(255);not null"`
	Commune     Commune `gorm:"foreignKey:CommuneUUID;references:UUID"`

	AsmUUID   string `json:"asm_uuid" gorm:"type:varchar(255);not null"`
	Asm       string `json:"asm" gorm:"default:''"`
	SupUUID   string `json:"sup_uuid" gorm:"type:varchar(255);not null"`
	Sup       string `json:"sup" gorm:"default:''"`
	DrUUID    string `json:"dr_uuid" gorm:"type:varchar(255);not null"`
	Dr        string `json:"dr" gorm:"default:''"`
	CycloUUID string `json:"cyclo_uuid" gorm:"type:varchar(255);not null"`
	Cyclo     string `json:"cyclo" gorm:"default:''"`

	// TotalPOS  int    `json:"total_pos"`
	Signature string `json:"signature"`

	RoutePlanItems []RoutePlanItem `gorm:"foreignKey:RoutePlanUUID;references:UUID"`
}
