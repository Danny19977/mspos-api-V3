package models

import (
	"gorm.io/gorm"
)

type RutePlanItem struct {
	gorm.Model
	UUID string `gorm:"not null;unique" json:"uuid"`

	RoutePlanID string      `json:"routeplan_id" gorm:"type:varchar(255);not null"`
	RoutePlan   RoutePlan `gorm:"foreignKey:RoutePlanID;references:UUID"`

	PosUUID string `json:"pos_uuid" gorm:"type:varchar(255);not null"`
	Pos     Pos  `gorm:"foreignKey:PosUUID;references:UUID"`

	Status bool `json:"status"`
}
