package models

import "gorm.io/gorm"

type RoutePlan struct {
	gorm.Model
	UUID string `gorm:"not null;unique" json:"uuid"`

	UserUUID string `json:"user_uuid" gorm:"type:varchar(255);not null"`
	User     User   `gorm:"foreignKey:UserUUID;references:UUID"`

	ProvinceUUID string   `json:"province_uuid" gorm:"type:varchar(255);not null"`
	Province     Province `gorm:"foreignKey:ProvinceUUID;references:UUID"`
	
	AreaUUID     string   `json:"area_uuid" gorm:"type:varchar(255);not null"`
	Area         Area     `gorm:"foreignKey:AreaUUID;references:UUID"`
	
	SubAreaUUID  string   `json:"subarea_uuid" gorm:"type:varchar(255);not null"`
	SubArea      SubArea  `gorm:"foreignKey:SubAreaUUID;references:UUID"`
	
	CommuneUUID  string   `json:"commune_uuid" gorm:"type:varchar(255);not null"`
	Commune      Commune  `gorm:"foreignKey:CommuneUUID;references:UUID"`

	// TotalPOS  int    `json:"total_pos"`
	Signature string `json:"signature"`

	RutePlanItems []RutePlanItem `gorm:"foreignKey:RoutePlanID;references:UUID"`
}
