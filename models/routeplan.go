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

	// TotalPOS  int    `json:"total_pos"`
	Signature string `json:"signature"`

	// TotalRoutePlanItemActive int64 `json:"total_route_plan_item_active" gorm:"-"`
	// TotalRoutePlanItem       int64 `json:"total_route_plan_item" gorm:"-"`

	RoutePlanItems []RoutePlanItem `gorm:"foreignKey:RoutePlanUUID;references:UUID"`
}
